package system

import (
	"encoding/json"
	"gin/model/common/response"
	"github.com/gin-gonic/gin"
)

type SystemApiApi struct{}

//
// TestApi
// @Description:
// @receiver s
// @param c
//
func (s *SystemApiApi)TestApi(c *gin.Context) {

	response.OkWithMessage("test api",c)

}

//
// TestParamApi
// @Description:
// @receiver s
// @param c
//
func (s *SystemApiApi)TestParamApi(c *gin.Context){

	username := c.DefaultQuery("username", "小王子")
	address := c.Query("address")

	testMap:=make(map[string]string,10)
	testMap["username"]=username
	testMap["address"]=address
	response.OkWithData(testMap,c)

}

//
// TestPostFormApi
// @Description:
// @receiver s
// @param c
//
func (s *SystemApiApi) TestPostFormApi(c *gin.Context) {

	username := c.PostForm("username")
	address := c.PostForm("address")

	testMap:=make(map[string]string,10)
	testMap["username"]=username
	testMap["address"]=address
	response.OkWithData(testMap,c)

}

//
// TestJsonApi
// @Description:
// @receiver s
// @param c
//
func (s *SystemApiApi) TestJsonApi(c *gin.Context) {

	b, _ := c.GetRawData()  // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)

	//c.JSON(http.StatusOK, m)
	response.OkWithData(m,c)

}

//
// TestRestfulApi
// @Description:
// @receiver s
// @param c
//
func (s *SystemApiApi) TestRestfulApi(c *gin.Context) {

	username := c.Param("username")
	address := c.Param("address")

	testMap:=make(map[string]string,10)
	testMap["username"]=username
	testMap["address"]=address
	response.OkWithData(testMap,c)
}


