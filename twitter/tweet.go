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

func PostTweet(keys Keys, title string, room_id int, imageUrl string) error {
	config := oauth1.NewConfig(keys.Consumer.Key, keys.Consumer.Secret)
	token := oauth1.NewToken(keys.Access.Key, keys.Access.Secret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	buf, err := convertBuffer(imageUrl)
	if err != nil {
		return err
	}
	media, _, err := client.Media.Upload(buf.Bytes(), "tweet_image")
	if err != nil {
		return err
	}

	_, _, err = client.Statuses.Update("Starting live streaming: \n"+title+" https://live.bilibili.com/"+strconv.Itoa(room_id), &twitter.StatusUpdateParams{
		MediaIds: []int64{media.MediaID},
	})
	if err != nil {
		return err
	}

	return nil
}

func convertBuffer(imageUrl string) (*bytes.Buffer, error) {
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