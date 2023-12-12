package messages

import (
    
    "bytes"
	"net/http"
	"fmt"
    "encoding/json"
    
    models "github.com/SnapMsg-Inc/g1.gateway.api/models"
	"github.com/gin-gonic/gin"
)

var MESSAGES_URL = "https://messages-ms-marioax.cloud.okteto.net"

// RegisterToken godoc
// @Summary Register a new notification token
// @Param token_data body models.TokenData true "Token Data"
// @Schemes
// @Description register a new notification token
// @Tags messages methods
// @Accept json
// @Produce json
// @Success 200 
// @Router /messages/token [post]
// @Security Bearer
func RegisterToken(c *gin.Context) {
    var tokenData models.TokenData

    if err := c.ShouldBindJSON(&tokenData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	url := fmt.Sprintf("%s/register-token", MESSAGES_URL);

    var body bytes.Buffer
    json.NewEncoder(&body).Encode(tokenData)

    res, err := http.Post(url, "application/json", &body)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Reenviar la respuesta del servicio Python al cliente
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// NotifyMessage godoc
// @Summary Send a message notification
// @Description send a message notification
// @Tags messages methods
// @Accept json
// @Produce json
// @Param message_notification body models.MessageNotification true "Message Notification Data"
// @Success 200
// @Router /messages [post]
// @Security Bearer
func NotifyMessage(c *gin.Context) {
    var notification models.MessageNotification
    if err := c.ShouldBindJSON(&notification); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Serializar la notificación para enviarla al servicio de Python
    var body bytes.Buffer
    json.NewEncoder(&body).Encode(notification)

    // URL del endpoint de Python que maneja las notificaciones de mensajes
    url := fmt.Sprintf("%s/notify-message", MESSAGES_URL);
    res, err := http.Post(url, "application/json", &body)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Reenviar la respuesta del servicio Python al cliente
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}