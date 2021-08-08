package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/controller"
)

func Setup(tag controller.Tag) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tag", tag.Create).Methods("POST")
	r.HandleFunc("/tag/{tag_id}", tag.Create).Methods("DELETE")
	r.HandleFunc("/tag", tag.Create).Methods("GET").Queries("name", "*", "id", "[0-9]+")

	return r
}
