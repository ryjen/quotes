package storage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuoteSerialize(t *testing.T) {

	input := `
  {
    "id": "1234",
    "text": "testing123",
    "author": "test"
  }
  `

	var quote Quote
	err := json.Unmarshal([]byte(input), &quote)

	assert.NoError(t, err)

	assert.Equal(t, quote.ID, "1234")
}
