package models

import "time"

type PostInfo struct {
	UID      string   `json:"-" gorm:"primary_key"`
	Hashtags []string `json:"hashtags"`
	Text     string   `json:"text" gorm:"size:300"`
	MediaURI []string `json:"mediaURI"`
	Public   bool     `json:"ispublic"`
}

type Post struct {
	ID        uint      `json:"pid" gorm:"primary_key"`
	Nickname  string    `json:"nick"`
	Timestamp time.Time `json:"timestamp"`
	Hashtags  []string  `json:"hashtags"`
	Text      string    `json:"text" gorm:"size:300"`
	MediaURI  []string  `json:"mediaURI"`
	Likes     uint      `json:"likes"`
}
