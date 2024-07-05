package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Url        string `json:"url" gorm:"unique"`
	ShortenUrl string `json:"shorten_url"`
}
