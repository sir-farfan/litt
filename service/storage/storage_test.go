package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	params := ConnectionParams{
		Driver: "mysql",
		User:   "root",
		Pass:   "YOUR_PASSWORD_HERE",
		// Server: "localhost",
		DB: "litt",
	}

	db := ConnectToDB(params)
	defer db.Close()

	service := New(db)

	tests := []struct {
		name     string
		tag      string
		inserted int
		errStr   string
	}{
		{
			name:     "successful create",
			tag:      "first",
			inserted: 1,
		},
		{
			name:     "duplicated tag fails",
			tag:      "first",
			inserted: 0,
			errStr:   "Error 1062: Duplicate entry 'first' for key 'tag.tag'",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inserted, err := service.CreateTag(tc.tag)
			assert.Equal(t, inserted, tc.inserted)
			if len(tc.errStr) > 0 {
				assert.Equal(t, tc.errStr, err.Error())
			}
		})
	}
}
