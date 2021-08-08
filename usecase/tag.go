package usecase

import (
	"errors"
	"fmt"

	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/model"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/service/storage"
)

const (
	SearchParamError = "Can't specify only ID or Tag name at the same time"
)

type Tag struct {
	db *storage.DBService
}

func NewTag(db *storage.DBService) *Tag {
	return &Tag{db: db}
}

func (u *Tag) Create(tag *model.Tag) (*model.Tag, error) {
	if len(tag.Tag) == 0 {
		return nil, errors.New("Tag cannot be empty")
	}
	if len(tag.Tag) > 40 {
		return nil, errors.New("Maximum tag length is 40 characters")
	}

	tagID, err := u.db.CreateTag(tag.Tag)
	if err != nil {
		return nil, fmt.Errorf("Couldn't insert the tag in the Database: %w", err)
	}

	tag.ID = tagID

	return tag, nil
}

func (u *Tag) Delete(tagID int) (int, error) {
	affected, err := u.db.DeleteTagByID(tagID)
	if err != nil {
		return 0, fmt.Errorf("Couldn't delete tag %w", err)

	}
	return affected, nil
}

func (u *Tag) Search(tag *model.Tag) ([]model.Tag, error) {
	if tag.ID > 0 && len(tag.Tag) > 0 {
		return nil, errors.New(SearchParamError)
	}

	tags, err := u.db.SearchTag(tag.ID, tag.Tag)
	if err != nil {
		return nil, fmt.Errorf("Failed to query the DB: %w", err)
	}

	return tags, nil
}
