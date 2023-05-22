package sio2

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Read(cCtx *cli.Context) error {
	fmt.Println("added task: ", cCtx.Args().First())
	return nil
}

func Put(cCtx *cli.Context) error {
	fmt.Println("added task: ", cCtx.Args().First())
	return nil
}
