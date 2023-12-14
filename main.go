package main

import (
    "os"
    "fmt"

    docs "github.com/SnapMsg-Inc/g1.gateway.api/docs"
    gin "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"     // swagger embed files
    ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

    admin "github.com/SnapMsg-Inc/g1.gateway.api/controllers/admin"
    posts "github.com/SnapMsg-Inc/g1.gateway.api/controllers/posts"
    users "github.com/SnapMsg-Inc/g1.gateway.api/controllers/users"
    messages "github.com/SnapMsg-Inc/g1.gateway.api/controllers/messages"
    stats "github.com/SnapMsg-Inc/g1.gateway.api/controllers/stats"
    middlewares "github.com/SnapMsg-Inc/g1.gateway.api/middlewares"
)

var SRV_ADDR = os.Getenv("SRV_ADDR")


// @title SnapMsg API
// @version 1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and Firebase JWT.
// @tag.name users methods
// @tag.name posts methods
// @tag.name admin methods
func main() {
    fmt.Println("Starting SnapMsg API...")

    docs.SwaggerInfo.BasePath = "/"
    router := gin.Default() // router with Default middleware
    
    /* private routes */
    private := router.Group("/")
    private.Use(middlewares.Authentication())
    private.Use(middlewares.CORS())
    {
        /* users routes */
        private.POST("/users", users.Create)
        private.GET("/users", users.Get)
        private.GET("/users/me", users.GetMe)
        private.GET("/users/me/recommended", users.GetRecommended)
        private.PATCH("/users/me", users.Update)
        private.DELETE("/users/me", users.Delete)

        private.POST("/users/me/follows/:uid", users.Follow)
        private.DELETE("/users/me/follows/:uid", users.Unfollow)

        private.GET("/users/:uid/follows/:otheruid", users.GetFollow)
        private.GET("/users/:uid/follows", users.GetFollows)
        private.GET("/users/:uid/followers", users.GetFollowers)

        /* posts routes */
        private.POST("/posts", posts.Create)
        private.GET("/posts", posts.Get)
        private.GET("/posts/me", posts.GetMe)
        
        write_posts := private.Group("/posts")
        write_posts.Use(middlewares.PostAuthorization())
        {
            write_posts.PATCH("/:pid", posts.Update)
            write_posts.DELETE("/:pid", posts.Delete)
        }

        private.GET("/posts/feed", posts.GetFeed)
        private.GET("/posts/me/recommended", posts.GetRecommended)
        private.GET("/trendings", posts.GetTrendingTopics)

        private.GET("/posts/favs", posts.GetFavs)
        private.GET("/posts/favs/:pid", posts.Favs)
        private.POST("/posts/favs/:pid", posts.Fav)
        private.DELETE("/posts/favs/:pid", posts.Unfav)

        private.GET("/posts/likes/:pid", posts.GetLike)
        private.POST("/posts/likes/:pid", posts.Like)
        private.DELETE("/posts/likes/:pid", posts.Unlike)        

        private.GET("/posts/me/snapshares", posts.GetSnapshares)
        private.GET("/posts/snapshares/:pid", posts.IsSnapshared)
        private.DELETE("/posts/snapshares/:pid", posts.DeleteSnapshare)
        private.POST("/posts/snapshares/:pid", posts.CreateSnapshare)
        private.GET("/posts/stats/me", posts.GetMyStats)

        /* messaging routes */
        private.POST("/messages/token", messages.RegisterToken)
        private.POST("/messages", messages.NotifyMessage)

        
        /* stats routes */
        private.POST("/stats", stats.PushStat)
        

        /* admin routes (must authorize) */
        admin_group := private.Group("/admin")
        admin_group.Use(middlewares.AdminAuthorization())
        {
            admin_group.POST("/:uid", admin.Create)
            admin_group.DELETE("/:uid", admin.Delete)

            admin_group.GET("/users/:uid", admin.GetUser)
            admin_group.POST("/users/:uid/block", admin.BlockUser)
            admin_group.DELETE("/users/:uid/block", admin.UnblockUser)

            admin_group.POST("/posts/:pid/block", admin.BlockPost)
            admin_group.DELETE("/posts/:pid/block", admin.UnblockPost)
            admin_group.GET("/posts", admin.GetPosts)
        }
    }

    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run(SRV_ADDR)
}
