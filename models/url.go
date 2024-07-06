package models

import "time"

type Url struct {
	ID           string     `gorm:"primarykey;autoIncrement:false;type:varchar(20)" json:"id"`
	Url          string     `json:"url" gorm:"type:text"`
	ShortenUrl   string     `json:"shorten_url" gorm:"type:varchar(30)"`
	UsageCount   int        `json:"usage_count" gorm:"default:0"`
	CreatedAt    time.Time  `json:"created_at"`
	LastAccessed *time.Time `json:"last_accessed"`
}
