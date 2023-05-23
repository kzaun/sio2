package sio2

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/urfave/cli/v2"
)

const (
	PutCommand = "put"
	PutAliases = "p"
	PutUsage   = "原数据去重全部写入数据，**** p/put+空格+原文档绝对路径"
)

func NewPutCommand(db *leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    PutCommand,
		Aliases: []string{PutAliases},
		Usage:   PutUsage,
		Action: func(cCtx *cli.Context) error {
			iter := db.NewIterator(nil, nil)
			for iter.Next() {

				key := iter.Key()
				value := iter.Value()
				fmt.Println(string(key))
				fmt.Println(DeserializeBlock(value))
			}
			return nil
		},
	}
}
