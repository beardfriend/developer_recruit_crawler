package main

import (
	"html/template"
	"net/http"
	"strings"

	"dev_recruitment_crawler/config"
	"dev_recruitment_crawler/engine"
	"dev_recruitment_crawler/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	config := config.LoadConfig()

	// Mongodb
	mongo := infrastructure.NewMongodb(config.MongoUrl)

	defer func() {
		infrastructure.CloseMongodb(mongo)
	}()

	g := gin.New()

	g.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	g.LoadHTMLGlob("templates/*.html")

	e := engine.NewEngine(mongo)
	g.GET("/", func(ctx *gin.Context) {
		resp := e.GetRecruitment()
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"recruitments": resp,
		})
	})

	g.GET("/cron", func(ctx *gin.Context) {
		e.CronRecruitment()
		ctx.JSON(200, nil)
	})

	g.Run(":3000")
}
