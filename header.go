package gobitcask

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Record struct {
	Key   string
	Value string
}

func Encode(buffer *bytes.Buffer, key string, value string) {
	keySize := uint64(len(key))
	valueSize := uint64(len(value))
	binary.Write(buffer, binary.LittleEndian, keySize)
	binary.Write(buffer, binary.LittleEndian, valueSize)
	binary.Write(buffer, binary.LittleEndian, []byte(key))
	binary.Write(buffer, binary.LittleEndian, []byte(value))
}
func Decode(buffer *bytes.Buffer) Record {
	keySize := uint64(0)
	if err := binary.Read(buffer, binary.LittleEndian, &keySize); err != nil {
		fmt.Println("keySize binary.Read failed:", err)
	}

	valueSize := uint64(0)
	if err := binary.Read(buffer, binary.LittleEndian, &valueSize); err != nil {
		fmt.Println("valueSize binary.Read failed:", err)
	}

	key := make([]byte, keySize)
	if err := binary.Read(buffer, binary.LittleEndian, &key); err != nil {
		fmt.Println("key binary.Read failed:", err)
	}

	value := make([]byte, valueSize)
	if err := binary.Read(buffer, binary.LittleEndian, &value); err != nil {
		fmt.Println("value binary.Read failed:", err)
	}

	record := Record{
		Key:   string(key),
		Value: string(value),
	}
	return record
}
