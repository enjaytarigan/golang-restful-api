package postgresql

import (
	"brodo-demo/entity"
	"database/sql"
	"errors"
)

type CategoryRepositoryPostgreSQL struct {
	Conn *sql.DB
}

func NewCategoryRepositoryPostgreSQL(conn *sql.DB) *CategoryRepositoryPostgreSQL {
	return &CategoryRepositoryPostgreSQL{
		Conn: conn,
	}
}

func (repo *CategoryRepositoryPostgreSQL) InsertOne(category entity.Category) (*entity.Category, error) {
	query := "INSERT INTO categories(name, created_by) VALUES($1, $2) RETURNING id, name, created_at, created_by"

	row := repo.Conn.QueryRow(query, category.Name, category.CreatedBy)

	newCategory := entity.Category{}

	if err := row.Scan(&newCategory.ID, &newCategory.Name, &newCategory.CreatedAt, &newCategory.CreatedBy); err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (repo *CategoryRepositoryPostgreSQL) UpdateById(category entity.Category) (*entity.Category, error) {
	query := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, name, created_at, created_by"

	row := repo.Conn.QueryRow(query, category.Name, category.ID)

	updatedCategory := &entity.Category{}

	if err := row.Scan(&updatedCategory.ID, &updatedCategory.Name, &updatedCategory.CreatedAt, &updatedCategory.CreatedBy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return updatedCategory, errors.New("category not found")
		}

		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (repo *CategoryRepositoryPostgreSQL) FindById(Id int) (entity.Category, error) {
	query := "SELECT id, name, created_at, created_by FROM categories WHERE id = $1"

	row := repo.Conn.QueryRow(query, Id)

	category := entity.Category{}

	if err := row.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return category, errors.New("category not found")
		}

		return category, err
	}

	return category, nil
}

func (repo *CategoryRepositoryPostgreSQL) FindAll() ([]entity.Category, error) {
	categories := []entity.Category{}
	query := "SELECT id, name, created_at, created_by FROM categories ORDER BY created_at DESC"

	rows, err := repo.Conn.Query(query)

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		category := entity.Category{}
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy); err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repo *CategoryRepositoryPostgreSQL) DeleteById(categoryId int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.Conn.Exec(query, categoryId)

	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowCount == 0 {
		return errors.New("category not found")
	}

	return nil
}
