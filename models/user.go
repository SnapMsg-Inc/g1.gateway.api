package models


type UserInfo struct {
	ID string `json:"-" gorm:"primary_key"`
	Email string `json:"email"`
	FullName string `json:"fullname"`
	Nick string `json:"nickname"`
	Interests []string `json:"interests"`
	Zone string `json:"zone"`
}

type User struct {
	ID string `json:"uid" gorm:"primary_key"`
	Email string `json:"email,omitempty"`
	FullName string `json:"fullname,omitempty"`
	Nick string `json:"nickname"`
	Interests []string `json:"interests"`
	Followers uint `json:"followers,omitempty"`
	Followings uint `json:"followings,omitempty"`
	Zone string `json:"zone,omitempty"`
}
