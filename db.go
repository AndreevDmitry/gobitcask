package gobitcask

import (
	"errors"
	"io"
	"os"
)

type Db struct {
	dir    string
	file   *os.File
	keyDir map[string]int64 //Key -> fileOffset)
}

func New(dir string) *Db {
	keyDir := fillKeydir(dir)
	//Active file
	file, err := os.OpenFile(dir+"/foobar.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	return &Db{
		dir:    dir,
		file:   file,
		keyDir: keyDir,
	}
}

func fillKeydir(dir string) map[string]int64 {
	keyDir := make(map[string]int64)
	file, err := os.Open(dir + "/foobar.db")
	if os.IsNotExist(err) {
		return keyDir
	}

	if err != nil {
		panic(err)
	}
	for {
		offset, _ := file.Seek(0, io.SeekCurrent)
		record, err := Decode(file)
		if err != nil {
			break
		}
		keyDir[record.Key] = offset
	}
	file.Close()
	return keyDir
}

func (db *Db) Put(key string, value string) {
	offset, _ := db.file.Seek(0, io.SeekCurrent)
	db.keyDir[key] = offset
	Encode(db.file, key, value)
}

func (db *Db) Get(key string) (string, error) {
	file, err := os.Open(db.dir + "/foobar.db")
	if err != nil {
		panic(err)
	}
	offset, ok := db.keyDir[key]
	if !ok {
		return "", errors.New("Bitcask: record not found")
	}

	file.Seek(offset, 0)
	record, err := Decode(file)
	if err != nil {
		panic(err)
	}
	return record.Value, nil
}

func (db *Db) Close() {
	db.file.Close()
}
