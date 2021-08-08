package storage

// We trust the Usecase layer performed all required validations, so this package
// will just (try to) perform any operation requested, no questions asked.

import (
	"database/sql"
	"fmt"
	"log"

	// needed for the golang generic sql driver
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/model"
)

type DBService struct {
	DB *sql.DB
}

type ConnectionParams struct {
	Driver string
	User   string
	Pass   string
	Server string
	DB     string
}

const (
	TagTable = "tag"
)

// New object with an open DB connection able to query the server
func New(db *sql.DB) *DBService {
	return &DBService{
		DB: db,
	}
}

// ConnectToDB should be moved to the startup config step at some point later
// so I don't want to have it highly coupled to this package
func ConnectToDB(p ConnectionParams) *sql.DB {
	source := fmt.Sprintf("%s:%s@%s/%s", p.User, p.Pass, p.Server, p.DB)
	db, err := sql.Open(p.Driver, source)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// CreateTag return the ID of the tag inserted
func (dbc *DBService) CreateTag(name string) (int, error) {
	query := "INSERT INTO " + TagTable + " (tag) VALUES (?)"

	res, err := dbc.DB.Exec(query, name)
	if err != nil {
		log.Printf("ERROR inserting tag: %v\n", err)
		return 0, err
	}

	inserted, _ := res.LastInsertId()
	log.Printf("ID of the insert: inserted %d\n", inserted)

	return int(inserted), nil
}

func (dbc *DBService) DeleteTag(tag string) (int, error) {
	query := "DELETE FROM " + TagTable + " WHERE tag = (?)"

	res, err := dbc.DB.Exec(query, tag)
	if err != nil {
		log.Printf("ERROR deleting tag: %v\n", err)
		return 0, err
	}

	affected, _ := res.RowsAffected()
	log.Printf("Number of rows affected: %d\n", affected)

	return int(affected), nil
}

func (dbc *DBService) SearchTag(tag string) ([]model.Tag, error) {
	var tags []model.Tag
	query := "SELECT * FROM " + TagTable + " WHERE tag = (?)"

	rows, err := dbc.DB.Query(query, tag)
	if err != nil {
		log.Printf("ERROR deleting tag: %v\n", err)
		return nil, err
	}

	for rows.Next() {
		var tag model.Tag
		err = rows.Scan(&tag.ID, &tag.Tag)

		if err != nil {
			log.Printf("Error scanning the search result %v\n", err)
			return nil, err
		}

		tags = append(tags, tag)
	}

	log.Printf("Rows found: %v\n", tag)

	return tags, nil
}
