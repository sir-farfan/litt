package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/model"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/usecase"
)

type Tag struct {
	useCase usecase.Tag
}

func NewTag(tag *usecase.Tag) *Tag {
	return &Tag{useCase: *tag}
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("creating tag")
	var tag model.Tag

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tagWithID, err := t.useCase.Create(&tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	m, err := json.Marshal(tagWithID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(m)
}

func (t *Tag) Search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching tag\n"))
}

func (t *Tag) Delete(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	// This function only gets called if the tag_id is provided in the correct format 😎
	tag := pathParams["tag_id"]
	tagID, _ := strconv.Atoi(tag)
	_, err := t.useCase.Delete(tagID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
