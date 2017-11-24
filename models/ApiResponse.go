package models

// Structure to respond to api
type ApiResponse struct {
	Status   string `json:"status"`
	ImageUID string `json:"image_uid"`
	Token    string `json:"token"`
}

type ApiRequest struct {
	ImageUID string `json:"image_uid"`
}
type Result struct {
	ImageUID string  `json:"image_uid"`
	Status   string  `json:"status"`
	Type1    float32 `json:"type1"`
	Type2    float32 `json:"type2"`
	Type3    float32 `json:"type3"`
}

type Message struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}
