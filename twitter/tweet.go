package twitter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Keys struct {
	Consumer struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"consumer"`
	Access struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"access"`
}

func PostTweet(keys Keys, tweetMessage string, title string, roomId int, imageUrl string) error {
	config := oauth1.NewConfig(keys.Consumer.Key, keys.Consumer.Secret)
	token := oauth1.NewToken(keys.Access.Key, keys.Access.Secret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	buf, err := getCoverImage(imageUrl)
	if err != nil {
		_, _, err = client.Statuses.Update(tweetMessage+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId), nil)
		if err != nil {
			return err
		}
	} else {
		media, _, err := client.Media.Upload(buf.Bytes(), "tweet_image")
		if err != nil {
			return err
		}

		_, _, err = client.Statuses.Update(tweetMessage+"\n"+title+" https://live.bilibili.com/"+strconv.Itoa(roomId), &twitter.StatusUpdateParams{
			MediaIds: []int64{media.MediaID},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getCoverImage(imageUrl string) (*bytes.Buffer, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data)
	return buf, nil
}
