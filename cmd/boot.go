package main

import (
	"FieldProgramming/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	api := g.Group("/api")
	api.Use(cors.Default())
	{
		api.POST("/conv", router.ConvertService)
	}
	err := g.Run(":3777")
	if err != nil {
		panic(err)
		return
	}
}
