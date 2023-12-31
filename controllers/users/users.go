package users

import (
	"fmt"
	"os"
	//	"io/ioutil"
	"bytes"
    "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/SnapMsg-Inc/g1.gateway.api/models"
	// middlewares "github.com/SnapMsg-Inc/g1.gateway.api/middlewares"
)

var USERS_URL = os.Getenv("USERS_URL")
var POSTS_URL = os.Getenv("POSTS_URL")
var MESSAGES_URL = os.Getenv("MESSAGES_URL")

// List users godoc
// @Summary List public users data filtering by query
// @Schemes
// @Description
// @Param uid query string false "user id"
// @Param email query string false "user email"
// @Param nick query string false "user nickname"
// @Param limit query int true "max results" default(100) maximum(100) minimum(0)
// @Param page query int true "page number" default(0) minimum(0)
// @Tags users methods
// @x-order "1"
// @Accept */*
// @Produce json
// @Success 200 array models.UserPublic
// @Router /users [get]
// @Security Bearer
func Get(c *gin.Context) {
	path_query := c.Request.URL.RequestURI()
	url := fmt.Sprintf("%s%s", USERS_URL, path_query)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get current user godoc
// @Summary Get private data of current user
// @Schemes
// @Description
// @Tags users methods
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Router /users/me [get]
// @Security Bearer
func GetMe(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// List recommended users godoc
// @Summary Recommend users for the user making the request
// @Schemes
// @Description
// @Tags users methods
// @x-order "2"
// @Accept json
// @Produce json
// @Success 200 array models.UserPublic
// @Router /users/me/recommended [get]
// @Security Bearer
func GetRecommended(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s/recommended", USERS_URL, uid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Create user godoc
// @Summary Create a user
// @Param userinfo body models.UserCreate true "User creation data"
// @Schemes
// @Description
// @x-order "0"
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users [post]
// @Security Bearer
func Create(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)

	// forward request
	//u_json, err := json.Marshal(&u);
    var user models.UserCreate;

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return;
    }
    fmt.Printf("[UserCreate] %s\b", user);
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid);
    var body bytes.Buffer;
    json.NewEncoder(&body).Encode(user);

	res, err := http.Post(url, "application/json", &body);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Update user godoc
// @Summary Update current user data
// @Param UserUpdate body models.UserUpdate true "data for update the user" 
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/me [patch]
// @Security Bearer
func Update(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)

    var data models.UserUpdate;
    bind_err := c.ShouldBindJSON(&data);
    
    if bind_err != nil {
        c.JSON(http.StatusBadRequest, gin.H{ "error" : bind_err.Error });
    }
    var body bytes.Buffer;
    json.NewEncoder(&body).Encode(data);

	req, _ := http.NewRequest("PATCH", url, &body)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Delete user godoc
// @Summary Delete the user making the request
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/me [delete]
// @Security Bearer
func Delete(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
    
	client := &http.Client{}
	url := fmt.Sprintf("%s/posts/%s", POSTS_URL, uid)
    req, _ := http.NewRequest("DELETE", url, nil)
	res, err := client.Do(req)

    if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    
	url = fmt.Sprintf("%s/users/%s", USERS_URL, uid)
	req, _ = http.NewRequest("DELETE", url, nil)
	res, err = client.Do(req)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Follow user godoc
// @Summary Follow a given user
// @Param uid path string true "user id to follow"
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/me/follows/{uid} [post]
// @Security Bearer
func Follow(c *gin.Context) {
	otheruid := c.Param("uid")
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s/follows/%s", USERS_URL, uid, otheruid)
	res, err := http.Post(url, "application/json", nil)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}

	if res.StatusCode == http.StatusOK {

        followerAlias, err := models.Uid2nick(uid)
    	if err != nil {

       	 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el alias del seguidor"})
        	return
    	}
		notifyURL := fmt.Sprintf("%s/notify-follow/%s/%s", MESSAGES_URL, followerAlias, otheruid)
		_, notifyErr := http.Post(notifyURL, "application/json", nil)
		if notifyErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo notificar al usuario"})
			return
		}
	}
	
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Unfollow user godoc
// @Summary Unfollow a given user
// @Param uid path string true "user id to unfollow"
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/me/follows/{uid} [delete]
// @Security Bearer
func Unfollow(c *gin.Context) {
	otheruid := c.Param("uid")
	uid := c.MustGet("FIREBASE_UID").(string)
	url := fmt.Sprintf("%s/users/%s/follows/%s", USERS_URL, uid, otheruid)
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get follow
// @Summary Check if uid follows otheruid
// @Param uid path string true "user id who follows"
// @Param otheruid path string true "the user id beeing followed"
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Failure 403
// @Failure 401
// @Router /users/{uid}/follows/{otheruid} [get]
// @Security Bearer
func GetFollow(c *gin.Context) {
	is_admin := c.MustGet("IS_ADMIN").(bool)
	me := c.MustGet("FIREBASE_UID").(string)
	uid := c.Param("uid")
	otheruid := c.Param("otheruid")

	if me != uid && me != otheruid && !is_admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "you cannot access this resource"})
		return
	}

	url := fmt.Sprintf("%s/users/%s/follows/%s", USERS_URL, uid, otheruid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// List follows godoc
// @Summary List follows of a given user
// @Param uid path string true "user id to list follows"
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{uid}/follows [get]
// @Security Bearer
func GetFollows(c *gin.Context) {
	me := c.MustGet("FIREBASE_UID").(string)
	is_admin := c.MustGet("IS_ADMIN").(bool)
	uid := c.Param("uid")

	// must exist mutual follow in order to list the follows of other user
	if me != uid && !followExist(me, uid) && !followExist(uid, me) && !is_admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "you cannot access this resource"})
		return
	}
	url := fmt.Sprintf("%s/users/%s/follows", USERS_URL, uid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// List followers godoc
// @Summary List followers of given user
// @Param uid path string true "user id to list followers"
// @Schemes
// @Description
// @Tags users methods
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{uid}/followers [get]
// @Security Bearer
func GetFollowers(c *gin.Context) {
	me := c.MustGet("FIREBASE_UID").(string)
	is_admin := c.MustGet("IS_ADMIN").(bool)
	uid := c.Param("uid")

	// must exist mutual follow in order to list the followers of other user
	if me != uid && !followExist(me, uid) && !followExist(uid, me) && !is_admin {
		c.JSON(http.StatusForbidden, gin.H{"error": "you cannot access this resource"})
		return
	}
	url := fmt.Sprintf("%s/users/%s/followers", USERS_URL, uid)
	res, err := http.Get(url)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

func followExist(uid string, followed string) bool {
	url := fmt.Sprintf("%s/users/%s/follows/%s", USERS_URL, uid, followed)
	res, err := http.Get(url)

	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

