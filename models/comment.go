package models

import "time"

type Comment struct {
	CommentID uint      `json:"id" gorm:"primary_key"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	ArticleID uint      `json:"-"`
}

type CommentCreateBody struct {
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required" time_format:"2006-01-02"`
}
