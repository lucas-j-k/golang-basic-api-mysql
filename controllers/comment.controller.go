package controllers

import (
	"example/go-mysql/config"
	"example/go-mysql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
*	CreateComment :: Insert new comment attached to an article ID
 */
func CreateComment(context *gin.Context) {
	var body models.CommentCreateBody

	// validate incoming JSON body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pathId := context.Param("id")

	// validate incoming path id as a valid uint
	articleId, err := strconv.ParseUint(pathId, 0, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
	}

	// Check article exists
	var article models.Article
	if err := config.Db.Where("article_id = ?", articleId).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
		return
	}

	// create insert object from the parsed body struct
	newComment := models.Comment{
		Body:      body.Body,
		CreatedAt: body.CreatedAt,
		ArticleID: article.ArticleID,
	}

	// insert into the DB
	if result := config.Db.Create(&newComment); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// response
	context.JSON(http.StatusOK, gin.H{"data": article})
}
