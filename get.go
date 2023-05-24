package sio2

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/urfave/cli/v2"
)

const (
	GetCommand = "get"
	GetAliases = "g"
	GetUsage   = "查询当前号码创建时间， **** g/get+空格+号码"
)

func NewGetCommand(db *leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    GetCommand,
		Aliases: []string{GetAliases},
		Usage:   GetUsage,
		Action: func(cCtx *cli.Context) error {
			key := strings.Trim(cCtx.Args().First(), " ")
			if ok, _ := db.Has([]byte(key), nil); !ok {
				fmt.Printf("当前号码不存在数据库: %s\n", key)
			} else {
				value, err := db.Get([]byte(key), nil)
				if err != nil {
					log.Fatal(err)
				}
				object := DeserializeBlock(value)
				timeLayout := "2006-01-02 15:04:05"
				timeStr := time.Unix(object.Created, 0).Format(timeLayout)
				fmt.Printf("号码：%s\n", object.Mobile)
				fmt.Printf("上传时间：%s\n", timeStr)
				fmt.Printf("上传完整路径：%s\n", object.FullName)
				fmt.Printf("上传时文件大小：%dKB\n", object.FileSize/1024)
			}
			return nil
		},
	}
}
