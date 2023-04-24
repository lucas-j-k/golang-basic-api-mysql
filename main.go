package main

import (
	"example/go-mysql/config"
	"example/go-mysql/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	router := gin.Default()

	// healthcheck
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "pong"})
	})

	// Articles
	router.GET("/articles", controllers.ListArticles)
	router.GET("/articles/:id", controllers.GetArticleById)
	router.POST("/articles", controllers.CreateArticle)
	router.PUT("/articles/:id", controllers.UpdateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)

	// Comments
	router.POST("/articles/:id/comments", controllers.CreateComment)

	// connect DB
	err := config.InitDB()
	if err != nil {
		panic("MySQL connection failed")
	}

	port := viper.Get("PORT").(string)

	// run server
	router.Run(fmt.Sprintf(":%v", port))
}
