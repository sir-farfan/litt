package usecase

import "gitlab.com/codelittinc/golang-interview-project-ismael-estrada/model"

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (u *Tag) Create(tag *model.Tag) (*model.Tag, error) {
	return &model.Tag{Tag: "yepee"}, nil
}
