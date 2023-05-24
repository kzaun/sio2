package sio2

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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
			fileName := cCtx.Args().First()
			fileData, err := os.Stat(fileName)
			if err != nil {
				log.Print("当前文件不存在")
				return nil
			}

			path, _ := filepath.Split(fileName)
			data, dataErr := ReadExcel(fileName)
			if dataErr != nil {
				log.Print("当前文件读取失败")
				return nil
			}

			numbers := []string{}
			for _, row := range data {
				for _, col := range row {
					numbers = append(numbers, col)
				}
			}
			if len(numbers) == 0 {
				fmt.Println("当前数据空白，检查数据")
				return nil
			}
			sorts := removeDuplication_map(numbers)

			iter := db.NewIterator(nil, nil)
		re:
			for _, v := range sorts {
				for iter.Next() {
					if string(iter.Key()) != v {
						data := &MetaData{
							Created:  time.Now().Unix(),
							FileName: fileData.Name(),
							FilePath: path,
							FileSize: fileData.Size(),
						}
						dataByte := data.Serialize()
						err = db.Put([]byte(v), dataByte, nil)
						if err != nil {
							return nil
						}
						fmt.Printf("号码：%v 成功存储 \n", v)
						break re
					} else {
						fmt.Printf("号码：%v 当前数据已经存在 \n", v)
						break re
					}
				}
			}

			return nil
		},
	}
}
