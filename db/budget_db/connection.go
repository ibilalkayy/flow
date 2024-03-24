package budget_db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ibilalkayy/flow/internal/middleware"
	"github.com/ibilalkayy/flow/internal/structs"
	_ "github.com/lib/pq"
)

func Connection2(values [6]string) (*sql.DB, error) {
	params := structs.DatabaseVariables{
		Host:     values[0],
		Port:     values[1],
		User:     values[2],
		Password: values[3],
		DBName:   values[4],
		SSLMode:  values[5],
	}

	connectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", params.Host, params.Port, params.User, params.Password, params.DBName, params.SSLMode)
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

func Connection() (*sql.DB, error) {
	var dv structs.DatabaseVariables

	if middleware.LoadEnvVariable("DB_HOST") != "" {
		dv = structs.DatabaseVariables{
			Host:     middleware.LoadEnvVariable("DB_HOST"),
			Port:     middleware.LoadEnvVariable("DB_PORT"),
			User:     middleware.LoadEnvVariable("DB_USER"),
			Password: middleware.LoadEnvVariable("DB_PASSWORD"),
			DBName:   middleware.LoadEnvVariable("DB_NAME"),
			SSLMode:  middleware.LoadEnvVariable("SSL_MODE"),
		}
	} else {
		return nil, errors.New("invalid host provided")
	}

	connectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dv.Host, dv.Port, dv.User, dv.Password, dv.DBName, dv.SSLMode)
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
