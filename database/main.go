package database

import (
	"errors"
)

type Database interface {
	Connect() error
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)
	Exec(query string, args ...interface{}) error
	Close() error
}

func NewDatabase(dbType string, config map[string]string) (Database, error) {
	switch dbType {
	case "mysql":
		return NewMySQL(config["username"], config["password"], config["host"], config["port"], config["dbname"])
	default:
		return nil, errors.New("unsupported database type")
	}
}
