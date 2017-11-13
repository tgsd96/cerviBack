package models

import "github.com/jinzhu/gorm"

type ImageStatus struct {
	gorm.Model
	ImageKey string  `json:"image_key" sql:"type:VARCHAR(255)"`
	UserID   string  `json:"user_id" sql:"type:VARCHAR(255)"`
	Status   string  `json:"status" sql:"type:VARCHAR(255)"`
	Type1    float32 `json:"type_1" sql:"type:DECIMAL(3,3)"`
	Type2    float32 `json:"type_2" sql:"type:DECIMAL(3,3)"`
	Type3    float32 `json:"type_3" sql:"type:DECIMAL(3,3)"`
}
