package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bitfield/script"
	"github.com/urfave/cli"
)

func main() {

	var command string
	var regex string
	var verbose bool

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "regex,r",
			Value:       ".",
			Usage:       "regex for cluster selection",
			Destination: &regex,
		},
		cli.StringFlag{
			Name:        "command,c",
			Value:       "",
			Usage:       "command to run on cluster",
			Destination: &command,
		},
		cli.BoolFlag{
			Name:        "verbose,v",
			Usage:       "print cluster name and command before every output",
			Destination: &verbose,
		},
	}

	app.Action = func(c *cli.Context) error {
		regex := regexp.MustCompile(regex)

		clusters_rune, _ := script.Exec("kubectx").MatchRegexp(regex).String()

		clusters := strings.Split(strings.TrimSpace(clusters_rune), "\n")
		fmt.Println(clusters)

		for _, cluster := range clusters {

			if verbose {
				fmt.Println("[ DEBUG ] CONTEXT: " + cluster)
				fmt.Println("[ DEBUG ] COMMAND: kubectl " + command)

			}

			script.Exec("kubectl " + command + " --context " + cluster).Stdout()
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
