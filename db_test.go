package gobitcask

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDb_Get(t *testing.T) {
	t.Run("SimpleGetSet", func(t *testing.T) {
		db := New(t.TempDir())
		db.Put("key1", "value1")
		db.Put("key2", "value2")
		value := db.Get("key1")
		assert.Equal(t, "value1", value)
		value = db.Get("key2")
		assert.Equal(t, "value2", value)
	})

	t.Run("NotFound", func(t *testing.T) {

	})

	t.Run("DoublePut", func(t *testing.T) {
		db := New(t.TempDir())
		db.Put("key1", "value1")
		db.Put("key1", "value2")
		value := db.Get("key1")
		assert.Equal(t, "value2", value)
	})
}
