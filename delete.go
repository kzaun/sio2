package sio2

import (
	"fmt"

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
			fmt.Println("completed task: ", cCtx.Args().First())
			return nil
		},
	}
}
