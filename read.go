package sio2

import (
	"fmt"

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
			fmt.Println("completed task: ", cCtx.Args().First())
			return nil
		},
	}
}
