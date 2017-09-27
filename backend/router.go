package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route := gin.Default()
	http.Handle("/", route)

	route.GET("/", func(g *gin.Context){
		g.String(http.StatusOK, "I'm go-appengine.")
	})
}
