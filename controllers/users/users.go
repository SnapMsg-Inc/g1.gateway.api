package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/SnapMsg-Inc/g1.gateway.api/models"
)


// List users godoc
// @Summary List users filtering by query 
// @Param uid query string false "user id"
// @Param email query string false "user email"
// @Param nick query string false "user nickname"
// @Param maxresults query int true "max results"
// @Param page query int true "page number"
// @Schemes
// @Description 
// @Tags users methods 
// @Accept */* 
// @Produce json
// @Success 200 array models.User
// @Router /users [get]
// @Security Bearer
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
}


// List recommended users godoc
// @Summary Recommend users for the user making the request
// @Schemes
// @Description 
// @Tags users methods 
// @Accept json
// @Produce json
// @Success 200 array models.User
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
// @Tags users methods 
// @Accept json 
// @Produce json
// @Success 200 
// @Router /users [post]
// @Security Bearer
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "not implemented yet" })
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
// @Router /users [put]
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



