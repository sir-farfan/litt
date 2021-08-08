package storage

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/codelittinc/golang-interview-project-ismael-estrada/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	createTag = fmt.Sprintf("create%d", rand.Intn(100000))
	searchTag = fmt.Sprintf("search%d", rand.Intn(100000))
	deleteTag = fmt.Sprintf("delete%d", rand.Intn(100000))
}

var (
	createTag string
	searchTag string
	deleteTag string
)

func getConnection() *sql.DB {
	params := ConnectionParams{
		Driver: "mysql",
		User:   "root",
		Pass:   "YOUR_PASSWORD_HERE",
		// Server: "localhost",
		DB: "litt",
	}

	return ConnectToDB(params)
}

func TestCreateTag(t *testing.T) {
	db := getConnection()
	defer db.Close()

	service := New(db)

	tests := []struct {
		name     string
		tag      string
		inserted bool
		errStr   string
	}{
		{
			name:     "successful create",
			tag:      createTag,
			inserted: true,
		},
		{
			name:     "duplicated tag fails",
			tag:      createTag,
			inserted: false,
			errStr:   "Error 1062: Duplicate entry '" + createTag + "' for key 'tag.tag'",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inserted, err := service.CreateTag(tc.tag)
			assert.Equal(t, inserted > 0, tc.inserted)
			if len(tc.errStr) > 0 {
				assert.Equal(t, tc.errStr, err.Error())
			}
		})
	}
}

func TestDeleteTag(t *testing.T) {
	db := getConnection()
	defer db.Close()

	service := New(db)
	service.CreateTag(deleteTag)

	tests := []struct {
		name  string
		tag   string
		count int
		err   error
	}{
		{
			name:  "existing tag",
			tag:   deleteTag,
			count: 1,
		},
		{
			name:  "non existing",
			tag:   deleteTag + createTag + searchTag,
			count: 0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			affected, err := service.DeleteTag(tc.tag)
			assert.Equal(t, affected, tc.count)
			assert.Equal(t, err, tc.err)
		})
	}
}

func TestSearchTag(t *testing.T) {
	db := getConnection()
	defer db.Close()

	service := New(db)
	service.CreateTag(searchTag)

	tests := []struct {
		name       string
		tag2search string
		tags       []model.Tag
		count      int
	}{
		{
			name:       "found a tag",
			tag2search: searchTag,
			count:      1,
		},
		{
			name:       "tag not found",
			tag2search: createTag + searchTag + deleteTag,
			count:      0,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tags, _ := service.SearchTag(tc.tag2search)
			assert.Equal(t, tc.count, len(tags))
			if tc.count > 0 {
				assert.Equal(t, tags[0].Tag, tc.tag2search)
			}
		})
	}
}
