package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	repository "github.com/bxcodec/go-clean-arch/author/repository/mysql"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err := db.Close()
		require.NoError(t, err)
	}()

	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(1, "Iman Tumorang", time.Now(), time.Now())

	query := "SELECT id, name, created_at, updated_at FROM author WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlAuthorRepository(db)

	anArticle, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
