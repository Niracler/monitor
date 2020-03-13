package v1

import (
	"log"
	"net/http"

	"gamenews.niracler.com/monitor/model"
	"gamenews.niracler.com/monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetGames(c *gin.Context) {
	page, _ := com.StrTo(c.DefaultQuery("p", "1")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("ps", "16")).Int()
	page = (page - 1) * pageSize
	maps := make(map[string]interface{})
	count, games := service.GetGames(page, pageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"count":   count,
		"results": games,
	})
}

func AddGame(c *gin.Context) {
	game := &model.Game{}

	err := c.BindJSON(game)
	if err != nil {
		log.Panicln(err)
	}

	service.AddGame(game)

	c.JSON(http.StatusCreated, game)
}

func EditGame(c *gin.Context) {
}

func DeleteGame(c *gin.Context) {
}
