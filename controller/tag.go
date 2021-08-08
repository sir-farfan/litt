package controller

import "net/http"

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating tag\n"))
}
