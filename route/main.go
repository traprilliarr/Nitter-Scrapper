package route

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/eyenote-corp/nitter-scrapper/controller"
)

func MainRoute(r *gin.Engine) {
	//cont := controller.NewArmadaController(serv)

	r.POST("/:chatId", controller.ScrapData)
	//Armada(r, servArmada)
}

//func Armada(r *gin.Engine, serv service.ArmadaService) {
//	cont := controller.NewArmadaController(serv)
//	route := r.Group("/armada")
//	route.GET("/download", cont.DownloadArmadaData)
//	route.GET("/:id", cont.GetArmadaDetails)
//	route.GET("", cont.GetArmadaData)
//}
