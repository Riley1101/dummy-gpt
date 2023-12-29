package database

import (
	"database/sql"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserQuery struct {
	CreateUserTable string
	TableExists     string
	InsertUser      string
}

var UserQuerySQL = UserQuery{
	CreateUserTable: `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`,

	TableExists: `SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users')`,

	InsertUser: `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`,
}

type UserDb struct {
	DB *sql.DB
}

func (u *UserDb) CheckUserExists(username string) (bool, error) {
	var exists bool
	err := u.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (u *UserDb) CreateUser(user *User) error {
	_, err := u.DB.Exec(UserQuerySQL.InsertUser, user.Username, user.Password, user.Email)
	return err
}

func (u *UserDb) CheckTableExists() (bool, error) {
	var exists bool
	err := u.DB.QueryRow(UserQuerySQL.TableExists).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (u *UserDb) CreateUserTable() error {
	exists, err := u.CheckTableExists()
	if exists {
		return nil
	}
	_, err = u.DB.Exec(UserQuerySQL.CreateUserTable)
	return err
}
