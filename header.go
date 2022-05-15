package gobitcask

import (
	"encoding/binary"
	"io"
)

type Record struct {
	Key   string
	Value string
}

func Encode(buffer io.Writer, key string, value string) {
	keySize := uint64(len(key))
	valueSize := uint64(len(value))
	binary.Write(buffer, binary.LittleEndian, keySize)
	binary.Write(buffer, binary.LittleEndian, valueSize)
	binary.Write(buffer, binary.LittleEndian, []byte(key))
	binary.Write(buffer, binary.LittleEndian, []byte(value))
}

func Decode(buffer io.Reader) (Record, error) {
	keySize := uint64(0)
	if err := binary.Read(buffer, binary.LittleEndian, &keySize); err != nil {
		return Record{}, err
	}

	valueSize := uint64(0)
	if err := binary.Read(buffer, binary.LittleEndian, &valueSize); err != nil {
		return Record{}, err
	}

	key := make([]byte, keySize)
	if err := binary.Read(buffer, binary.LittleEndian, &key); err != nil {
		return Record{}, err
	}

	value := make([]byte, valueSize)
	if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
		return Record{}, err
	}

	record := Record{
		Key:   string(key),
		Value: string(value),
	}
	return record, nil
}
