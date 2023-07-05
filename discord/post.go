package discord

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/xerrors"
)

type Keys struct {
	Token string `yaml:"token"`
	ChannelId string	`yaml:"channelId"`
}

func PostDiscord(keys Keys, message string, title string, roomId int, imageUrl string) error {
	dg, err := discordgo.New("Bot " + keys.Token)
	if err != nil {
		return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
	}

	buf, err := getCoverImage(imageUrl)
	if err != nil {
		_, err = dg.ChannelMessageSend(keys.ChannelId, message+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId))
		if err != nil {
			return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
		} else {
			return nil
		}
	} else {
		_, err = dg.ChannelFileSend(keys.ChannelId, "tweet_image", buf)
		if err != nil {
			return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
		}

		_, err = dg.ChannelMessageSend(keys.ChannelId, message+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId))
		if err != nil {
			return xerrors.Errorf("discord.PostDiscord error occurred: %w", err)
		} else {
			return nil
		}
	}
}

func getCoverImage(imageUrl string) (*bytes.Buffer, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, xerrors.Errorf("discord.getCoverImage error occurred: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("discord.getCoverImage error occurred: %w", err)
	}

	buf := bytes.NewBuffer(data)
	return buf, nil
}