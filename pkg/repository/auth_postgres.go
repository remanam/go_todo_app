package repository

import (
	"context"
	"fmt"
	"log"
	"todo"

	"github.com/jackc/pgx/v4"
)

type AuthPostgres struct {
	conn *pgx.Conn
}

func NewAuthPostgres(conn *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{
		conn: conn,
	}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {

	var id int

	query := fmt.Sprintf(`INSERT INTO %s (firstname, username, password_hash) 
						  VALUES ($1, $2, $3) 
						  RETURNING id`, usersTable)
	// Ошибка conn closed

	row := r.conn.QueryRow(context.Background(), query, user.Firstname, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		log.Println("request insert failed")
		return 0, err
	}

	return id, nil
}
