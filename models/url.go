package models

import "time"

type Url struct {
	ID         string    `gorm:"primarykey,autoIncrement:false" json:"id"`
	Url        string    `json:"url" gorm:"unique"`
	ShortenUrl string    `json:"shorten_url"`
	CreatedAt  time.Time `json:"created_at"`
}
