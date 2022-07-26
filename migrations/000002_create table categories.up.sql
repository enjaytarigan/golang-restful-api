CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    created_by INT NOT NULL,
    FOREIGN KEY(created_by) REFERENCES users(id)
)