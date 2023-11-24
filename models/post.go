package models

import "time"

type Post struct {
	ID        string    `json:"pid" gorm:"primary_key"`
	UID       string    `json:"uid"`
	Timestamp time.Time `json:"timestamp"`
	Hashtags  []string  `json:"hashtags"`
	Text      string    `json:"text" gorm:"size:300"`
	MediaURI  []string  `json:"media_uri"`
	Likes     uint      `json:"likes"`
}

type PostCreate struct {
	UID      string   `json:"uid" swaggerignore:"true"`
	Text     string   `json:"text" gorm:"size:300"`
	Hashtags []string `json:"hashtags,omitempty"`
	MediaURI []string `json:"media_uri,omitempty"`
	Private bool      `json:"is_private"`
}

type PostUpdate struct {
	Text     string   `json:"text,omitempty" gorm:"size:300"`
	Hashtags []string `json:"hashtags,omitempty"`
	MediaURI []string `json:"media_uri,omitempty"`
	Private bool      `json:"is_private,omitempty"`
}

type PostQuery struct {
	Nick string `json:"nick,omitempty"`
	Text string   `json:"text,omitempty" gorm:"size:300"`
	Hashtags []string `json:"hashtags,omitempty"`
    Limit uint `json:"limit,default=100" binding:"required,min=0,max=100"` 
    Page uint `json:"page,default=0" binding:"required,min=0"` 
}
