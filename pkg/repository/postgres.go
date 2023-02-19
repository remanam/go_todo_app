package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// МИГРАЦИИ пример
//migrate -path ./schema -database postgresql://postgres:00001111@localhost:5432/postgres?sslmode=disable -verbose up

type Config struct {
	// Example  "postgres://postgres:00001111@localhost:5432/todo_erik"
	//"postgres://postgres:00001111@localhost:5432/postgres"
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

const usersTable = "users"

func NewPostgresDB(cfg Config) (*pgx.Conn, error) {

	DATABASE_URL := fmt.Sprintf("postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName)
	fmt.Println(DATABASE_URL)

	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		return nil, err
	}
	//defer conn.Close(context.Background())

	fmt.Println("Connected to database succsessfully")
	return conn, nil

}
