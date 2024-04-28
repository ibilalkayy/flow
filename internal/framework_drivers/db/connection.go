package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ibilalkayy/flow/internal/entities"
	"github.com/ibilalkayy/flow/internal/middleware"
	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	var dv entities.DatabaseVariables

	if middleware.LoadEnvVariable("DB_HOST") != "" {
		dv = entities.DatabaseVariables{
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

func TableExists(tableName string) (bool, error) {
	db, err := Connection()
	if err != nil {
		return false, err
	}

	var exists bool
	query := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name = $1
	)`

	err = db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
