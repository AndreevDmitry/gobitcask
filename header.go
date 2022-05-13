package gobitcask

import (
	"bytes"
	"encoding/binary"
)

type Record struct {
	Key   string
	Value string
}

func Encode(buffer *bytes.Buffer, key string, value string) {
	keySize := uint64(len(key))
	valueSize := len(value)
	binary.Write(buffer, binary.LittleEndian, keySize)

	buffer.Write([]byte{byte(valueSize)})
	buffer.Write([]byte(key))
	buffer.Write([]byte(value))
}
func Decode(buffer *bytes.Buffer) Record {
	keySize := uint64(0)
	binary.Read(buffer, binary.LittleEndian, &keySize)
	valueSize, err := buffer.ReadByte()
	if err != nil {
		panic(err)
	}
	key := make([]byte, keySize)
	value := make([]byte, valueSize)
	buffer.Read(key)
	buffer.Read(value)

	record := Record{
		Key:   string(key),
		Value: string(value),
	}
	return record
}
