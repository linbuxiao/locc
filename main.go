package main

import (
	"github.com/linbuxiao/locc/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "locc",
		Usage: "Just in local.",
		Commands: []*cli.Command{
			cmd.ClockCMD,
			cmd.MemoCMD,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
