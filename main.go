package main

import (
	"dev_recruitment_crawler/engine"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	e := engine.NewEngine()
	g.GET("/", func(ctx *gin.Context) {
		resp := e.GetRecruitment()
		ctx.JSON(200, resp)
	})

	g.Run(":3000")
}
