package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const usage = `A cli for registry v2 api`

func main() {
	app := &cli.App{
		Name:  "ctrsploit",
		Usage: usage,
		Commands: []*cli.Command{

		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
