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

// UserExists user exists
func UserExists(c *gin.Context) {

}

// AddUser add user
func AddUser(c *gin.Context) {
	var user models.User
	_ = c.ShouldBindJSON(&user)

	if user.Username == "" {
		log.Println("username is required")
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "the username is required",
		})
		return
	}
	if user.Password == "" {
		log.Println("password is required")
		c.JSON(http.StatusOK, gin.H{
			"code": code.ERROR,
			"msg":  "the password is required",
		})
		return
	}
	// check if user exists
	resultCode := models.CheckUserExists(user.Username)
	if resultCode == code.SUCCESS {
		// crypto password
		user.Password = utils.GetMD5(user.Password)
		resultCode = models.CreateUser(user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": user,
		"msg":  code.GetMsg(resultCode),
	})
}

// GetUsers get users
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", utils.Page))
	size, err := strconv.Atoi(c.DefaultQuery("size", utils.Size))
	if err != nil {
		log.Println("GetUsers page strconv Error: ", err)
		return
	}
	page = (page - 1) * size
	resultCode, user := models.GetUserList(page, size)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": user,
		"msg":  code.GetMsg(resultCode),
	})
}

// EditUser edit user
func EditUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("EditUser id strconv Error: ", err)
		return
	}
	var user models.User
	_ = c.ShouldBindJSON(&user)
	// check if user exists
	resultCode := models.CheckUserExists(user.Username)
	if resultCode == code.SUCCESS {
		resultCode = models.EditUser(id, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"data": user,
		"msg":  code.GetMsg(resultCode),
	})
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("DeleteUser id strconv Error: ", err)
		return
	}
	resultCode := models.DeleteUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": resultCode,
		"msg":  code.GetMsg(resultCode),
	})
}
