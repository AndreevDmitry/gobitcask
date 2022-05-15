package gobitcask

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDb_Get(t *testing.T) {
	t.Run("SimpleGetSet", func(t *testing.T) {
		// Arrange
		db := New(t.TempDir())

		// Act
		db.Put("key1", "value1")
		db.Put("key2", "value2")

		// Assert
		value, _ := db.Get("key1")
		assert.Equal(t, "value1", value)
		value, _ = db.Get("key2")
		assert.Equal(t, "value2", value)
	})

	t.Run("NotFound", func(t *testing.T) {
		// Arrange
		db := New(t.TempDir())

		// Act
		_, err := db.Get("key1")

		// Assert
		assert.EqualError(t, err, "Bitcask: record not found")
	})

	t.Run("DoublePut", func(t *testing.T) {
		db := New(t.TempDir())
		db.Put("key1", "value1")
		db.Put("key1", "value2")
		value, _ := db.Get("key1")
		assert.Equal(t, "value2", value)
	})
	t.Run("OpenCloseOpen", func(t *testing.T) {
		// Arrange
		tempDir := t.TempDir()
		db := New(tempDir)
		db.Put("key1", "value1")
		db.Close()

		// Act
		db = New(tempDir)

		// Assert
		value, err := db.Get("key1")
		assert.Equal(t, "value1", value)
		assert.NoError(t, err)
	})

}
