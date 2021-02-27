package bilibili

import (
	"bilibili-live-notificator/client"
	"context"
	"net/http"
	"time"
)

type RoomInfo struct {
	RoomID     int    `json:"room_id"`
	LiveStatus int    `json:"live_status"`
	Title      string `json:"title"`
}

type RoomInfoRequest struct {
	ID int `json:"id"`
}

type RoomInfoResponse struct {
	Data RoomInfo `json:"data"`
}

func GetRoomInfo(id string) (*RoomInfo, error) {
	client, _ := client.NewClient(
		"https://api.live.bilibili.com",
		&http.Client{},
		"string",
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	httpRequest, err := client.NewRequest(ctx, "GET", "/room/v1/Room/get_info", "id="+id, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse RoomInfoResponse
	if err := client.DecodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse.Data, nil
}
