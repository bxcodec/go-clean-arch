package category_test

import (
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/models"
	categoryRepo "github.com/bxcodec/go-clean-arch/repository/mysql/category"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "tag", "updated_at", "created_at"}).
		AddRow(1, "name 1", "Tag 1", time.Now(), time.Now()).
		AddRow(2, "name 2", "Tag 2", time.Now(), time.Now())

	query := "SELECT id,name,tag,updated_at, created_at FROM category WHERE ID > \\? LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := categoryRepo.NewMysqlCategoryRepository(db)
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
	rows := sqlmock.NewRows([]string{"id", "name", "tag", "updated_at", "created_at"}).
		AddRow(1, "name 1", "Tag 1", time.Now(), time.Now())

	query := "SELECT id,name,tag,updated_at, created_at FROM category WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := categoryRepo.NewMysqlCategoryRepository(db)

	num := int64(5)
	anCategory, err := a.GetByID(num)
	assert.NoError(t, err)
	assert.NotNil(t, anCategory)
}

func TestStore(t *testing.T) {

	ar := &models.Category{
		Name:      "Judul",
		Tag:       "Tag",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT  category SET name=\\? , tag=\\? , updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Name, ar.Tag, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))
	// mock.ExpectExec(query).WithArgs(ar.Name, ar.Tag, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))
	a := categoryRepo.NewMysqlCategoryRepository(db)

	lastId, err := a.Store(ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), lastId)
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "tag", "updated_at", "created_at"}).
		AddRow(1, "name 1", "Tag 1", time.Now(), time.Now())

	query := "SELECT id,name,tag,updated_at, created_at FROM category WHERE title = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := categoryRepo.NewMysqlCategoryRepository(db)

	name := "name 1"
	anCategory, err := a.GetByName(name)
	assert.NoError(t, err)
	assert.NotNil(t, anCategory)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM category WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := categoryRepo.NewMysqlCategoryRepository(db)

	num := int64(12)
	anCategoryStatus, err := a.Delete(num)
	assert.NoError(t, err)
	assert.True(t, anCategoryStatus)
}
