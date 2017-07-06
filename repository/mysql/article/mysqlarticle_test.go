package article_test

import (
	"testing"
	"time"

	articleRepo "github.com/bxcodec/go-clean-arch/repository/mysql/article"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", time.Now(), time.Now()).
		AddRow(2, "title 2", "Content 2", time.Now(), time.Now())

	query := "SELECT id,title,content,updated_at, created_at FROM article WHERE ID > \\? LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)
	cursor := "sampleCursor"
	num := int64(5)
	list, err := a.Fetch(cursor, num)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", time.Now(), time.Now())

	query := "SELECT id,title,content,updated_at, created_at FROM article WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleRepo.NewMysqlArticleRepository(db)

	num := int64(5)
	anArticle, err := a.GetByID(num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
