package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	"github.com/bxcodec/go-clean-arch/models"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	mockArticles := []models.Article{
		models.Article{
			ID: 1, Title: "title 1", Content: "content 1",
			Author: models.Author{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		models.Article{
			ID: 2, Title: "title 2", Content: "content 2",
			Author: models.Author{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Title, mockArticles[0].Content,
			mockArticles[0].Author.ID, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt).
		AddRow(mockArticles[1].ID, mockArticles[1].Title, mockArticles[1].Content,
			mockArticles[1].Author.ID, mockArticles[1].UpdatedAt, mockArticles[1].CreatedAt)

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)
	cursor := articleRepo.EncodeCursor(mockArticles[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)

	num := int64(5)
	anArticle, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &models.Article{
		Title:     "Judul",
		Content:   "Content",
		CreatedAt: now,
		UpdatedAt: now,
		Author: models.Author{
			ID:   1,
			Name: "Iman Tumorang",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	query := "INSERT  article SET title=\\? , content=\\? , author_id=\\?, updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.Author.ID, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}

func TestGetByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()
	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE title = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)

	title := "title 1"
	anArticle, err := a.GetByTitle(context.TODO(), title)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	query := "DELETE FROM article WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	num := int64(12)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ar := &models.Article{
		ID:        12,
		Title:     "Judul",
		Content:   "Content",
		CreatedAt: now,
		UpdatedAt: now,
		Author: models.Author{
			ID:   1,
			Name: "Iman Tumorang",
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	query := "UPDATE article set title=\\?, content=\\?, author_id=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.Author.ID, ar.UpdatedAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}
