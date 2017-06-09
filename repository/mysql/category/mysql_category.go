package category

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

func (m *mysqlCategoryRepository) fetch(query string, args ...interface{}) ([]*models.Category, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Category, 0)
	for rows.Next() {
		t := new(models.Category)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Tag,
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

func (m *mysqlCategoryRepository) Fetch(articleID int64) ([]*models.Category, error) {

	query := `SELECT id,name,tag,updated_at, created_at
  						FROM category WHERE ID > ? `

	return m.fetch(query, articleID)

}

func NewMysqlCategoryRepository(Conn *sql.DB) repository.CategoryRepository {

	return &mysqlCategoryRepository{Conn}
}
