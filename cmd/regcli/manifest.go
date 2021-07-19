package main

import (
	"github.com/docker/distribution/manifest/schema2"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

var manifest = &cli.Command{
	Name:    "manifest",
	Aliases: []string{"m"},
	Usage:   "manifest's api",
	Subcommands: []*cli.Command{
		pull,
		push,
	},
}

var pull = &cli.Command{
	Name:  "pull",
	Usage: "pull a manifest",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "repository",
			Aliases:  []string{"r"},
			Usage:    "the name of repository",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "reference",
			Aliases:  []string{"t"},
			Value:    "latest",
			Usage:    "reference",
			Required: true,
		},
	},
	Action: func(context *cli.Context) (err error) {
		repository := context.String("repository")
		reference := context.String("reference")
		manifest, err := Registry.GetManifest(repository, reference)
		if err != nil {
			return
		}
		man, err := schema2.FromStruct(manifest)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		marshaled, err := man.MarshalJSON()
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		log.Logger.Info("\n", string(marshaled))
		return
	},
}

var push = &cli.Command{
	Name:  "push",
	Usage: "upload a manifest",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "repository",
			Aliases:  []string{"r"},
			Usage:    "the name of repository",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "reference",
			Aliases:  []string{"t"},
			Value:    "latest",
			Usage:    "reference",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "manifest_file",
			Aliases:  []string{"m", "manifest"},
			Usage:    "manifest",
			Required: true,
		},
	},
	Action: func(context *cli.Context) (err error) {
		repository := context.String("repository")
		reference := context.String("reference")
		manifestFile := context.String("manifest_file")
		content, err := ioutil.ReadFile(manifestFile)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		var manifest schema2.DeserializedManifest
		err = manifest.UnmarshalJSON(content)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		err = Registry.PutManifest(repository, reference, manifest.Manifest)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		return
	},
}
