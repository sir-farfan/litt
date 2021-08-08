package storage

import (
	"database/sql"
	"fmt"
	"log"

	// needed for the golang generic sql driver
	_ "github.com/go-sql-driver/mysql"
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
	fmt.Println(source)
	db, err := sql.Open(p.Driver, source)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (dbc *DBService) CreateTag(name string) (int, error) {
	// var ctx context.Context
	query := "INSERT INTO " + TagTable + " (tag) VALUES (?)"

	log.Println(query)

	res, err := dbc.DB.Exec(query, name)
	if err != nil {
		log.Printf("ERROR inserting tag: %v\n", err)
		return 0, err
	}

	inserted, _ := res.LastInsertId()
	log.Printf("Result of the insert: inserted %d\n", inserted)

	return int(inserted), nil
}
