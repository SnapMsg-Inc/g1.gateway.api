package posts

import (
    "fmt"
    "os"
    "bytes"
    "strings"
	"net/http"
    "encoding/json"
    
    models "github.com/SnapMsg-Inc/g1.gateway.api/models"
	"github.com/gin-gonic/gin"
)

var USERS_URL = os.Getenv("USERS_URL")
var POSTS_URL = os.Getenv("POSTS_URL")


// Get posts godoc
// @Summary Get posts filtering by query
// @Param hashtags query []string false "hashtags"
// @Param nick query string false "author's nickname"
// @Param text query string false "text to match"
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts [get]
// @Security Bearer
func Get(c *gin.Context) {
    // bindear la query a un struct (modelo PostQuery)
    // convertir nick a uid con el ms de users
    // enviar la request al ms de posts
    var query models.PostQuery;
    bind_err := c.ShouldBindQuery(&query);
    
    if bind_err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error" : bind_err.Error });
    }
    path_query := strings.Split(c.Request.URL.RequestURI(), "?")[1];
    nick := c.Query("nick");
    // fetch user's uid 
    fmt.Printf("[INFO] %s\n", nick);

    if nick != "" {
        url := fmt.Sprintf("%s/users?nick=%s", USERS_URL, nick);
        fmt.Printf("[INFO] %s\n", url);
        res, err := http.Get(url);
    
        if (err != nil) {
            c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
            return;
        }
        var user []models.UserPublic;
        defer res.Body.Close();
        err = json.NewDecoder(res.Body).Decode(&user);
        
        if (err != nil) {
            c.JSON(http.StatusInternalServerError, gin.H{ "error" : "cannot parse body" });
            return;
        }
        path_query += "&uid=" + user[0].ID;
    }
    url := fmt.Sprintf("%s/posts?%s", POSTS_URL, path_query);
    fmt.Printf("[INFO] %s\n", url);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}

// Get posts of current user (public and privates) godoc
// @Summary Get posts filtering by query
// @Param hashtags query []string false "hashtags"
// @Param text query string false "text to match"
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/me [get]
// @Security Bearer
func GetMe(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
    query := strings.Split(c.Request.URL.RequestURI(), "?")[1];
    url := fmt.Sprintf("%s/posts?uid=%s&%s&private=true&public=true", POSTS_URL, uid, query);
    fmt.Printf("[url] %s\n", url);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}

// Create post godoc
// @Summary Create a new post
// @Param PostCreate body models.PostCreate true "data for the new post"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts [post]
// @Security Bearer
func Create(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string)
   
    // set PostCreate data
    var post models.PostCreate;
    c.ShouldBindJSON(&post);
    post.UID = uid;
    // post.Nick = user[0].Nick;
    
    var body bytes.Buffer;
    json.NewEncoder(&body).Encode(post);
    
    url := fmt.Sprintf("%s/posts", POSTS_URL);
    res, err := http.Post(url, "application/json", &body);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}

// Update post godoc
// @Summary Update some fields of a post
// @Param pid path string true "post id to update"
// @Param PostUpdate body models.PostUpdate true "data for update the post" 
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Failure 403  
// @Router /posts/{pid} [patch]
// @Security Bearer
func Update(c *gin.Context) {
    pid := c.Param("pid");
	url := fmt.Sprintf("%s/posts/%s", POSTS_URL, pid);
    fmt.Printf("[url] %s [PID] %s\n", url, pid);
	req, _ := http.NewRequest("PATCH", url, c.Request.Body);
	client := &http.Client{};
	res, err := client.Do(req);

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
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
	url := fmt.Sprintf("%s/posts/%s", POSTS_URL, c.Param("pid"));
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get feed godoc
// @Summary Get feed of the user making the request
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/feed [get]
// @Security Bearer
func GetFeed(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string)
    query := strings.Split(c.Request.URL.RequestURI(), "?")[1];

    // fetch follows list
    url := fmt.Sprintf("%s/users/%s/follows", USERS_URL, uid);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
    }
    var follows []models.UserPublic;
    err = json.NewDecoder(res.Body).Decode(&follows);
    
    if (err != nil) {
        c.JSON(http.StatusInternalServerError, gin.H{ "error" : "cannot parse body" });
        return;
    }

    if (len(follows) == 0) {
        c.JSON(http.StatusOK, gin.H{"data" : []string{}});
        return;
    }

    // parse follows uid to http query format
    for _, follow := range follows {
        query += "&uid=" + follow.ID;
    }
    url = fmt.Sprintf("%s/posts?%s&private=True&public=True", POSTS_URL, query);

    // fetch (private and public) posts of followed
    fmt.Printf("[URL] %s\n", url)
    res, err = http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
}

// Get recommended godoc
// @Summary Get recommended posts for a user
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/recommended [get]
// @Security Bearer
func GetRecommended(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string)
    path_query := strings.Split(c.Request.URL.RequestURI(), "?")[1];
    url := fmt.Sprintf("%s/posts/%s/recommended?%s", POSTS_URL, uid, path_query);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
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
// @Router /posts/likes/{pid} [post]
// @Security Bearer
func Like(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
    url := fmt.Sprintf("%s/posts/%s/likes/%s", POSTS_URL, uid, pid);
    res, err := http.Post(url, "application/json", nil);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
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
// @Router /posts/likes/{pid} [delete]
// @Security Bearer
func Unlike(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
	url := fmt.Sprintf("%s/posts/%s/likes/%s", POSTS_URL, uid, pid)
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// List favs godoc
// @Summary List user fav posts
// @Param limit query int true "limit" default(100) maximum(100) minimum(0)
// @Param page query int true "page" default(0) minimum(0)
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200 array models.Post
// @Router /posts/favs [get]
// @Security Bearer
func GetFavs(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
    path_query := strings.Split(c.Request.URL.RequestURI(), "?")[1];
    url := fmt.Sprintf("%s/posts/%s/favs?%s", POSTS_URL, uid, path_query);
    res, err := http.Get(url);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil);
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
// @Router /posts/favs/{pid} [post]
// @Security Bearer
func Fav(c *gin.Context) {
	uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
    url := fmt.Sprintf("%s/posts/%s/favs/%s", POSTS_URL, uid, pid);
    res, err := http.Post(url, "application/json", nil);

    if (err != nil) {
        c.JSON(res.StatusCode, gin.H{ "error" : err.Error });
        return;
    }

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
// @Router /posts/favs/{pid} [delete]
// @Security Bearer
func Unfav(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
    url := fmt.Sprintf("%s/posts/%s/favs/%s", POSTS_URL, uid, pid)
    req, _ := http.NewRequest("DELETE", url, nil)
    client := &http.Client{}
    res, err := client.Do(req)

    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get Like godoc
// @Summary Check if current user liked a given post
// @Schemes
// @Description
// @Param pid path string true "pid to check like"
// @Tags posts methods
// @Accept json
// @Produce json
// @Failure 404
// @Success 200
// @Router /posts/likes/{pid} [get]
// @Security Bearer
func GetLike(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
    url := fmt.Sprintf("%s/posts/%s/likes/%s", POSTS_URL, uid, pid);
    fmt.Println("entre");
    fmt.Println(url);

    res, err := http.Get(url);
    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get Favs  godoc
// @Summary Check if current user favs a given post
// @Schemes
// @Description
// @Param pid path string true "pid to check fav"
// @Tags posts methods
// @Accept json
// @Produce json
// @Failure 404
// @Success 200
// @Router /posts/favs/{pid} [get]
// @Security Bearer
func Favs(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
    pid := c.Param("pid");
    url := fmt.Sprintf("%s/posts/%s/favs/%s", POSTS_URL, uid, pid);
    res, err := http.Get(url);

    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// Get Trending Topics godoc
// @Summary Get Trending Topics
// @Description Retrieves a list of trending topics.
// @Tags posts methods
// @Accept json
// @Produce json
// @Param limit query int true "Limit of topics to retrieve" default(10) maximum(100) minimum(0)
// @Param page query int false "Page number for pagination" default(0) minimum(0)
// @Success 200 {array} models.Post
// @Router /trendings [get]
// @Security Bearer
func GetTrendingTopics(c *gin.Context) {
    limit := c.DefaultQuery("limit", "100")
    page := c.DefaultQuery("page", "0")

    url := fmt.Sprintf("%s/trendings?limit=%s&page=%s", POSTS_URL, limit, page)
    res, err := http.Get(url)
        fmt.Println("URL de la petición:", url)

    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}


// Delete snapshare godoc
// @Summary Delete snapshare owned by current user
// @Param pid path string true "snapshare id to delete"
// @Schemes
// @Description
// @Tags posts methods
// @Accept json
// @Produce json
// @Success 200
// @Router /posts/{pid}/snapshares [delete]
// @Security Bearer
func DeleteSnapshare(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string);
	url := fmt.Sprintf("%s/posts/%s/snapshares/%s", POSTS_URL, uid, c.Param("pid"));
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if (err != nil) {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// GetSnapshares godoc
// @Summary Get snapshares of current user
// @Description Retrieves the snapshares associated with the current authenticated user.
// @Tags posts methods
// @Accept json
// @Produce json
// @Param limit query int false "Limit of snapshares to retrieve" default(100) maximum(100) minimum(0)
// @Param page query int false "Page number for pagination" default(0) minimum(0)
// @Success 200
// @Router /posts/me/snapshares [get]
// @Security Bearer
func GetSnapshares(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string)

    limit := c.DefaultQuery("limit", "100")
    page := c.DefaultQuery("page", "0")

    url := fmt.Sprintf("%s/posts/%s/snapshares?limit=%s&page=%s", POSTS_URL, uid, limit, page)
    res, err := http.Get(url)

    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// CreateSnapshare godoc
// @Summary Create a snapshare
// @Description Creates a new snapshare for a given post by the current user.
// @Tags posts methods
// @Accept json
// @Produce json
// @Param pid path string true "Post ID"
// @Success 200 
// @Router /posts/snapshares/{pid} [post]
// @Security Bearer
func CreateSnapshare(c *gin.Context) {
    pid := c.Param("pid")
    uid := c.MustGet("FIREBASE_UID").(string)

    url := fmt.Sprintf("%s/posts/%s/snapshares/%s", POSTS_URL, uid, pid)

    res, err := http.Post(url, "application/json", nil)

    if err != nil {
        c.JSON(res.StatusCode, gin.H{"error": err.Error()})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}

// IsSnapshared godoc
// @Summary Check if a post is snapshared by current user
// @Description Checks if the current authenticated user has snapshared a specific post.
// @Tags posts methods
// @Accept json
// @Produce json
// @Param pid path string true "Post ID"
// @Success 200 
// @Failure 404 
// @Router /posts/snapshares/{pid} [get]
// @Security Bearer
func IsSnapshared(c *gin.Context) {
    uid := c.MustGet("FIREBASE_UID").(string)
    pid := c.Param("pid")

    url := fmt.Sprintf("%s/posts/%s/snapshares/%s", POSTS_URL, uid, pid)

    res, err := http.Get(url)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if res.StatusCode == http.StatusNotFound {
        c.JSON(http.StatusNotFound, gin.H{"error": "not snapshared"})
        return
    }
    c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}
