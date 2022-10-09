package mysql

import (
	"context"
	"database/sql"

	"github.com/bxcodec/go-clean-arch/article/repository"
	"github.com/bxcodec/go-clean-arch/domain"
)

const AuthorRepoName = "author_mysql_repo"

var _ repository.Repository = &mysqlAuthorRepo{}
var _ domain.AuthorRepository = &mysqlAuthorRepo{}

type mysqlAuthorRepo struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewMysqlAuthorRepository(db *sql.DB) domain.AuthorRepository {
	return &mysqlAuthorRepo{
		DB: db,
	}
}

func (m *mysqlAuthorRepo) Name() string {
	return AuthorRepoName
}

func (m *mysqlAuthorRepo) Driver() string {
	return DriverMySQL
}

func (m *mysqlAuthorRepo) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *mysqlAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}
