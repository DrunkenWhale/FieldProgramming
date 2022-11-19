package main

import (
	"FieldProgramming/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	api := g.Group("/api")
	{
		api.POST("/conv", router.ConvertService)
	}
	err := g.Run(":3777")
	if err != nil {
		panic(err)
		return
	}
}
