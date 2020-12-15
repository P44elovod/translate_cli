package app

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func Start() error {

	app := cli.NewApp()
	app.Name = "Translation CLI"
	app.Usage = "To Piglatina translation and vowels to digits(vice versa) shifting"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "text",
			Usage: "source text",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "vdencrypt",
			Usage:     "Encrypt vowels to digits",
			UsageText: "Vowels to digits encryption",
			Flags:     flags,
			Action: func(c *cli.Context) error {
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, c.String("text"), "result text")
				return nil
			},
		},
		{
			Name:      "dvdecrypt",
			Usage:     "Decrypt vowels to digits",
			UsageText: "Digits to vowels dencryption",
			Flags:     flags,
			Action: func(c *cli.Context) error {
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, c.String("text"), "result text")
				return nil
			},
		},
		{
			Name:      "piglatina",
			Usage:     "Translate English to Piglatina",
			UsageText: "English to Piglatina translation",
			Flags:     flags,
			Action: func(c *cli.Context) error {
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, c.String("text"), "result text")
				return nil
			},
		},
	}

	return app.Run(os.Args)

}
