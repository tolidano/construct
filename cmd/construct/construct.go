package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tolidano/construct"
	"github.com/urfave/cli"
)

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
			Name: "verbose, v",
			Usage: "Verbose output",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("conf") == "" {
			fmt.Println("You must specify a configuration file.")
			os.Exit(1)
		} else {
            yamlConfig := construct.ParseFile(c.String("conf"))
			box := lockbox.New(yamlConfig)
			if c.Bool("list") {
				// List secrets
				secrets, err := box.All()
				if err != nil {
					fmt.Println(err)
				} else {
					for _, secret := range secrets {
						fmt.Println(*secret)
					}
				}
			} else if c.String("value") != "" {
				// Set secret
				box.Set(c.String("section"), c.String("secret"), c.String("value"))
				fmt.Println("Saved")
			} else if c.String("secret") != "" {
				// Get secret
				fmt.Println(box.Get(c.String("section"), c.String("secret")))
			} else {
				// Show config
				fmt.Printf("%+v\n", yamlConfig)
			}
		}
		return nil
	}

	app.Run(os.Args)
}
