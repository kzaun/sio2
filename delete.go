package sio2

import (
	"fmt"
	"log"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/urfave/cli/v2"
)

const (
	DeleteUsage   = "删除数据库中某一条数据，**** d/delete+空格+号码"
	DeleteCommand = "delete"
	DeleteAliases = "d"
)

func NewDeleteCommand(db *leveldb.DB) *cli.Command {
	return &cli.Command{
		Name:    DeleteCommand,
		Aliases: []string{DeleteAliases},
		Usage:   DeleteUsage,
		Action: func(cCtx *cli.Context) error {
			key := strings.Trim(cCtx.Args().First(), " ")
			if ok, _ := db.Has([]byte(key), nil); !ok {
				fmt.Printf("当前号码不存在数据库: %s\n", key)
			} else {
				err := db.Delete([]byte(key), nil)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("当前号码%s删除成功！ \n", key)
			}
			return nil
		},
	}
}
