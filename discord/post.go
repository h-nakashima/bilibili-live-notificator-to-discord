package discord

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/xerrors"
)

type Keys struct {
	Token string `yaml:"token"`
	ChannelId string	`yaml:"channelId"`
}

// PostDiscord sends a message to a Discord channel with an optional image.
func PostDiscord(keys Keys, message string, title string, roomId int, imageUrl string) error {
	dg, err := discordgo.New("Bot " + keys.Token)
	if err != nil {
		return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
	}

	reader, fileName, err := getImageFromURL(imageUrl)
	if err != nil {
		_, err = dg.ChannelMessageSend(keys.ChannelId, message+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId))
		if err != nil {
			return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
		} else {
			return nil
		}
	} else {
		defer reader.Close()

		// Prepare the message with the image and text.
		messageFile := &discordgo.File{
			Name:   fileName,
			Reader: reader,
		}

		messageSend := &discordgo.MessageSend{
			Content: message+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId),
			Files:   []*discordgo.File{messageFile},
		}

		// Send the message.
		_, err = dg.ChannelMessageSendComplex(keys.ChannelId, messageSend)
		if err != nil {
			return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
		} else {
			return nil
		}
	}
}

// getImageFromURL retrieves an image from a URL and returns the image data and filename.
func getImageFromURL(imageURL string) (io.ReadCloser, string, error) {
    // Create an HTTP request and retrieve the image data.
    response, err := http.Get(imageURL)
    if err != nil {
        return nil, "", err
    }

    if response.StatusCode != 200 {
        return nil, "", fmt.Errorf("failed to get image from URL: HTTP %v", response.StatusCode)
    }

    // Extract the filename from the URL.
    fileName := path.Base(response.Request.URL.Path)

    return response.Body, fileName, nil
}
