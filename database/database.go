package database

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/ilyakaznacheev/cleanenv"
)

type DataBase struct {
	db *sql.DB
}

var (
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrNotFound       = errors.New("not found")
	ErrInvalid        = errors.New("invalid operation")
)

func parseError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return ErrDuplicateEntry
		case 1452:
			return ErrInvalid
		default:
			return err
		}
	}
	return err
}

func Connect() (*DataBase, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, parseError(err)
	}
	return &DataBase{db}, nil
}

