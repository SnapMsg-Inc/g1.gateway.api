package admin

import (
    "os"
    "fmt"
    "strings"
	"net/http"
    "encoding/json"

	gin "github.com/gin-gonic/gin"
	models "github.com/SnapMsg-Inc/g1.gateway.api/models"
    
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
// @Router /admin/{uid} [post]
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


// Get admin godoc
// @Summary Check if user is admin 
// @Param uid path string true "user id"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200
// @Failure 404
// @Router /admin/{uid} [get]
// @Security Bearer
func Get(c *gin.Context) {
    uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid);
    res, err := http.Get(url);

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()});
		return;
	}
    var user models.User;
    err = json.NewDecoder(res.Body).Decode(&user);
    
    if (err != nil) {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot parse body"});
        return;
    }
    
    if (!user.IsAdmin) {
        c.JSON(http.StatusNotFound, gin.H{"message": "admin does not exist"});
        return;
    }
    c.JSON(http.StatusOK, gin.H{"message": "admin exist"});
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
// @Router /admin/{uid} [delete]
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


// Get user godoc
// @Summary Get specific user 
// @Param uid path string true "user id"
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /admin/users/{uid} [get]
// @Security Bearer
func GetUser(c *gin.Context) {
    uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
    res, err := http.Get(url);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}   
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
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
	uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
    res, err := patchField(url, "is_blocked", true);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "user blocked"})
}


// Unblock any user godoc
// @Summary Unblock a given user 
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
	uid := c.Param("uid");
	url := fmt.Sprintf("%s/users/%s", USERS_URL, uid)
    res, err := patchField(url, "is_blocked", false);

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
    c.JSON(res.StatusCode, gin.H{"message": "user unblocked"})
}


// Get any posts godoc
// @Summary Get all posts filtered by query 
// @Param hashtags query []string false "hashtags" collectionFormat(multi)
// @Param nick query string false "author's nickname"
// @Param text query string false "text to match"
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags admin methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /admin/posts [get]
// @Security Bearer
func GetPosts(c *gin.Context) {
    var query models.PostQuery;
    
    if bind_err := c.ShouldBindQuery(&query); bind_err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error" : bind_err.Error });
    }
    url := fmt.Sprintf("%s/posts?private=true&public=true&blocked=true&%s", POSTS_URL, query.String());
    fmt.Printf("[INFO] %s\n", url);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
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

