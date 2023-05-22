package sio2

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type MetaData struct {
	FileSize int
	FileName string
	FilePath string
	Created  time.Duration
}

func (meta *MetaData) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(meta)
	if err != nil {
		log.Panic(err)
	}
	//4.返回result的字节数组
	return result.Bytes()
}

func DeserializeBlock(data []byte) *MetaData {
	var meta MetaData
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&meta)
	if err != nil {
		log.Panic(err)
	}
	return &meta
}
