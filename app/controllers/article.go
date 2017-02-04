package controllers

import (
	"../models"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
)

// ================== article start ========================

func IndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	render(c, "index.html", gin.H{"title": "index", "payload": articles})
}

func ArticleDetailPage(c *gin.Context) {
	//根据完整文件名渲染模板，并传递参数
	//render(c, "article.html", gin.H{"title": "test article", "payload": "article content"})
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, "article.html", gin.H{"title": article.Title, "payload": article})
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}

}

func ArticlePage(c *gin.Context) {
	render(c, "create-article.html", gin.H{"title": "Create New Article"})
}

//func ArticleCreate(c *gin.Context) {
//render(c, "submission-successful.html", gin.H{"title": "Submission Successful"})
//}
func ArticleCreate(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := models.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		render(c, "submission-successful.html", gin.H{"title": "Submission Successful", "payload": a})
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, templateName string, data gin.H) {
	//loggedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

// ================== article end ========================
