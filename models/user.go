package models


type UserInfo struct {
	Email string `json:"email"`
	FullName string `json:"fullname"`
	Birthdate string `json:"birthdate"`
	Nick string `json:"nick"`
	Picture string `json:"pic"`
	Interests []string `json:"interests"`
	Zone string `json:"zone"`
}

type User struct {
	ID string `json:"uid" gorm:"primary_key"`
	Email string `json:"email,omitempty"`
	FullName string `json:"fullname,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Nick string `json:"nick"`
	Picture string `json:"pic"`
	Interests []string `json:"interests"`
	Zone string `json:"zone,omitempty"`
	IsAdmin bool `json:"is_admin,omitempty"`
}

type UserPublic struct {
	ID string `json:"uid"`
	Nick string `json:"nick"`
	Picture string `json:"pic"`
	Interests []string `json:"interests"`
}

