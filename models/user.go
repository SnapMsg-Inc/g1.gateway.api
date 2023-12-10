package models

type User struct {
	ID        string             `json:"uid" gorm:"primary_key"`
	Email     string             `json:"email,omitempty"`
	FullName  string             `json:"fullname,omitempty"`
	Birthdate string             `json:"birthdate,omitempty" example:"YYYY-MM-DD"`
	Alias	  string             `json:"alias"`
	Nick      string             `json:"nick"`
	Follows   int                `json:"follows"`
	Followers int                `json:"followers"`
	Picture   string             `json:"pic"`
	Interests []string           `json:"interests"`
	Ocupation string             `json:"ocupation,omitempty"`
	Zone      map[string]float32 `json:"zone,omitempty" example:"latitude:0.00,longitude:0.00"`
	IsAdmin   bool               `json:"is_admin,omitempty"`
}

type UserCreate struct {
	Email     string             `json:"email" binding:"required"`
	FullName  string             `json:"fullname" binding:"required"`
	Birthdate string             `json:"birthdate" example:"YYYY-MM-DD" binding:"required"`
	Alias	  string             `json:"alias" binding:"required"`
	Nick      string             `json:"nick" binding:"required"`
	Picture   string             `json:"pic" binding:"required"`
	Interests []string           `json:"interests" binding:"required"`
	Ocupation string             `json:"ocupation,omitempty" binding:"required"`
	Zone      map[string]float32 `json:"zone,omitempty" example:"latitude:0.00,longitude:0.00" binding:"required"`
}

type UserPublic struct {
	ID        string   `json:"uid"`
	Alias	  string   `json:"alias"`
	Nick      string   `json:"nick"`
	Follows   int      `json:"follows"`
	Followers int      `json:"followers"`
	Picture   string   `json:"pic"`
	Interests []string `json:"interests"`
}

type UserUpdate struct {
	Alias	  string             `json:"alias,omitempty"`
	Nick      string             `json:"nick,omitempty"`
	Picture   string             `json:"pic,omitempty"`
	Interests []string           `json:"interests,omitempty"`
	Ocupation string             `json:"ocupation,omitempty"`
	Zone      map[string]float32 `json:"zone,omitempty" example:"latitude:0.00,longitude:0.00"`
}

