package models

type TokenData struct {
    UID string `json:"user_id"`
    Token  string `json:"token"`
}

type MessageNotification struct {
    SenderAlias    string `json:"sender_alias"`
    ReceiverID     string `json:"receiver_id"`
    MessageContent string `json:"message_content"`
}