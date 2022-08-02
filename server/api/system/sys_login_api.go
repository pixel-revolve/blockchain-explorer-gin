package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct {}

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginJson
// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
func (s *LoginApi) LoginJson(c *gin.Context)  {

	var login Login

	if err := c.ShouldBind(&login); err == nil {
		fmt.Printf("login info:%#v\n", login)
		c.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}


