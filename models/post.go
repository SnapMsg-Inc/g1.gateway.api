package models

import "time"

type Post struct {
	ID        uint      `json:"pid" gorm:"primary_key"`
	UID       string    `json:"uid"`
	Nickname  string    `json:"nick"`
	Timestamp time.Time `json:"timestamp"`
	Hashtags  []string  `json:"hashtags"`
	Text      string    `json:"text" gorm:"size:300"`
	MediaURI  []string  `json:"media_uri"`
	Likes     uint      `json:"likes"`
}

type PostCreate struct {
	Hashtags []string `json:"hashtags"`
	Text     string   `json:"text" gorm:"size:300"`
	MediaURI []string `json:"mediaURI"`
	Private bool      `json:"is_private"`
}


