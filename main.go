package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"bilibili-live-notificator/bilibili"
	"bilibili-live-notificator/discord"

	"gopkg.in/yaml.v2"
)

type config struct {
	Discord discord.Keys `yaml:"discord"`
}

func main() {
	// TODO: Add test code

	app := &cli.App{
		Name:    "bilibili-live-notificator",
		Usage:   "It detects starting the live streaming on Bilibili and notifies a Discord.",
		Version: "0.0.5",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "room-id",
				Usage:    "Bilibili room ID",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "api-keys-file",
				Usage:    "API keys file",
				Aliases:  []string{"k"},
				Required: true,
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		// Load API keys file
		file, err := ioutil.ReadFile(c.String("api-keys-file"))
		if err != nil {
			return xerrors.Errorf("failed to read API keys file: %w", err)
		}
		var config config
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			return xerrors.Errorf("failed to parse API keys file: %w", err)
		}

		roomInfo, err := bilibili.GetRoomInfo(c.String("room-id"))
		liveStatus := *roomInfo.LiveStatus
		if err != nil {
			return xerrors.Errorf("failed to get live-status from bilibili: %w", err)
		}

		for {
			sleepingTime := 5 + rand.Intn(5)
			log.Println("Sleep " + strconv.Itoa(sleepingTime) + "sec")
			time.Sleep(time.Duration(sleepingTime) * time.Second)

			roomInfo, err := bilibili.GetRoomInfo(c.String("room-id"))
			if err != nil {
				log.Printf("%+v\n", err)
			} else {
				if liveStatus == *roomInfo.LiveStatus {
					if *roomInfo.LiveStatus == 1 {
						err = discord.PostDiscord(config.Discord, "Started live streaming at "+time.Now().Format("15:04:05 MST: "), *roomInfo.Title, *roomInfo.RoomID, *roomInfo.ImageUrl)
					} else {
						err = discord.PostDiscord(config.Discord, "Finished live streaming at "+time.Now().Format("15:04:05 MST: "), *roomInfo.Title, *roomInfo.RoomID, *roomInfo.ImageUrl)
					}
					if err != nil {
						log.Printf("%+v\n", err)
					}
				}
				liveStatus = *roomInfo.LiveStatus
			}
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
