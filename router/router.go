package router

import "github.com/gin-gonic/gin"

func Initialize(){
	//Init gin with deflaut config
	r := gin.Default()
	//Router root
	r.GET("/", func(ctx *gin.Context){
		ctx.JSON(
			200,
			gin.H{
				"message" : "on",
			},
		)
	})
	//rodando o gin
	r.Run("7070",)
}