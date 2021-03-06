package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	"bilibili-live-notificator/bilibili"
	"bilibili-live-notificator/twitter"

	"gopkg.in/yaml.v2"
)

type config struct {
	Twitter twitter.Keys `yaml:"twitter"`
}

func main() {
	// TODO: Use xerrors
	// TODO: Add test code

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
			panic(err)
		}
		var config config
		err = yaml.Unmarshal(file, &config)

		liveStatus := -1

		for {
			roomInfo, err := bilibili.GetRoomInfo(c.String("room-id"))
			if err != nil {
				return err
			}
			if liveStatus != *roomInfo.LiveStatus {
				if *roomInfo.LiveStatus == 1 {
					err = twitter.PostTweet(config.Twitter, "Started live streaming: ", *roomInfo.Title, *roomInfo.RoomID, *roomInfo.ImageUrl)
				} else {
					err = twitter.PostTweet(config.Twitter, "Finished live streaming: ", *roomInfo.Title, *roomInfo.RoomID, *roomInfo.ImageUrl)
				}
				if err != nil {
					return err
				}
			}
			liveStatus = *roomInfo.LiveStatus
			sleepingTime := 5 + rand.Intn(5)
			log.Println("Sleep " + strconv.Itoa(sleepingTime) + "sec")
			time.Sleep(time.Duration(sleepingTime) * time.Second)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
