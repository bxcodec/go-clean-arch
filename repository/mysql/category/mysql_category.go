package category

import (
	"database/sql"
	"fmt"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/bxcodec/go-clean-arch/repository"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

func (m *mysqlCategoryRepository) fetch(query string, args ...interface{}) ([]*models.Category, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		fmt.Println("ERROR DB , ", err.Error())
		return nil, models.NewErrorInternalServer()
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
			fmt.Println(" NEVER NEVER ", err, "")
			return nil, models.NewErrorInternalServer()
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlCategoryRepository) Fetch(cursor string, num int64) ([]*models.Category, error) {

	query := `SELECT id,name,tag,updated_at, created_at
  						FROM category WHERE ID > ? LIMIT ?`

	return m.fetch(query, cursor, num)

}
func (m *mysqlCategoryRepository) GetByID(id int64) (*models.Category, error) {
	query := `SELECT id,name,tag,updated_at, created_at
  						FROM category WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &models.Category{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NewErrorNotFound()
	}

	return a, nil
}

func (m *mysqlCategoryRepository) GetByName(title string) (*models.Category, error) {
	query := `SELECT id,name,tag,updated_at, created_at
  						FROM category WHERE title = ?`

	list, err := m.fetch(query, title)
	if err != nil {
		return nil, err
	}

	a := &models.Category{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NewErrorNotFound()
	}
	return a, nil
}

func (m *mysqlCategoryRepository) Store(a *models.Category) (int64, error) {

	// query := `INSERT INTO category(title, content, created_at , updated_at) VALUES(? , ? , ?, ? )`
	query := `INSERT  category SET name=? , tag=? , updated_at=? , created_at=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, models.NewErrorInternalServer()
	}
	res, err := stmt.Exec(a.Name, a.Tag, a.CreatedAt, a.UpdatedAt)
	if err != nil {

		return 0, models.NewErrorInternalServer()
	}
	return res.LastInsertId()
}
func NewMysqlCategoryRepository(Conn *sql.DB) repository.CategoryRepository {

	return &mysqlCategoryRepository{Conn}
}

func (m *mysqlCategoryRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM category WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, models.NewErrorInternalServer()
	}
	res, err := stmt.Exec(id)
	if err != nil {

		return false, models.NewErrorInternalServer()
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, models.NewErrorInternalServer()
	}
	if rowsAfected <= 0 {
		return false, models.NewErrorInternalServer()
	}

	return true, nil
}
