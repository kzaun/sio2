package main

import (
	"log"
	"os"

	"github.com/aszeta/sio2"
	"github.com/urfave/cli/v2"
)

const DbName = "sio2.db"

func main() {
	db := sio2.NewDB(DbName)
	commands := []*cli.Command{
		sio2.NewReadCommand(db),
		sio2.NewGetCommand(db),
		sio2.NewDeleteCommand(db),
		sio2.NewPutCommand(db),
	}

	app := &cli.App{
		EnableBashCompletion: true,
		Commands:             commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
