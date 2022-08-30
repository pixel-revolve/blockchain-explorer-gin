package web

import "github.com/gin-gonic/gin"

func LaunchWeb(r *gin.Engine) {
	r.LoadHTMLFiles("web/templates/*/*")
	r.LoadHTMLGlob("web/templates/explorer.html")
	r.Static("static", "web/templates/static")
	r.Static("svg", "web/templates/svg")

	r.GET("/explorer", func(context *gin.Context) {
		context.HTML(200, "explorer.html", nil)
	})
}
