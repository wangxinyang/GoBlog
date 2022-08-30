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

// AddCategory add category
func AddCategory(c *gin.Context) {
	var category models.Category
	_ = c.ShouldBindJSON(&category)

	if category.Name == "" {
		log.Println("category name is required")
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "the category name is required",
		})
		return
	}
	// check if user exists
	resultCode := models.CheckCategoryExists(category.Name)
	if resultCode == code.SUCCESS {
		resultCode = models.CreateCategory(category)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": category,
		"msg":  code.GetMsg(resultCode),
	})
}

// GetCategoryList get category list
func GetCategoryList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", utils.Page))
	size, err := strconv.Atoi(c.DefaultQuery("size", utils.Size))
	if err != nil {
		log.Println("GetCategoryList page strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "GetCategoryList page strconv Error",
		})
		return
	}
	page = (page - 1) * size
	resultCode, category := models.GetCategoryList(page, size)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": category,
		"msg":  code.GetMsg(resultCode),
	})
}

// EditCategory edit category
func EditCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("EditCategory id strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "EditCategory id strconv Error",
		})
		return
	}
	var category models.Category
	_ = c.ShouldBindJSON(&category)
	// check if user exists
	resultCode := models.CheckCategoryExists(category.Name)
	if resultCode == code.SUCCESS {
		resultCode = models.EditCategory(id, category)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": category,
		"msg":  code.GetMsg(resultCode),
	})
}

// DeleteCategory delete category
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("DeleteCategory id strconv Error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "DeleteCategory id strconv Error",
		})
		return
	}
	resultCode := models.DeleteCategoryById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"msg":  code.GetMsg(resultCode),
	})
}
