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
	var v Variables

	if os.Getenv("POSTGRES_HOST") != "" {
		v = Variables{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("POSTGRES_SSL"),
		}
	} else {
		v = Variables{
			Host:     middleware.LoadEnvVariable("LOCAL_HOST"),
			Port:     middleware.LoadEnvVariable("PORT"),
			User:     middleware.LoadEnvVariable("POSTGRES_USER"),
			Password: middleware.LoadEnvVariable("PASSWORD"),
			DBName:   middleware.LoadEnvVariable("DBNAME"),
			SSLMode:  middleware.LoadEnvVariable("SSLMODE"),
		}
	}

	connectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", v.Host, v.Port, v.User, v.Password, v.DBName, v.SSLMode)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Table(basePath, filename string, number int) (*sql.DB, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	query, err := os.ReadFile(basePath + filename)
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
