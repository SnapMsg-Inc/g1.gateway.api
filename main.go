package main

import (
	//"net/http"
	docs "github.com/SnapMsg-Inc/g1.gateway.api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	posts "github.com/SnapMsg-Inc/g1.gateway.api/controllers/posts"
	users "github.com/SnapMsg-Inc/g1.gateway.api/controllers/users"
	admin "github.com/SnapMsg-Inc/g1.gateway.api/controllers/admin"
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
	router.POST("/users", users.Create)

	/* private routes */
	private := router.Group("/")
	private.Use(middlewares.Authentication())
	{
		/* users routes */
		private.GET("/users", users.Get)
		private.GET("/users/recommended", users.GetRecommended)
		private.PUT("/users", users.Update)
		private.DELETE("/users", users.Delete)

		private.POST("/users/follow/:uid", users.Follow)
		private.DELETE("/users/follow/:uid", users.Unfollow)

		/* posts routes */
		private.GET("/posts", posts.Get)
		private.GET("/posts/feed", posts.GetFeed)
		private.GET("/posts/recommended", posts.GetRecommended)
		private.POST("/posts", posts.Create)
		private.PUT("/posts/:pid", posts.Update)
		private.DELETE("/posts/:pid", posts.Delete)

		private.POST("/posts/like/:pid", posts.Like)
		private.DELETE("/posts/like/:pid", posts.Unlike)

		private.GET("/posts/fav", posts.GetFavs)
		private.POST("/posts/fav/:pid", posts.Fav)
		private.DELETE("/posts/fav/:pid", posts.Unfav)

		/* messaging routes */

		/* admin routes (must authorize) */
		admin_group := router.Group("/admin")
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
