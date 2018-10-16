package model

type User struct {
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type Video struct {
	VideoName string `json:"video_name"`
	VideoAuthor string `json:"video_author"`
	VideoID string `json:"video_id"`
}
