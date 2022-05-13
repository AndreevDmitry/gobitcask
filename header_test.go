package gobitcask

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	t.Run("Simple case", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Encode(buffer, "key1", "value1")
		result := Decode(buffer)
		assert.Equal(t, "key1", result.Key)
		assert.Equal(t, "value1", result.Value)
	})

	t.Run("Long key", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		key := strings.Repeat("x", 300)
		Encode(buffer, key, "value1")
		result := Decode(buffer)
		assert.Equal(t, key, result.Key)
		assert.Equal(t, "value1", result.Value)
	})
}
