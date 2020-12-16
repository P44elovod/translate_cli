package app

import (
	"fmt"
	"os"
	"strings"
	"translate_cli/util"

	"github.com/urfave/cli"
)

func Start() error {

	app := cli.NewApp()
	app.Name = "Translation CLI"
	app.Usage = "To Piglatin translation and vowels to digits(vice versa) shifting"

	app.Commands = []cli.Command{
		{
			Name:      "encrypt",
			Usage:     "Encrypt vowels to digits",
			UsageText: "Vowels to digits encryption",
			Action: func(c *cli.Context) error {
				source := strings.Join(c.Args().Tail(), " ")
				result := util.ShiftLettersToDigits(source)
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, source, result)
				return nil
			},
		},
		{
			Name:      "decrypt",
			Usage:     "Decrypt vowels to digits",
			UsageText: "Digits to vowels dencryption",
			Action: func(c *cli.Context) error {
				source := strings.Join(c.Args().Tail(), " ")
				result := util.ShiftDigitsToLetters(source)
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, source, result)
				return nil
			},
		},
		{
			Name:      "piglatin",
			Usage:     "Translate English to Piglatin",
			UsageText: "English to Piglatina translation",
			Action: func(c *cli.Context) error {
				var source string
				if c.NArg() > 0 {
					source = strings.Join(c.Args().Tail(), " ")
				}
				result := util.TranslateTextToPiglatin(source)
				fmt.Printf("%s: \n source: %s\n result: %s\n", c.Command.UsageText, source, result)
				return nil
			},
		},
	}

	return app.Run(os.Args)

}
