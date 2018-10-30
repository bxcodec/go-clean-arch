package repository_test

import (
	"context"
	"testing"
	"time"

	articleRepo "github.com/bxcodec/go-clean-arch/v2/article/repository"
	models "github.com/bxcodec/go-clean-arch/v2/models"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", 1, time.Now(), time.Now()).
		AddRow(2, "title 2", "Content 2", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, author_id, updated_at, created_at FROM article WHERE ID > \\? LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)
	cursor := "sampleCursor"
	num := int64(5)
	list, err := a.Fetch(context.TODO(), cursor, num)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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
	defer db.Close()

	query := "INSERT  article SET title=\\? , content=\\? , author_id=\\?, updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.Author.ID, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	lastId, err := a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), lastId)
}

func TestGetByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
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
	defer db.Close()

	query := "DELETE FROM article WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	num := int64(12)
	anArticleStatus, err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
	assert.True(t, anArticleStatus)
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
	defer db.Close()

	query := "UPDATE article set title=\\?, content=\\?, author_id=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.Author.ID, ar.UpdatedAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := articleRepo.NewMysqlArticleRepository(db)

	s, err := a.Update(context.TODO(), ar)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
