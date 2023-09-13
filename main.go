package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	users "github.com/SnapMsg-Inc/g1.gateway.api/controllers/users"
	posts "github.com/SnapMsg-Inc/g1.gateway.api/controllers/posts"
	//	admin "github.com/SnapMsg-Inc/g1.gateway.api/controllers/admin"
	docs "github.com/SnapMsg-Inc/g1.gateway.api/docs"
)


// @title SnapMsg API
// @version 1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and Firebase JWT.
func main() {
	router := gin.Default() // router with Default middleware
	docs.SwaggerInfo.BasePath = "/"

	/* users routes */
	router.GET("/users", users.Get)
	router.GET("/users/recommended", users.GetRecommended)
	router.POST("/users", users.Create)
	router.PUT("/users", users.Update)
	router.DELETE("/users", users.Delete)

	router.POST("/users/follow/:uid", users.Follow)
	router.DELETE("/users/follow/:uid", users.Unfollow)

	/* posts routes */
	router.GET("/posts", posts.Get)
	router.GET("/posts/feed", posts.GetFeed)
	router.GET("/posts/recommended", posts.GetRecommended)
	router.POST("/posts", posts.Create)
	router.PUT("/posts/:pid", posts.Update)
	router.DELETE("/posts/:pid", posts.Delete)

	router.POST("/posts/like/:pid", posts.Like)
	router.DELETE("/posts/like/:pid", posts.Unlike)

	router.GET("/posts/fav", posts.GetFavs)
	router.POST("/posts/fav/:pid", posts.Fav)
	router.DELETE("/posts/fav/:pid", posts.Unfav)

	/* messaging routes */
	/* admin routes */
	//router.PUT("/admin/users/:uid", admin.Create)
	//router.DELETE("/admin/users/:uid", admin.DeleteAnyUser)
	//router.DELETE("/admin/posts/:pid", admin.DeleteAnyPost)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":3000") // service running in port 3000
}
