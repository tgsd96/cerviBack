package models

// Structure to respond to api
type ApiResponse struct {
	Status   string `json:"status"`
	ImageUID string `json:"image_uid"`
	Token    string `json:"token"`
}
