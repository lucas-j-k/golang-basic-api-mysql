package models

import "time"

type Article struct {
	ArticleID   uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"createdAt"`
	PublishedAt time.Time `json:"publishedAt"`
	Comments    []Comment `json:"comments" gorm:"foreignkey:ArticleID"`
}

type ArticleCreateBody struct {
	Title       string    `json:"title" binding:"required"`
	Subtitle    string    `json:"subtitle" binding:"required"`
	Body        string    `json:"body" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required" time_format:"2006-01-02"`
	PublishedAt time.Time `json:"publishedAt" binding:"required" time_format:"2006-01-02"`
}

type ArticleUpdateBody struct {
	Title       string    `json:"title" binding:"required"`
	Subtitle    string    `json:"subtitle" binding:"required"`
	Body        string    `json:"body" binding:"required"`
	PublishedAt time.Time `json:"publishedAt" binding:"required" time_format:"2006-01-02"`
}
