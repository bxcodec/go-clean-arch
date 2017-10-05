package mysql

import (
	"database/sql"

	models "github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func (m *mysqlArticleRepository) fetch(query string, args ...interface{}) ([]*models.Article, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {

		return nil, models.INTERNAL_SERVER_ERROR
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

			return nil, models.INTERNAL_SERVER_ERROR
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
func (m *mysqlArticleRepository) GetByID(id int64) (*models.Article, error) {
	query := `SELECT id,title,content,updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &models.Article{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *mysqlArticleRepository) GetByTitle(title string) (*models.Article, error) {
	query := `SELECT id,title,content,updated_at, created_at
  						FROM article WHERE title = ?`

	list, err := m.fetch(query, title)
	if err != nil {
		return nil, err
	}

	a := &models.Article{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, models.NOT_FOUND_ERROR
	}
	return a, nil
}

func (m *mysqlArticleRepository) Store(a *models.Article) (int64, error) {

	query := `INSERT  article SET title=? , content=? , updated_at=? , created_at=?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {

		return 0, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(a.Title, a.Content, a.CreatedAt, a.UpdatedAt)
	if err != nil {

		return 0, models.INTERNAL_SERVER_ERROR
	}
	return res.LastInsertId()
}

func (m *mysqlArticleRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM article WHERE id = ?"

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, models.INTERNAL_SERVER_ERROR
	}
	res, err := stmt.Exec(id)
	if err != nil {

		return false, models.INTERNAL_SERVER_ERROR
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, models.INTERNAL_SERVER_ERROR
	}
	if rowsAfected <= 0 {
		return false, models.INTERNAL_SERVER_ERROR
	}

	return true, nil
}
func (m *mysqlArticleRepository) Update(ar *models.Article) (*models.Article, error) {
	query := `UPDATE article set title=?, content=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, nil
	}
	res, err := stmt.Exec(ar.Title, ar.Content, ar.UpdatedAt, ar.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect < 1 {
		return nil, models.INTERNAL_SERVER_ERROR
	}

	return ar, nil
}

func NewMysqlArticleRepository(Conn *sql.DB) repository.ArticleRepository {

	return &mysqlArticleRepository{Conn}
}
