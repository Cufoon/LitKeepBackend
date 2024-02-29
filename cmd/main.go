package main

import (
	"log"
	"os"

	"cufoon.litkeep.service/app"
	"github.com/urfave/cli/v2"
)

func main() {
	println("<-----Litkeep for Ms. Tang!----->")
	clic := &cli.App{
		Name:  "Litkeep Backend Service",
		Usage: "Just record your life bill!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./dev.yaml",
				Usage:   "config file to run the backend service",
				EnvVars: []string{"LITKEEP_CONFIG"},
			},
		},
		Action: func(cCtx *cli.Context) error {
			cf := cCtx.String("config")
			app.StartAPP(cf)
			return nil
		},
	}
	if err := clic.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
