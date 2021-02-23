package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "bilibili-live-notificator",
		Usage:   "It detects starting the live streaming on Bilibili and notifies a Twitter.",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "room-id",
				Usage:    "Bilibili room ID",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "twitter-api-key",
				Usage:    "Twitter API key",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "watching",
				Aliases: []string{"w"},
				Usage:   "Service mode",
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("B roomId = ", c.String("room-id"))
		fmt.Println("twitter-api-key = ", c.String("twitter-api-key"))
		fmt.Println("watching", c.String("watching"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
