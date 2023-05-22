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

func NewPutCommand(*leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    PutCommand,
		Aliases: []string{PutAliases},
		Usage:   PutUsage,
		Action: func(cCtx *cli.Context) error {
			fmt.Println("completed task: ", cCtx.Args().First())
			// err := db.View(func(tx *bolt.Tx) error {

			// 	//3.打开BlockBucket表，获取表对象

			// 	b := tx.Bucket([]byte(TableName))

			// 	//4.Get()方法通过key读取value
			// 	if b != nil {
			// 		data := b.Get([]byte("l"))
			// 		fmt.Printf("%s\n", data)
			// 		data = b.Get([]byte("ll"))
			// 		fmt.Printf("%s\n", data)
			// 	}

			// 	return nil
			// })
			// if err != nil {
			// 	log.Panic(err)
			// }
			return nil
		},
	}
}
