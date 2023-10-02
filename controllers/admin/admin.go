package admin

import (
	_ "github.com/SnapMsg-Inc/g1.gateway.api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create admin user godoc
// @Summary Create an admin user giving existing user
// @Param uid path string true "user id to become admin"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/:uid [post]
// @Security Bearer
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Delete any user godoc
// @Summary Delete a given user
// @Param uid path string true "user id to delete"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/:uid [delete]
// @Security Bearer
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Delete any post godoc
// @Summary Delete a given post
// @Param pid path string true "post id to delete"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/post/:pid [delete]
// @Security Bearer
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}
