package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(username, password, host, port, dbName string) (*MySQL, error) {
	connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", username, password, host, port, dbName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

func (m *MySQL) Connect() error {
	return m.db.Ping()
}

func (m *MySQL) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
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

func (m *MySQL) Exec(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	return err
}

func (m *MySQL) Close() error {
	return m.db.Close()
}
