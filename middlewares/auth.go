package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"

	option "google.golang.org/api/option"

	models "github.com/SnapMsg-Inc/g1.gateway.api/models"
)

var app *firebase.App
var Auth *auth.Client

func init() {
	var err error
	opt := option.WithCredentialsFile("fb-key.json")
	ctx := context.Background()
	app, err = firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	Auth, err = app.Auth(ctx)

	if err != nil {
		log.Fatalf("error creating auth client: %v", err)
	}
}

func authenticate(c *gin.Context) {
	// parse bearer token (`Bearer <id token>`)
	bearer_token := strings.SplitN(c.GetHeader("Authorization"), " ", 2)

	if bearer_token[0] != "Bearer" || len(bearer_token) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid bearer token format"})
		c.Abort()
		return
	}
	token_str := bearer_token[1]

	//ctx := context.Background()
	token, err := Auth.VerifyIDToken(c, token_str)
	if err != nil {
		// Unauthenticated
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthenticated"})
		c.Abort()
		return
	}
	c.Set("FIREBASE_ID_TOKEN", token)
	c.Set("FIREBASE_UID", token.UID)
	c.Set("IS_ADMIN", false)
	c.Next()
}

func Authentication() gin.HandlerFunc {
	return authenticate
}

func authorize(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s", os.Getenv("USERS_URL"), uid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "users service unavailable"})
		c.Abort()
		return
	}
	defer res.Body.Close()
	var users []models.User
	json.NewDecoder(res.Body).Decode(users)

	if len(users) == 0 || users[0].IsAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"status": "forbidden operation"})
		c.Abort()
		return
	}
	c.Set("IS_ADMIN", true)
	c.Next()
}

func Authorization() gin.HandlerFunc {
	return authorize
}
