package sio2

import (
	"bytes"
	"encoding/gob"
	"log"
)

type MetaData struct {
	FileSize int64
	FileName string
	FilePath string
	Created  int64
}

func (meta *MetaData) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(meta)
	if err != nil {
		log.Panic(err)
	}
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

func removeDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
func inArray(n int, f func(int) bool) bool {

	for i := 0; i < n; i++ {
		if f(i) {
			return true
		}
	}
	return false
}
