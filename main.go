package main

import (
	"os"

	"github.com/HenriqueOtsuka/goexpert-challenge-1/cmd"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/cfg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "goexpert-challenge-1"
	app.Usage = "Challenge 1 for GoExpert"
	app.Flags = cfg.Flags
	app.Action = cli.ActionFunc(run)
	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal("failed to run CLI app")
	}
	logrus.Info("Shutting down...")
}

func run(c *cli.Context) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(cfg.App.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse log level")
	}
	logrus.SetLevel(level)
	logrus.Info("Starting application...")
	router := cmd.SetupRouter()
	logrus.Info("Router setup complete")

	err = router.Run(":8080")
	if err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}

	return err
}
