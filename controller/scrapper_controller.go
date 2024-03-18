package controller

import (
	"fmt"
	arutek "github.com/arutek/backend-go-package"
	"github.com/gin-gonic/gin"
	"gitlab.com/eyenote-corp/nitter-scrapper/dto"
	"gitlab.com/eyenote-corp/nitter-scrapper/service"
	"net/http"
	"os"
)

func ScrapData(ctx *gin.Context) {
	var inputData dto.ScrapperDto
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H(arutek.Error(err, err.Error())))
		return
	}
	domain := os.Getenv("NITTER_INSTANCE")
	url := fmt.Sprintf("%s/search?f=tweet&q=from:%s", domain, inputData.Account)
	obj, err := service.Scrape(url, ctx.Param("chatId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H(arutek.Error(err, err.Error())))
		return
	}
	res := arutek.Response("SAVED", obj, -1)
	ctx.JSON(http.StatusOK, gin.H(res))
}
