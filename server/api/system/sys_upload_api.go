package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UploadApi struct {}

func (u *UploadApi) Upload(c *gin.Context) {

	// 单个文件
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	dst := fmt.Sprintf("D:/tempFiles/%s", file.Filename)
	// 上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})


}