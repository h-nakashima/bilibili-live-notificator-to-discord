package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"bilibili-live-notificator/bilibili"
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
		roomInfo, err := bilibili.GetRoomInfo(c.String("room-id"))
		if err != nil {
			return err
		}
		fmt.Println(roomInfo.Title)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
