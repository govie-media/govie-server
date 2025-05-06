package database

import (
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewPostgreSQL(username, password, host, port, dbName string) (*PostgreSQL, error) {
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", username, password, host, port, dbName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgreSQL{db: db}, nil
}

func (m *PostgreSQL) Connect() error {
	return m.db.Ping()
}

func (m *PostgreSQL) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err = rows.Scan(&row)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

func (m *PostgreSQL) Exec(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	return err
}

func (m *PostgreSQL) Close() error {
	return m.db.Close()
}
