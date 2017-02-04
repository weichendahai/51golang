package main

import (
	//. "./app/controllers"
	"./app/controllers"
	"./app/middleware"
	//"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//加载模板
	router.LoadHTMLGlob("resources/views/*")
	//添加中间件
	router.Use(middleware.Middleware)
	// Initialize the routes
	//initializeRoutes()

	// Handle the index route
	router.GET("/", controllers.IndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", controllers.LoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", controllers.LoginAuth)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", controllers.LoginOut)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", controllers.RegisterPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", controllers.Register)
	}

	// Group article related routes together
	articleRoutes := router.Group("/article")
	{
		// Handle GET requests at /article/view/some_article_id
		articleRoutes.GET("/view/:article_id", controllers.ArticleDetailPage)

		// Handle the GET requests at /article/create
		// Show the article creation page
		// Ensure that the user is logged in by using the middleware
		articleRoutes.GET("/create", controllers.ArticlePage)

		// Handle POST requests at /article/create
		// Ensure that the user is logged in by using the middleware
		articleRoutes.POST("/create", controllers.ArticleCreate)
	}

	//注册接口
	//router.GET("/simple/server/get", GetHandler)
	//router.POST("/simple/server/post", PostHandler)
	//router.PUT("/simple/server/put", PutHandler)
	//router.DELETE("/simple/server/delete", DeleteHandler)

	//router.GET("/index", IndexHandler)
	//监听端口
	http.ListenAndServe(":1987", router)
}
