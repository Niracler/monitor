package v1

import (
	"net/http"

	"gamenews.niracler.com/monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetUserOperation(c *gin.Context) {
	page, _ := com.StrTo(c.DefaultQuery("p", "1")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("ps", "10")).Int()
	page = (page - 1) * pageSize
	maps := make(map[string]interface{})
	count, uoList := service.GetUserOperation(page, pageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"results": uoList,
	})
}

func GetVisitCount(c *gin.Context) {
	page, _ := com.StrTo(c.DefaultQuery("p", "1")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("ps", "10")).Int()
	page = (page - 1) * pageSize
	maps := make(map[string]interface{})
	count, vcList := service.GetVisitorCount(page, pageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"results": vcList,
	})
}
