// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)
	if loggedIn {
		render(c, gin.H{"title": "Home Page", "payload": articles, "is_logged_in": true}, "index.html")
	} else {
		render(c, gin.H{"title": "Home Page", "payload": articles, "is_logged_in": false}, "index.html")
	}

}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getSingleArticle(articleID); err == nil {
			render(c, gin.H{"title": article.Title, "message": article, "is_logged_in": true}, "article.html")
		} else {
			render(c, gin.H{"title": "404 Not Found", "message": http.StatusNotFound}, "404.html")
		}

	} else {
		render(c, gin.H{"title": "404 Not Found", "message": "Oops! Sorry we cant find any article with the ID"}, "404.html")
	}
}

func getArticleUnAuthenticated(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getSingleArticle(articleID); err == nil {
			render(c, gin.H{"title": article.Title, "message": article}, "article.html")
		} else {
			render(c, gin.H{"title": "404 Not Found", "message": http.StatusNotFound}, "404.html")
		}

	} else {
		render(c, gin.H{"title": "404 Not Found", "message": "Oops! Sorry we cant find any article with the ID"}, "404.html")
	}
}

func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title":        "Create New Article",
		"is_logged_in": true}, "create-article.html")
}
func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":        "Submission Successful",
			"payload":      a,
			"is_logged_in": true}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
