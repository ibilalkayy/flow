package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/ibilalkayy/flow/usecases/middleware"
	_ "github.com/lib/pq"
)

type MyConnection struct {
	*handler.Handler
}

func (MyConnection) Connection() (*sql.DB, error) {
	myEnv := middleware.MyEnv{}
	deps := interfaces.Dependencies{
		Env: myEnv,
	}
	h := handler.NewHandler(deps)
	myEnv.Handler = h

	var dv entities.DatabaseVariables
	if h.Deps.Env.LoadEnvVariable("DB_HOST") != "" {
		dv = entities.DatabaseVariables{
			Host:     h.Deps.Env.LoadEnvVariable("DB_HOST"),
			Port:     h.Deps.Env.LoadEnvVariable("DB_PORT"),
			User:     h.Deps.Env.LoadEnvVariable("DB_USER"),
			Password: h.Deps.Env.LoadEnvVariable("DB_PASSWORD"),
			DBName:   h.Deps.Env.LoadEnvVariable("DB_NAME"),
			SSLMode:  h.Deps.Env.LoadEnvVariable("SSL_MODE"),
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

func (h MyConnection) Table(filename string, number int) (*sql.DB, error) {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return nil, err
	}

	query, err := os.ReadFile(filename)
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

func (h MyConnection) TableExists(tableName string) (bool, error) {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return false, err
	}

	var exists bool
	tableName = strings.ToLower(tableName)
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
