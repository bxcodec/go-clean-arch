package article

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func (m *mysqlArticleRepository) fetch(query string, args ...interface{}) ([]*models.Article, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Article, 0)
	for rows.Next() {
		t := new(models.Article)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlArticleRepository) Fetch(cursor string, num int64) ([]*models.Article, error) {

	query := `SELECT id,title,content,updated_at, created_at
  						FROM article WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}

func NewMysqlArticleRepository(Conn *sql.DB) repository.ArticleRepository {

	return &mysqlArticleRepository{Conn}
}
