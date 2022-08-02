package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/model/platform"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io/ioutil"
)

type AssetTransferService struct {
}

func (assetTransferService *AssetTransferService) PutAsset(ctx *gin.Context) {
	data := struct {
		Id   string `json:"id"`
		Data string `json:"data"`
	}{}
	err := ctx.Bind(&data)
	if err != nil {
		return
	}
	result := utils.SubmitTransaction("CreateAsset", cast.ToString(data.Id), cast.ToString(data.Data))
	var response platform.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return
	}
	ctx.JSON(200, response)
}

func (assetTransferService *AssetTransferService) ReadAsset(ctx *gin.Context) {
	id := ctx.Query("id")
	result := utils.SubmitTransaction("ReadAsset", id)
	var response platform.Response
	err := json.Unmarshal(result, &response)
	if err != nil {
		return
	}
	ctx.JSON(200, response)
}

func (assetTransferService *AssetTransferService) ReadAssetHistory(ctx *gin.Context) {
	id := ctx.Query("id")
	result := utils.SubmitTransaction("GetAssetHistory", id)
	var response platform.Response
	err := json.Unmarshal(result, &response)
	if err != nil {
		return
	}
	ctx.JSON(200, response)
}

func (assetTransferService *AssetTransferService) QueryAssetByAssetData(ctx *gin.Context) {
	var bodyData map[string]interface{}
	var buffer bytes.Buffer
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	_type := ctx.Query("type")
	err := json.Unmarshal(data, &bodyData)
	if err != nil {
		return
	}
	buffer.WriteString(`{"selector":`)
	if _type == "and" {
		buffer.WriteString(`{"$and":[`)

	} else if _type == "or" {
		buffer.WriteString(`{"$or":[`)
	}
	for key, value := range bodyData {
		_regex := fmt.Sprintf(`{"%s":{"$regex": "%s"}},`, key, value)
		buffer.WriteString(_regex)
	}
	buffer.Truncate(buffer.Len() - 1)
	buffer.WriteString(`]}}`)
	fmt.Println(buffer.String())
	result := utils.SubmitTransaction("QueryAssets", buffer.String())
	var response platform.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return
	}
	ctx.JSON(200, response)
}
