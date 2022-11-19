package main

import (
	"FieldProgramming/middleware"
	"FieldProgramming/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	api := g.Group("/api")
	api.Use(middleware.Cors())
	{
		api.POST("/conv", router.ConvertService)
	}
	err := g.Run(":3777")
	if err != nil {
		panic(err)
		return
	}
}
