package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/controller"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/router"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/service/storage"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/usecase"
)

func getSQLConnection() *storage.DBService {
	params := storage.ConnectionParams{
		Driver: "mysql",
		User:   "root",
		Pass:   "YOUR_PASSWORD_HERE",
		// Server: "localhost",
		DB: "litt",
	}

	dbc := storage.ConnectToDB(params)
	return storage.New(dbc)
}

func setupServer() *http.Server {
	db := getSQLConnection()
	tagUsecase := usecase.NewTag(db)
	tagController := controller.NewTag(tagUsecase)

	r := router.Setup(*tagController)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	return srv
}

func main() {
	fmt.Println("Server started... maybe")
	setupServer()
}
