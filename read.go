package sio2

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/urfave/cli/v2"
)

const (
	ReadCommand = "read"
	ReadAliases = "r"
	ReadUsage   = "读取Excel数据，标注与数据库重复项 并生成u开头的文档，**** r/read+空格+文档绝对路径"
)

// 这边主要有两个逻辑
// 1.读取当前传入文档，遍历固定格式电话号码并以此为key向数据库查询，查到则生成第三行式批注，否则继续查
// 2.生成 不一样色块的excel以及包含元数据批注并导出新的文档
func NewReadCommand(db *leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    ReadCommand,
		Aliases: []string{ReadAliases},
		Usage:   ReadUsage,
		Action: func(cCtx *cli.Context) error {
			fileName := cCtx.Args().First()
			_, err := os.Stat(fileName)
			if err != nil {
				log.Print("当前文件不存在")
				return nil
			}

			data, dataErr := ReadExcel(fileName)
			if dataErr != nil {
				log.Print("当前文件读取失败")
				return nil
			}
			iter := db.NewIterator(nil, nil)

			for _, row := range data {
				for _, colCell := range row {
					for iter.Next() {

						if string(iter.Key()) != colCell {
							continue
						}
						key := iter.Key()
						value := iter.Value()
						fmt.Println(string(key))
						fmt.Println(string(value))
					}

					fmt.Print(colCell, "\t")
				}
				fmt.Println()
			}
			fmt.Println("completed task: ", fileName)
			return nil
		},
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
