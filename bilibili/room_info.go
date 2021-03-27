package bilibili

import (
	"bilibili-live-notificator/client"
	"context"
	"net/http"
	"time"

	"golang.org/x/xerrors"
)

type RoomInfo struct {
	RoomID     *int    `json:"room_id"`
	LiveStatus *int    `json:"live_status,omitempty"`
	Title      *string `json:"title"`
	ImageUrl   *string `json:"user_cover"`
}

type RoomInfoRequest struct {
	ID int `json:"id"`
}

type RoomInfoResponse struct {
	Data RoomInfo `json:"data"`
}

// TODO: テストコードを書く
func GetRoomInfo(id string) (*RoomInfo, error) {
	client, _ := client.NewClient(
		"https://api.live.bilibili.com",
		&http.Client{},
		"string", // TODO: UAを設定する
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	httpRequest, err := client.NewRequest(ctx, "GET", "/room/v1/Room/get_info", "id="+id, nil)
	if err != nil {
		return nil, xerrors.Errorf("failed to create new request: %w", err)
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, xerrors.Errorf("failed to do request: %w", err)
	}

	var apiResponse RoomInfoResponse
	if err := client.DecodeBody(httpResponse, &apiResponse); err != nil {
		return nil, xerrors.Errorf("failed to decode room info response: %w", err)
	}

	return &apiResponse.Data, nil
}
