package main

import (
	"FieldProgramming/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	g := gin.Default()
	api := g.Group("/api")
	api.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	{
		api.POST("/conv", router.ConvertService)
	}
	err := g.Run(":3777")
	if err != nil {
		panic(err)
		return
	}
}
