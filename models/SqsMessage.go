package models

type SqsMessage struct {
	ImageKey string `json:"image_key"`
	UserID   string `json:"user_id"`
}
