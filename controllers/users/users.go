package users

import (
	"os"
	"fmt"
//	"io/ioutil"
//	"bytes"
//	"encoding/json"
	"net/http"
	
	"github.com/gin-gonic/gin"
//	models "github.com/SnapMsg-Inc/g1.gateway.api/models"
//	middlewares "github.com/SnapMsg-Inc/g1.gateway.api/middlewares"
)

var USERS_URL = os.Getenv("USERS_URL");


// List users godoc
// @Summary List users filtering by query 
// @Schemes
// @Description
// @Param uid query string false "user id"
// @Param email query string false "user email"
// @Param nick query string false "user nickname"
// @Param maxresults query int true "max results"
// @Param page query int true "page number"
// @Tags users methods 
// @x-order "1"
// @Accept */* 
// @Produce json
// @Success 200 array models.UserPublic
// @Router /users [get]
// @Security Bearer
func Get(c *gin.Context) {
	path_query := c.Request.URL.RequestURI();
	url := fmt.Sprintf("%s%s", USERS_URL, path_query);
	res, err := http.Get(url);

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error" : err.Error()});
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}


// Get current user godoc
// @Summary Get private data of current user 
// @Schemes
// @Description
// @Tags users methods 
// @Accept */* 
// @Produce json
// @Success 200 {} models.User
// @Router /users/me [get]
// @Security Bearer
func GetMe(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string);
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid);
	res, err := http.Get(url);

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error" : err.Error()});
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
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
// @Router /users/recommended [get]
// @Security Bearer
func GetRecommended(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
}


// Create user godoc
// @Summary Create a user  
// @Param userinfo body models.UserInfo true "User info"
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
	//var u models.UserInfo;
	//c.BindJSON(&u);
	uid := c.MustGet("FIREBASE_UID").(string);
	
	// forward request
	//u_json, err := json.Marshal(&u);
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid);
	res, err := http.Post(url, "application/json", c.Request.Body);
	
	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error" : err.Error()});
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}


// Update user godoc
// @Summary Update some user data 
// @Param nick query string false "new nickname"
// @Param interests query []string false "new interests"
// @Description 
// @Tags users methods 
// @Accept json
// @Produce json
// @Success 200 
// @Router /users [patch]
// @Security Bearer
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
}


// Delete user godoc
// @Summary Delete the user making the request
// @Schemes
// @Description 
// @Tags users methods 
// @Accept json
// @Produce json
// @Success 200 
// @Router /users [delete]
// @Security Bearer
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
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
// @Router /users/follow/{uid} [post]
// @Security Bearer
func Follow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
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
// @Router /users/follow/{uid} [delete]
// @Security Bearer
func Unfollow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
}



