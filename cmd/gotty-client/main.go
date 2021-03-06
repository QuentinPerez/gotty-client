package main

import (
	"os"
	"path"

	"github.com/moul/gotty-client"
	"github.com/moul/gotty-client/vendor/github.com/Sirupsen/logrus"
	"github.com/moul/gotty-client/vendor/github.com/codegangsta/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/gotty-client"
	app.Version = VERSION
	app.Usage = "GoTTY client for your terminal"
	app.ArgsUsage = "GOTTY_URL"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, D",
			Usage:  "Enable debug mode",
			EnvVar: "GOTTY_CLIENT_DEBUG",
		},
	}

	app.Action = Action

	app.Run(os.Args)
}

func Action(c *cli.Context) {
	if len(c.Args()) != 1 {
		logrus.Fatalf("usage: gotty-client [GoTTY URL]")
	}

	// setting up logrus
	logrus.SetOutput(os.Stderr)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// create Client
	url := c.Args()[0]
	client, err := gottyclient.NewClient(url)
	if err != nil {
		logrus.Fatalf("Cannot create client: %v", err)
	}

	// loop
	if err = client.Loop(); err != nil {
		logrus.Fatalf("Communication error: %v", err)
	}
}
