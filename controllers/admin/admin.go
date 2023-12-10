package admin

import (
    "os"
    "fmt"
    "strings"
	"net/http"

	gin "github.com/gin-gonic/gin"
	_ "github.com/SnapMsg-Inc/g1.gateway.api/models"
)

var USERS_URL = os.Getenv("USERS_URL");
var POSTS_URL = os.Getenv("POSTS_URL");


/* boilerplate code for patch single field */
func patchField(url string, k string, v any) (*http.Response, error) {
    json_str := fmt.Sprintf(`{"%s": %#v}`, k, v);
    body := strings.NewReader(json_str);
    fmt.Printf("[INFO] %s\n", body);
	req, _ := http.NewRequest("PATCH", url, body)
	client := &http.Client{}
	return client.Do(req)
}


// Create admin user godoc
// @Summary Create an admin user giving existing user
// @Param uid path string true "user id to become admin"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/{uid} [post]
// @Security Bearer
func Create(c *gin.Context) {
	uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
    res, err := patchField(url, "is_admin", true);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "admin added"})
}


// Delete admin user godoc
// @Summary Remove admin status from existing user
// @Param uid path string true "user id of the admin"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/{uid} [delete]
// @Security Bearer
func Delete(c *gin.Context) {
	uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
    res, err := patchField(url, "is_admin", false);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "admin deleted"})
}


// Block any user godoc
// @Summary Block a given user
// @Param uid path string true "user id to block"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/{uid}/block [post]
// @Security Bearer
func BlockUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}


// Unblock any user godoc
// @Summary Unblock a blocked user
// @Param uid path string true "user id to unblock"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/users/{uid}/block [delete]
// @Security Bearer
func UnblockUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}


// Get any posts godoc
// @Summary Get all posts filtered by query 
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/posts [get]
// @Security Bearer
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}


// Block any post godoc
// @Summary Block a given post
// @Param pid path string true "post id to block"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/posts/{pid}/block [post]
// @Security Bearer
func BlockPost(c *gin.Context) {
	pid := c.Param("pid");
	url := fmt.Sprintf("%s/posts/%s", POSTS_URL, pid)
    res, err := patchField(url, "is_blocked", true);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "post blocked"})
}


// Unblock any post godoc
// @Summary Unblock a given post
// @Param pid path string true "post id to unblock"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Router /admin/posts/{pid}/block [delete]
// @Security Bearer
func UnblockPost(c *gin.Context) {
	pid := c.Param("pid");
	url := fmt.Sprintf("%s/posts/%s", POSTS_URL, pid)
    res, err := patchField(url, "is_blocked", false);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "post unblocked"})
}

