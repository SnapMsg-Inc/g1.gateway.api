package models

import (
    "os"
    "fmt"
    "time"
    "net/http"
    "encoding/json"
)

var MESSAGES_URL = os.Getenv("MESSAGES_URL")

type MentionNotification struct {
    MentionedUserIDs []string `json:"mentioned_user_ids"`
    MentioningUserID string   `json:"mentioning_user_id"`
    MessageContent   string   `json:"message_content"`
}

type Post struct {
	ID        string    `json:"pid" gorm:"primary_key"`
	UID       string    `json:"uid"`
	Timestamp time.Time `json:"timestamp"`
	Hashtags  []string  `json:"hashtags"`
	Text      string    `json:"text" gorm:"size:300"`
	MediaURI  []string  `json:"media_uri"`
	Likes     uint      `json:"likes"`
    IsBlocked bool      `json:"is_blocked"`
}

type PostCreate struct {
	UID      string   `json:"uid" swaggerignore:"true"`
	Text     string   `json:"text" gorm:"size:300"`
	Hashtags []string `json:"hashtags,omitempty"`
	MediaURI []string `json:"media_uri,omitempty"`
	Private bool      `json:"is_private"`
    MentionedUserIDs  []string `json:"mentioned_user_ids,omitempty"` // Agregar esta línea
}

type PostUpdate struct {
	Text     string   `json:"text,omitempty" gorm:"size:300"`
	Hashtags []string `json:"hashtags,omitempty"`
	MediaURI []string `json:"media_uri,omitempty"`
	Private bool      `json:"is_private,omitempty"`
}

type PostQuery struct {
	Nick string `form:"nick,omitempty" json:"nick,omitempty"`
	Text string   `form:"text,omitempty" json:"text,omitempty" gorm:"size:300"`
	Hashtags []string `form:"hashtags,omitempty" json:"hashtags,omitempty"`
    Limit uint `form:"limit,default=100" json:"limit,default=100" binding:"required,min=0,max=100"` 
    Page uint `form:"page,default=100" json:"page,default=0" binding:"required,min=0"` 
}

func nick2uid(nick string) []string {
    base_url := os.Getenv("USERS_URL");
    url := fmt.Sprintf("%s/users?nick=%s", base_url, nick);
    fmt.Printf("[INFO] %s\n", url);
    res, err := http.Get(url);
    var uids []string; 

    if (err == nil) {
        var users []UserPublic;
        defer res.Body.Close();
        json.NewDecoder(res.Body).Decode(&users);
    
        for _, user := range users {
            uids = append(uids, user.ID);
        } 
    }
    return uids;
}

func (q PostQuery) String() string {
    qstr := ""

    /*  optional qparams  */
    if q.Text != "" { qstr += fmt.Sprintf("text=%s&", q.Text) }
    for _, h := range q.Hashtags { 
        qstr += fmt.Sprintf("hashtags=%%23%s&", h[1:])
    }

    if q.Nick != "" { 
        for _, uid := range nick2uid(q.Nick) {
            qstr += fmt.Sprintf("uid=%s&", uid);
        }
    }

    /*  required qparams  */
    qstr += fmt.Sprintf("limit=%d&page=%d", q.Limit, q.Page)
    fmt.Printf("[INFO] qstr: %s\n", qstr);
    return qstr;
}

// func NotifyMention(mentionedUserIDs []string, mentioningUserID, messageContent string) {
//     notification := MentionNotification{
//         MentionedUserIDs: mentionedUserIDs,
//         MentioningUserID: mentioningUserID,
//         MessageContent:   messageContent,
//     }

//     var body bytes.Buffer
//     json.NewEncoder(&body).Encode(notification)

//     url := fmt.Sprintf("%s/notify-mention", MESSAGES_URL)
//     _, err := http.Post(url, "application/json", &body)
//     if err != nil {
//         fmt.Sprintf("Error al enviar notificación de mención: %v", err)
//     }
// }
