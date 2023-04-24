package controllers

import (
	"example/go-mysql/config"
	"example/go-mysql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
* ListArticles :: Query for list from the Articles table
 */
func ListArticles(context *gin.Context) {
	var articles []models.Article

	if result := config.Db.Preload("Comments").Find(&articles); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	context.JSON(http.StatusOK, gin.H{"data": articles})
}

/*
* GetArticleById :: Query for list from the Articles table
 */
func GetArticleById(context *gin.Context) {
	pathId := context.Param("id")
	var article models.Article

	// validate incoming path id as a valid uint
	articleId, err := strconv.ParseUint(pathId, 0, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article id"})
	}

	// Find the article by ID in the database
	findErr := config.Db.Preload("Comments").Where(&models.Article{ArticleID: uint(articleId)}).First(&article).Error

	if findErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": article})
}

/*
* CreateArticle :: Insert a new Article record into the Articles table
 */
func CreateArticle(context *gin.Context) {
	var body models.ArticleCreateBody

	// validate incoming JSON body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create insert object from the parsed body struct
	article := models.Article{
		Title:       body.Title,
		Subtitle:    body.Subtitle,
		Body:        body.Body,
		CreatedAt:   body.CreatedAt,
		PublishedAt: body.PublishedAt,
	}

	// insert into the DB
	if result := config.Db.Create(&article); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// response
	context.JSON(http.StatusOK, gin.H{"data": article})
}

/*
* UpdateArticle :: Update an existing article record
 */
func UpdateArticle(context *gin.Context) {
	var body models.ArticleUpdateBody

	// validate incoming JSON body
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check article exists in DB
	pathId := context.Param("id")

	// validate incoming path id as a valid uint
	articleId, err := strconv.ParseUint(pathId, 0, 64)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Invalid article id"})
	}

	// Check article exists
	var article models.Article
	if err := config.Db.Where("article_id = ?", articleId).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
		return
	}

	// create insert object from the parsed body struct
	if result := config.Db.Model(&article).Updates(&body); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// response
	context.JSON(http.StatusOK, gin.H{"data": article})
}

/*
* DeleteArticle :: Delete an article and associated comments
 */
func DeleteArticle(context *gin.Context) {

	// check article exists in DB
	pathId := context.Param("id")

	// validate incoming path id as a valid uint
	articleId, err := strconv.ParseUint(pathId, 0, 64)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Invalid article id"})
	}

	// Check article exists
	var article models.Article
	if err := config.Db.Where(&models.Article{ArticleID: uint(articleId)}).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Article not found"})
		return
	}

	// Delete article with cascade to comments
	if result := config.Db.Select("Comments").Delete(&article); result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	// response
	context.JSON(http.StatusOK, gin.H{"data": articleId})
}
