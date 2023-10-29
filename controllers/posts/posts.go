package posts

import (
    "fmt"
    "os"
	"net/http"
    "encoding/json"

	_ "github.com/SnapMsg-Inc/g1.gateway.api/models"
	"github.com/gin-gonic/gin"
)

// Get posts godoc
// @Summary Get posts filtering by query
// @Param hashtags query []string false "hashtags"
// @Param nick query string false "author's nickname"
// @Param text query string false "text to match"
// @Param query int true "max results"
// @Param page query int true "page"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts [get]
// @Security Bearer
func Get(c *gin.Context)
{
    path_query := c.Request.URL.RequestURI();
    url := fmt.Sprintf("%s%s", POSTS_URL, path_query);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err,Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}

// Create post godoc
// @Summary Create a new post
// @Param postinfo body models.PostInfo true "data for the new post"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts [post]
// @Security Bearer
func Create(c *gin.Context)
{
    uid := c.MustGet("FIREBASE_UID").(string);

    // fetch user's nickname
    url := fmt.Sprintf("%s/users?uid=%s", USERS_URL, uid);
    res, err := http.Get(url);
    
    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err,Error });
        return;
    }
    
}

// Update post godoc
// @Summary Update some fields of a post
// @Param pid path string true "post id to update"
// @Param text body string false "new text for the post"
// @Param hashtags body []string false "new hashtags for the post"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/{pid} [patch]
// @Security Bearer
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Delete post godoc
// @Summary Delete post owned by current user
// @Param pid path string true "post id to delete"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/{pid} [delete]
// @Security Bearer
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Get feed godoc
// @Summary Get feed of the user making the request
// @Param limit query int true "max results"
// @Param page query int true "page"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/feed [get]
// @Security Bearer
func GetFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Get recommended godoc
// @Summary Get recommended posts for a user
// @Param limit query int true "max results"
// @Param page query int true "page"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/recommended [get]
// @Security Bearer
func GetRecommended(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Like post godoc
// @Summary Add a like to a post
// @Param pid path string true "post id to like"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/like/{pid} [post]
// @Security Bearer
func Like(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Unlike post godoc
// @Summary Unlike a post
// @Param pid path string true "post id to unlike"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/like/{pid} [delete]
// @Security Bearer
func Unlike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// List favs godoc
// @Summary List user fav posts
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/fav [get]
// @Security Bearer
func GetFavs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Add fav godoc
// @Summary Add a post to user favs
// @Param pid path string true "post id to mark as fav"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/fav/{pid} [post]
// @Security Bearer
func Fav(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}

// Unfav a post godoc
// @Summary Remove a post from user favs
// @Param pid path string true "post id to unfav"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/fav/{pid} [delete]
// @Security Bearer
func Unfav(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "not implemented yet"})
}
