package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"dev_recruitment_crawler/engine"
	"dev_recruitment_crawler/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	g := gin.New()

	g.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	g.LoadHTMLGlob("templates/*.html")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// mongodb
	mongo := infrastructure.NewMongodb(os.Getenv("MONGODB_URL"))
	defer func() {
		infrastructure.CloseMongodb(mongo)
	}()
	e := engine.NewEngine(mongo)

	type Request struct {
		Position string `form:"position,default=backend" binding:"oneof=frontend backend"`
		Career   int    `form:"career,default=1"`
	}

	go e.CronRecruitment()

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Hour().Do(e.CronRecruitment)
		<-s.Start()
	}()

	g.GET("/", func(ctx *gin.Context) {
		req := new(Request)

		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.HTML(http.StatusBadRequest, "400.html", nil)
			return
		}

		resp := e.GetRecruitment(req.Career, req.Position)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"recruitments": resp,
		})
	})

	g.Run(":2000")
}
