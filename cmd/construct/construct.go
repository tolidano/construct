package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tolidano/construct"
	"github.com/urfave/cli"
)

// App set-up
func init() {

}

// Command-line utility entry point.
func main() {
	app := cli.NewApp()
	app.Name = "Construct"
	app.Usage = "Cloud platform administration"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Usage: "Load YAML configuration from `FILE`",
		},
		cli.StringFlag{
			Name:  "command, d",
			Usage: "Execute a `COMMAND`",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Verbose output",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("conf") == "" {
			fmt.Println("You must specify a configuration file.")
			os.Exit(1)
		} else {
			yamlConfig := construct.ParseFile(c.String("conf"))
			if c.String("command") != "" {
				// Execute command
			} else {
				// Show config
				fmt.Printf("%+v\n", yamlConfig)
			}
		}
		return nil
	}

	app.Run(os.Args)
}
