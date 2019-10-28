package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/7onetella/password/api/model"
	"github.com/boltdb/bolt"
	_ "github.com/lib/pq"
)

// DataSource is a wrapper for sql.DB
type DataSource struct {
	db *sql.DB
}

// NewDataSource returns a new instance of DataSource
func NewDataSource() *DataSource {
	return &DataSource{db}
}

// Query queries database
func (ds *DataSource) Query(query string, args ...interface{}) ([]model.Password, error) {

	var rows *sql.Rows
	var err error
	log.Println("query:", query)
	rows, err = ds.db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return passwordRowsMapper(rows)
}

// OpenDB opens password database
func OpenDB(path string) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

// CreatePassword creates password record in db
func CreatePassword(p model.Password) (string, error) {
	p.ID = getUUID()

	statement := `
		INSERT INTO passwords (id, title, url, username, password, notes, tags, admin_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	tags := strings.Join(p.Tags, ",")
	_, err := db.Exec(statement, p.ID, p.Title, p.URL, p.Username, p.Password, p.Notes, tags, p.AdminID)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}

// ReadPassword reads password record from db
func ReadPassword(ID string) (model.Password, error) {
	row := db.QueryRow("SELECT * FROM passwords WHERE id = $1", ID)

	return passwordRowMapper(row)
}

// UpdatePassword updates password record in db
func UpdatePassword(p model.Password) error {
	statement := `
		UPDATE passwords SET title=$1, url=$2, username=$3, password=$4, notes=$5, tags=$6 WHERE id=$7`
	tags := strings.Join(p.Tags, ",")
	_, err := db.Exec(statement, p.Title, p.URL, p.Username, p.Password, p.Notes, tags, p.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePassword deletes password record in db
func DeletePassword(id string) error {
	statement := `
		DELETE FROM passwords WHERE ID=$1`
	_, err := db.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}

// FindPasswordsByTitle will find password records by username
func FindPasswordsByTitle(rid, title, nextToken string, size int, adminID string) ([]model.Password, error) {
	ds := NewDataSource()
	return ds.Query(`SELECT * FROM passwords WHERE title like $1 AND id > $2 AND admin_id = $4 ORDER BY id ASC LIMIT $3`, title, nextToken, size, adminID)
}

// ListAllPasswords lists all passwords with pagination
func ListAllPasswords(nextToken string, size int, adminID string) ([]model.Password, error) {
	ds := NewDataSource()
	return ds.Query("SELECT * FROM passwords WHERE id > $1 AND admin_id = $3 ORDER BY id ASC LIMIT $2", nextToken, size, adminID)
}

// DeleteAllPasswords deletes all passwords
func DeleteAllPasswords() error {
	statement := `
		DELETE FROM passwords`
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	return nil
}

// good article on pagination
// https://www.citusdata.com/blog/2016/03/30/five-ways-to-paginate/

func passwordRowsMapper(rows *sql.Rows) ([]model.Password, error) {
	passwords := []model.Password{}

	for {
		if !rows.Next() {
			return passwords, nil
		}
		p := model.Password{}

		var tagsRaw string
		rows.Scan(&p.ID, &p.Title, &p.URL, &p.Username, &p.Password, &p.Notes, &tagsRaw, &p.AdminID)
		p.Tags = strings.Split(tagsRaw, ",")

		passwords = append(passwords, p)
	}
}

func passwordRowMapper(row *sql.Row) (model.Password, error) {
	p := model.Password{}

	var tagsRaw string
	err := row.Scan(&p.ID, &p.Title, &p.URL, &p.Username, &p.Password, &p.Notes, &tagsRaw, &p.AdminID)
	if err != nil {
		return p, err
	}
	p.Tags = strings.Split(tagsRaw, ",")

	return p, nil
}
