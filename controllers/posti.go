package controllers

import (
	// "log"
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DB represents the database connection object.
// First() is a method provided by GORM to retrieve the first record that matches the specified conditions.
// The &post argument is a pointer to the variable where the retrieved data will be stored.
// id is the value used to filter the records. In this case, it's typically the ID of the post you're trying to find.

func FindPosts(c *gin.Context) {
	id := c.Param("id")
	// id := c.Query("id")
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id") // Get the ID of the post to update

	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update post fields
	post.Title = input.Title
	post.Content = input.Content

	// Save the updated post to the database
	if err := models.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := models.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
