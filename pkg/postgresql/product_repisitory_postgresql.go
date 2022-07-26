package postgresql

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"database/sql"
	"errors"
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
		products(name, description, main_img, price, category_id, created_by)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	row := repo.Conn.QueryRow(query, product.Name, product.Description, product.MainImg, product.Price, product.CategoryId, product.CreatedBy)

	if err := row.Scan(&product.ID, &product.CreatedAt); err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepositoryPostgreSQL) FindAllAndCount(param repository.FindAllProductsParam) ([]entity.Product, int, error) {
	query := `
		SELECT p.id, p.name, p.description, p.main_img, p.price, p.created_by, p.category_id, c.name AS category, p.created_at
		FROM products AS p
		INNER JOIN categories AS c ON c.id = p.category_id
		LIMIT $1 OFFSET $2
	`

	queryCount := `
		SELECT COUNT(p.id) AS total_product FROM products AS p
	`

	var rows = new(sql.Rows)
	var err error
	var row *sql.Row

	if param.MinPrice > 0 && param.MaxPrice > param.MinPrice {
		query = `
			SELECT p.id, p.name, p.description, p.main_img, p.price, p.created_by, p.category_id, c.name AS category, p.created_at
			FROM products AS p
			INNER JOIN categories AS c ON c.id = p.category_id
			WHERE p.price BETWEEN $1 AND $2
			LIMIT $3 OFFSET $4
		`

		queryCount = `
			SELECT COUNT(p.id) AS total_product FROM products AS p
			WHERE p.price BETWEEN $1 AND $2
		`

		rows, err = repo.Conn.Query(query, param.MinPrice, param.MaxPrice, param.Limit, param.Skip)
		row = repo.Conn.QueryRow(queryCount, param.MinPrice, param.MaxPrice)

	} else {
		rows, err = repo.Conn.Query(query, param.Limit, param.Skip)
		row = repo.Conn.QueryRow(queryCount)

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
	if err := row.Scan(&count); err != nil {
		return products, count, err
	}

	return products, count, nil
}

func (repo *ProductRepositoryPostgreSQL) FindById(productId int) (entity.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.main_img, p.category_id, p.created_at, p.created_by, c.name AS category_name
		FROM products AS p
		INNER JOIN categories AS c ON c.id = p.category_id
		WHERE p.id = $1
	`

	row := repo.Conn.QueryRow(query, productId)

	product := entity.Product{}

	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.MainImg, &product.CategoryId, &product.CreatedAt, &product.CreatedBy, &product.Category); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Product{}, errors.New("product not found")
		}

		return entity.Product{}, err
	}

	return product, nil
}