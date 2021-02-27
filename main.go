package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"bilibili-live-notificator/bilibili"
	"bilibili-live-notificator/client"
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
		client, _ := client.NewClient(
			"https://api.live.bilibili.com/room/v1/Room/get_info",
			&http.Client{},
			"string",
		)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		httpRequest, _ := client.NewRequest(ctx, "GET", "?id=22159299&from=room", nil)
		httpResponse, _ := client.HTTPClient.Do(httpRequest)
		io.Copy(os.Stdout, httpResponse.Body)
		var apiResponse bilibili.RoomInfoResponse
		if err := client.DecodeBody(httpResponse, &apiResponse); err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
