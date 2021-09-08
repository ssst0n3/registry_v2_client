package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/registry"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `A cli for registry v2 api`

var Registry registry.Registry

func main() {
	app := &cli.App{
		Name:  "reg cli",
		Usage: usage,
		Commands: []*cli.Command{
			manifest,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "registry_address",
				Aliases:  []string{"r"},
				Required: true,
				Usage:    "the address of registry",
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Usage:   "username",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "password",
			},
			&cli.BoolFlag{
				Name:    "insecure",
				Aliases: []string{"k"},
				Value:   true,
				Usage:   "insecure",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Aliases: []string{"d"},
				Value: false,
				Usage: "debug",
			},
		},
		Before: func(context *cli.Context) (err error) {
			debug := context.Bool("debug")
			if debug {
				log.Logger.Level = logrus.DebugLevel
			} else {
				log.Logger.Level = logrus.InfoLevel
				log.Logger.SetReportCaller(false)
			}
			address := context.String("registry_address")
			username := context.String("username")
			password := context.String("password")
			insecure := context.Bool("insecure")
			Registry = registry.NewRegistry(address, username, password, insecure)
			return
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Logger.Fatal(err)
	}
}
