package sio2

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/urfave/cli/v2"
)

const (
	ReadCommand = "read"
	ReadAliases = "r"
	ReadUsage   = "读取Excel数据，标注与数据库重复项 并生成u开头的文档，**** r/read+空格+文档绝对路径"
)

func NewReadCommand(db *leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    ReadCommand,
		Aliases: []string{ReadAliases},
		Usage:   ReadUsage,
		Action: func(cCtx *cli.Context) error {
			// err := db.Update(func(tx *bolt.Tx) error {

			// 	//2.通过Bucket()方法打开BlockBucket表
			// 	b := tx.Bucket([]byte(TableName))
			// 	//3.通过Put()方法往表里面存储数据
			// 	if b != nil {
			// 		err := b.Put([]byte("l"), []byte("Send $100 TO Bruce"))
			// 		if err != nil {
			// 			log.Panic("数据存储失败..")
			// 		}
			// 	}
			// 	return nil
			// })

			// err = db.View(func(tx *bolt.Tx) error {

			// 	//3.打开BlockBucket表，获取表对象

			// 	b := tx.Bucket([]byte(TableName))
			// 	//4.Get()方法通过key读取value
			// 	if b != nil {
			// 		data := b.Get([]byte("l"))
			// 		fmt.Printf("%s\n", data)
			// 		data = b.Get([]byte("ll"))
			// 		fmt.Printf("%s\n", data)
			// 	}
			// 	// fmt.Println(b.Get([]byte("ll")))

			// 	return nil
			// })
			// if err != nil {
			// 	log.Panic(err)
			// }
			// fmt.Println("completed task: ", cCtx.Args().First())
			return nil
		},
	}
}
