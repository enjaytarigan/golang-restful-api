package postgresql

import (
	"brodo-demo/entity"
	"brodo-demo/repository"
	"database/sql"
	"errors"
)

type UserRepositoryPostgreSQL struct {
	Conn *sql.DB
}

func NewUserRepositoryPostgreSQL(connection *sql.DB) repository.UserRepository {
	return &UserRepositoryPostgreSQL{
		Conn: connection,
	}
}

func (repo *UserRepositoryPostgreSQL) Insert(user entity.User) (userId int, err error) {
	queryCommand := "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id"
	row := repo.Conn.QueryRow(queryCommand, user.Username, user.Password)

	if err := row.Scan(&userId); err != nil {
		return userId, err
	}

	return userId, nil
}

func (repo *UserRepositoryPostgreSQL) VerifyAvailableUsername(username string) bool {
	query := "SELECT username FROM users WHERE username = $1"

	row := repo.Conn.QueryRow(query, username)

	var tempUsername string

	if err := row.Scan(&tempUsername); err != nil {
		return true // username available for register
	}

	return false // username already used
}

func (repo *UserRepositoryPostgreSQL) FindByUsername(username string) (entity.User, error) {
	user := entity.User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"

	row := repo.Conn.QueryRow(query, username)

	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}

		return user, err
	}

	return user, nil
}

func (repo *UserRepositoryPostgreSQL) VerifyUserIsExist(userId int) error {
	return nil
}
