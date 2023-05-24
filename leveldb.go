package sio2

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func NewDB(dbName string) *leveldb.DB {
	db, err := leveldb.OpenFile(dbName, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
