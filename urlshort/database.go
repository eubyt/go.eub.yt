package urlshort

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	db *sql.DB
}

type Result struct {
	id   int
	url  string
	code string
}

var DATABASE = Connection{}

func (c *Connection) Connect(name string) {
	var err error
	nameFormat := fmt.Sprintf("/go-eub-yt/%s.db", name)
	c.db, err = sql.Open("sqlite3", nameFormat)
	if err != nil {
		log.Fatal("error opening/create database: ", err)
	}
}

func (c *Connection) CreateTable() {
	statement, err := c.db.Prepare("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY, url TEXT, code TEXT)")
	if err != nil {
		log.Fatal("error createTable: ", err)
	}
	statement.Exec()
}

func (c *Connection) Insert(url, code string) (Result, error) {
	statement, err := c.db.Prepare("INSERT INTO urls (url, code) VALUES (?, ?)")
	if err != nil {
		log.Fatal("error preparing inset: ", err)
	}
	result, err := statement.Exec(url, code)
	if err != nil {
		log.Fatal("error executing inset: ", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("error getting last insert id: ", err)
	}
	return Result{int(id), url, code}, nil
}

func (c *Connection) CheckExistURL(url string) (bool, error) {
	var count int
	statement, err := c.db.Prepare("SELECT COUNT(*) FROM urls WHERE url = ?")
	if err != nil {
		log.Fatal("error prepare checkExistURL: ", err)
	}
	err = statement.QueryRow(url).Scan(&count)
	if err != nil {
		log.Fatal("error querying row: ", err)
	}
	return count > 0, nil
}

func (c *Connection) CheckExistCode(code string) (bool, error) {
	var count int
	statement, err := c.db.Prepare("SELECT COUNT(*) FROM urls WHERE code = ?")
	if err != nil {
		log.Fatal("error prepare checkExistCode: ", err)
	}
	err = statement.QueryRow(code).Scan(&count)
	if err != nil {
		log.Fatal("error querying row: ", err)
	}
	return count > 0, nil
}

func (c *Connection) GetCode(url string) (string, error) {
	var code string
	statement, err := c.db.Prepare("SELECT code FROM urls WHERE url = ?")
	if err != nil {
		log.Fatal("error preparing getCode: ", err)
	}
	err = statement.QueryRow(url).Scan(&code)
	if err != nil {
		log.Fatal("error querying row: ", err)
	}
	return code, nil
}

func (c *Connection) GetURL(code string) (string, error) {
	var url string
	statement, err := c.db.Prepare("SELECT url FROM urls WHERE code = ?")
	if err != nil {
		log.Fatal("error preparing getURL: ", err)
	}
	err = statement.QueryRow(code).Scan(&url)
	if err != nil {
		log.Fatal("error querying row: ", err)
	}
	return url, nil
}

func (c *Connection) AllResult() ([]Result, error) {
	var results []Result
	rows, err := c.db.Query("SELECT id, url, code FROM urls")
	if err != nil {
		log.Fatal("error querying rows: ", err)
	}
	for rows.Next() {
		var result Result
		err = rows.Scan(&result.id, &result.url, &result.code)
		if err != nil {
			log.Fatal("error scanning row: ", err)
		}
		results = append(results, result)
	}
	return results, nil
}
