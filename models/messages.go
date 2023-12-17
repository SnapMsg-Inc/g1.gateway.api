package models

import (
    "fmt"
    "net/http"
    "bytes"
    "encoding/json"
    "os"
)
type TokenData struct {
    UID string `json:"user_id"`
    Token  string `json:"token"`
}

type MessageNotification struct {
    SenderAlias    string `json:"sender_alias"`
    ReceiverID     string `json:"receiver_id"`
    MessageContent string `json:"message_content"`
}

func NotifyMention(mentionedUserIDs []string, mentioningUserID, messageContent string) {
    notification := MentionNotification{
        MentionedUserIDs: mentionedUserIDs,
        MentioningUserID: mentioningUserID,
        MessageContent:   messageContent,
    }

    var body bytes.Buffer
    json.NewEncoder(&body).Encode(notification)

    url := fmt.Sprintf("%s/notify-mention", MESSAGES_URL)
    _, err := http.Post(url, "application/json", &body)
    if err != nil {
        fmt.Sprintf("Error al enviar notificación de mención: %v", err)
    }
}

func Uid2nick(uid string) (string, error) {
    baseUrl := os.Getenv("USERS_URL")
    url := fmt.Sprintf("%s/users/%s", baseUrl, uid)
    fmt.Printf("[INFO] %s\n", url)

    res, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    var user UserPublic
    if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
        return "", err
    }

    return user.Alias, nil
}