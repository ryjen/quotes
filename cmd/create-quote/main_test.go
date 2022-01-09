package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"micrantha.com/quotes/internal/storage"
)

func init() {
	db = storage.NewMemoryDB()
}

func TestHandleEvents(t *testing.T) {

	res, err := createQuote(context.TODO(), storage.NewQuote{
		Text:   "testing handle events",
		Author: "testing",
	})

	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	var value map[string]string
	err = json.Unmarshal([]byte(res.Body), &value)
	assert.NoError(t, err)

	id := value["id"]
	assert.NotNil(t, id)

	saved, err := db.GetQuote(context.TODO(), id)

	assert.NoError(t, err)
	assert.NotNil(t, saved)

	assert.Equal(t, saved.Text, "testing handle events")
	assert.Equal(t, saved.Author, "testing")
}
