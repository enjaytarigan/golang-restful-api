CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    main_img TEXT NOT NULL,
    price BIGINT NOT NULL,
    category_id INT NOT NULL,
    created_at  TIMESTAMP DEFAULT current_timestamp,
    created_by INT NOT NULL,
    FOREIGN KEY(created_by) REFERENCES users(id),
    FOREIGN KEY(category_id) REFERENCES categories(id)
);