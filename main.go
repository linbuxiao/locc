package main

import (
	"github.com/linbuxiao/locc/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "locc",
		Commands: []*cli.Command{
			cmd.ClockCMD,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
