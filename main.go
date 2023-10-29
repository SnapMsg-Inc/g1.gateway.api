package main

import (
    //"net/http"
    docs "github.com/SnapMsg-Inc/g1.gateway.api/docs"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"     // swagger embed files
    ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

    admin "github.com/SnapMsg-Inc/g1.gateway.api/controllers/admin"
    posts "github.com/SnapMsg-Inc/g1.gateway.api/controllers/posts"
    users "github.com/SnapMsg-Inc/g1.gateway.api/controllers/users"
    middlewares "github.com/SnapMsg-Inc/g1.gateway.api/middlewares"
)

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
    docs.SwaggerInfo.BasePath = "/"
    router := gin.Default() // router with Default middleware

    /* create user is public  */

    /* private routes */
    private := router.Group("/")
    private.Use(middlewares.Authentication())
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
        private.PATCH("/posts/:pid", posts.Update)
        private.DELETE("/posts/:pid", posts.Delete)

        private.GET("/posts/feed", posts.GetFeed)
        private.GET("/posts/recommended", posts.GetRecommended)

        private.GET("/posts/fav", posts.GetFavs)
        private.POST("/posts/fav/:pid", posts.Fav)
        private.DELETE("/posts/fav/:pid", posts.Unfav)

        private.POST("/posts/like/:pid", posts.Like)
        private.DELETE("/posts/like/:pid", posts.Unlike)

        /* messaging routes */


        /* admin routes (must authorize) */
        admin_group := private.Group("/admin")
        admin_group.Use(middlewares.Authorization())
        {
            admin_group.PUT("/users/:uid", admin.Create)
            admin_group.DELETE("/users/:uid", admin.DeleteUser)
            admin_group.DELETE("/posts/:pid", admin.DeletePost)
        }
    }

    router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Run(":3000") // service running in port 3000
}
