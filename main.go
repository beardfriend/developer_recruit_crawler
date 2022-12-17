package main

import (
	"sync"

	"dev_recruitment_crawler/model"
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
		var wg sync.WaitGroup
		resp := make([]*model.Recruitment, 0)
		wg.Add(2)
		go func() {
			defer wg.Done()
			resp = append(resp, e.provider.GetRecruitment(1, "backend")...)
		}()

		go func() {
			defer wg.Done()
			resp = append(resp, e2.provider.GetRecruitment(1, "backend")...)
		}()
		wg.Wait()

		ctx.JSON(200, resp)
	})

	engine.Run(":3000")
}
