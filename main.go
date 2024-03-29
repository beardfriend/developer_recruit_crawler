package main

import (
	"html/template"
	"net/http"
	"strings"

	"dev_recruitment_crawler/engine"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	g.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	g.LoadHTMLGlob("templates/*.html")

	e := engine.NewEngine(nil)

	type Request struct {
		Position string `form:"position,default=dataEngineer" binding:"oneof=frontend backend dataEngineer"`
		Career   int    `form:"career,default=1"`
	}

	g.Static("/templates", "./templates/")

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
