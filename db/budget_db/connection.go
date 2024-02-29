package budget_db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/ibilalkayy/flow/internal/middleware"
	_ "github.com/lib/pq"
)

type Variables struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Connection() (*sql.DB, error) {
	v := Variables{
		Host:     middleware.LoadEnvVariable("host"),
		Port:     middleware.LoadEnvVariable("port"),
		User:     middleware.LoadEnvVariable("user"),
		Password: middleware.LoadEnvVariable("password"),
		DBName:   middleware.LoadEnvVariable("dbname"),
		SSLMode:  middleware.LoadEnvVariable("sslmode"),
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", v.Host, v.Port, v.User, v.Password, v.DBName, v.SSLMode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Table(filename string, number int) (*sql.DB, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	query, err := os.ReadFile("db/budget_db/migrations/" + filename)
	if err != nil {
		return nil, err
	}

	requests := strings.Split(string(query), ";")[number]
	stmt, err := db.Prepare(requests)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return db, nil
}
