package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("chainwatch")

func main() {
	if err := logging.SetLogLevel("*", "info"); err != nil {
		log.Fatal(err)
	}
	log.Info("Starting chainwatch")

	app := &cli.App{
		Name:    "lotus-chainwatch",
		Usage:   "Devnet token distribution utility",
		Version: build.UserVersion(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "db",
				EnvVars: []string{"LOTUS_DB"},
				Value:   "postgres://postgres:password@192.168.168.10:5432/postgres?sslmode=disable",
			},
			&cli.StringFlag{
				Name:    "log-level",
				EnvVars: []string{"GOLOG_LOG_LEVEL"},
				Value:   "debug",
			},
		},
		Commands: []*cli.Command{
			dotCmd,
			runCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
