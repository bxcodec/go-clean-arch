package usecase

import (
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/bxcodec/go-clean-arch/author"

	"github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*article.Article, string, error)
	GetByID(id int64) (*article.Article, error)
	Update(ar *article.Article) (*article.Article, error)
	GetByTitle(title string) (*article.Article, error)
	Store(*article.Article) (*article.Article, error)
	Delete(id int64) (bool, error)
}

type articleUsecase struct {
	articleRepos repository.ArticleRepository
	authorRepo   _authorRepo.AuthorRepository
}

type authorChanel struct {
	Author *author.Author
	Error  error
}

func NewArticleUsecase(a repository.ArticleRepository, ar _authorRepo.AuthorRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepos: a,
		authorRepo:   ar,
	}
}

func (a *articleUsecase) getAuthorDetail(item *article.Article, authorChan chan authorChanel) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Debug("Recovered in ", r)
		}
	}()

	res, err := a.authorRepo.GetByID(item.Author.ID)
	holder := authorChanel{
		Author: res,
		Error:  err,
	}
	authorChan <- holder
}
func (a *articleUsecase) getAuthorDetails(data []*article.Article) ([]*article.Article, error) {
	chAuthor := make(chan authorChanel)
	defer close(chAuthor)
	existingAuthorMap := make(map[int64]bool)
	totalCall := 0
	for _, item := range data {
		if _, ok := existingAuthorMap[item.Author.ID]; !ok {
			existingAuthorMap[item.Author.ID] = true
			go a.getAuthorDetail(item, chAuthor)
		}
		totalCall++
	}

	totalReceived := 0
	mapAuthor := make(map[int64]*author.Author)
	receivingDone := false
	for {
		select {
		case a := <-chAuthor:
			totalReceived++
			if a.Error == nil && a.Author != nil {
				mapAuthor[a.Author.ID] = a.Author
			}

		case <-time.After(time.Second * 1):
			logrus.Warn("Timeout when calling user detail")
			receivingDone = true
		}

		if totalReceived == len(existingAuthorMap) {
			receivingDone = true
		}

		if receivingDone {
			break
		}
	}

	// merge the author

	for index, item := range data {
		if a, ok := mapAuthor[item.Author.ID]; ok {
			data[index].Author = *a
		}
	}
	return data, nil
}

func (a *articleUsecase) Fetch(cursor string, num int64) ([]*article.Article, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := a.articleRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	listArticle, err = a.getAuthorDetails(listArticle)
	if err != nil {
		return nil, "", err
	}

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (a *articleUsecase) GetByID(id int64) (*article.Article, error) {

	res, err := a.articleRepos.GetByID(id)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor
	return res, nil
}

func (a *articleUsecase) Update(ar *article.Article) (*article.Article, error) {
	ar.UpdatedAt = time.Now()
	return a.articleRepos.Update(ar)
}

func (a *articleUsecase) GetByTitle(title string) (*article.Article, error) {

	res, err := a.articleRepos.GetByTitle(title)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor

	return res, nil
}

func (a *articleUsecase) Store(m *article.Article) (*article.Article, error) {

	existedArticle, _ := a.GetByTitle(m.Title)
	if existedArticle != nil {
		return nil, article.CONFLIT_ERROR
	}

	id, err := a.articleRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (a *articleUsecase) Delete(id int64) (bool, error) {
	existedArticle, _ := a.articleRepos.GetByID(id)
	logrus.Info("Masuk Sini")
	if existedArticle == nil {
		logrus.Info("Masuk Sini2")
		return false, article.NOT_FOUND_ERROR
	}
	logrus.Info("Masuk Sini3")

	return a.articleRepos.Delete(id)
}
