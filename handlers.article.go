package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles},
		"index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleById(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"Error Code":    http.StatusNotFound,
				"Error Message": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Error Code":    http.StatusNotFound,
			"Error Message": err.Error(),
		})
	}
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
