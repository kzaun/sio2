package sio2

import (
	"fmt"

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
			err := db.Put([]byte("key"), []byte("value"), nil)
			if err != nil {
				fmt.Println("111")
			}
			result, _ := db.Get([]byte("key"), nil)
			fmt.Println(string(result))
			fmt.Println("completed task: ", cCtx.Args().First())
			return nil
		},
	}
}
