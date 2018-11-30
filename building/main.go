package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hal-ms/driver/building/cnto"
)

func main() {
	r := gin.Default()

	r.POST("/:scene", cnto.Led)
	//r.GET("/:scene", cnto.Building)
	//r.POST("/:scene", cnto.Scene)
	r.Run(":8000")
}
