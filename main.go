package main

import (
	"dev_recruitment_crawler/provider"
	"dev_recruitment_crawler/provider/jumpit"
	"dev_recruitment_crawler/provider/wanted"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	provider provider.Provider
}

func main() {
	engine := gin.New()
	j := jumpit.NewJumpit()
	w := wanted.NewWanted()
	e := &Engine{
		provider: j,
	}

	e2 := &Engine{
		provider: w,
	}
	engine.GET("/", func(ctx *gin.Context) {
		resp := e.provider.GetRecruitment(1, "backend")
		resp2 := e2.provider.GetRecruitment(1, "backend")
		resp = append(resp, resp2...)

		ctx.JSON(200, resp)
	})

	engine.Run(":3000")
}
