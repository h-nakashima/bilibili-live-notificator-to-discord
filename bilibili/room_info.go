package bilibili

type RoomInfoRequest struct {
	ID int `json:"id"`
}

type RoomInfoResponse struct {
	Data struct {
		RoomID     int    `json:"room_id"`
		LiveStatus int    `json:"live_status"`
		Title      string `json:"title"`
	} `json:"data"`
}
