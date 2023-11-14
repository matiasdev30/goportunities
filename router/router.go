package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(){
	//Init gin with deflaut config
	r := gin.Default()
	//Router root
	InitializeRoutes(r)
	//rodando o gin
	r.Run()
}

