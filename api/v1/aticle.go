package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/code"
	"goblog/models"
	"goblog/utils"
	"log"
	"net/http"
	"strconv"
)

// AddArticle add article
func AddArticle(c *gin.Context) {
	var article models.Article
	_ = c.ShouldBindJSON(&article)

	if article.Title == "" {
		log.Println("article name is required")
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "the article name is required",
		})
		return
	}
	resultCode := models.CreateArticle(article)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": article,
		"msg":  code.GetMsg(resultCode),
	})
}

// GetArticleList get article list
func GetArticleList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", utils.Page))
	size, err := strconv.Atoi(c.DefaultQuery("size", utils.Size))
	if err != nil {
		log.Println("GetArticleList page strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "GetArticleList page strconv Error",
		})
		return
	}
	page = (page - 1) * size
	resultCode, article := models.GetArticleList(page, size)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": article,
		"msg":  code.GetMsg(resultCode),
	})
}

// EditArticle edit article
func EditArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("EditArticle id strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "EditArticle id strconv Error",
		})
		return
	}
	var article models.Article
	_ = c.ShouldBindJSON(&article)
	resultCode := models.EditArticle(id, article)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": article,
		"msg":  code.GetMsg(resultCode),
	})
}

// DeleteArticle delete article
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("DeleteArticle id strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "DeleteArticle id strconv Error",
		})
		return
	}
	resultCode := models.DeleteArticleById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"msg":  code.GetMsg(resultCode),
	})
}
