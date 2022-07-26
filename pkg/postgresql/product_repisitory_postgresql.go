package postgresql

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"database/sql"
)

type ProductRepositoryPostgreSQL struct {
	Conn *sql.DB
}

func NewProductRepositoryPostgreSQL(conn *sql.DB) repository.ProductRepository {
	return &ProductRepositoryPostgreSQL{
		Conn: conn,
	}
}

func (repo *ProductRepositoryPostgreSQL) InsertOne(product entity.Product) (*entity.Product, error) {
	query := `
		INSERT INTO 
		products(name, description, main_img, price, category_id, created_by, type)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	row := repo.Conn.QueryRow(query, product.Name, product.Description, product.MainImg, product.Price, product.CategoryId, product.CreatedBy, product.Type)

	if err := row.Scan(&product.ID, &product.CreatedAt); err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepositoryPostgreSQL) VerifyProductTypeIsExists(productTypeId int) error {
	query := "SELECT type FROM product_types WHERE id = $1"

	row := repo.Conn.QueryRow(query, productTypeId)

	var typeName string

	if err := row.Scan(&typeName); err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepositoryPostgreSQL) FindAllAndCount(param repository.FindAllProductsParam) ([]entity.Product, int, error) {
	query := `
		SELECT p.id, p.name, p.description, p.main_img, p.price, p.created_by, p.category_id, c.name AS category, p.created_at
		FROM products AS p
		INNER JOIN categories AS c ON c.id = p.category_id
	`

	queryCount := `
		SELECT COUNT(p.id) AS total_product FROM products AS p
		INNER JOIN categories AS c ON c.id = p.category_id
	`

	var rows = new(sql.Rows)
	var err error

	if param.Limit > 0 && param.Skip >= 0 {
		query = `
			SELECT p.id, p.name, p.description, p.main_img, p.price, p.created_by, p.category_id, c.name AS category, p.created_at
			FROM products AS p
			INNER JOIN categories AS c ON c.id = p.category_id
			LIMIT $1 OFFSET $2
		`
		rows, err = repo.Conn.Query(query, param.Limit, param.Skip)
	} else {
		rows, err = repo.Conn.Query(query)
	}

	if err != nil {
		return []entity.Product{}, 0, err
	}

	defer rows.Close()

	products := []entity.Product{}

	for rows.Next() {
		var product entity.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.MainImg, &product.Price, &product.CreatedBy, &product.CategoryId, &product.Category, &product.CreatedAt); err != nil {
			return []entity.Product{}, 0, err
		}

		products = append(products, product)
	}

	var count int
	row := repo.Conn.QueryRow(queryCount)

	if err := row.Scan(&count); err != nil {
		return products, count, err
	}

	return products, count, nil
}
