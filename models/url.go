package models

import "time"

type Url struct {
	ID         string    `gorm:"primarykey;autoIncrement:false;type:varchar(20)" json:"id"`
	Url        string    `json:"url" gorm:"type:text"`
	ShortenUrl string    `json:"shorten_url" gorm:"type:varchar(30)"`
	CreatedAt  time.Time `json:"created_at"`
}
