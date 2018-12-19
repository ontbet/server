package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	. "ontbet/models"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func QueryForAddress(c *gin.Context) {
	address := c.Param("address")
	p := &GuessResult{}
	result, err := p.GetResultForAddress(address)
	var jsonstr []byte
	if err == nil {
		jsonstr, _ = json.Marshal(result)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": string(jsonstr),
		"count":  len(result),
	})
}
func QueryLastInfo(c *gin.Context) {

	p := &GuessResult{}
	result, err := p.GeTop20Result()
	var jsonstr []byte
	if err == nil {
		jsonstr, _ = json.Marshal(result)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": string(jsonstr),
		"count":  len(result),
	})
}
