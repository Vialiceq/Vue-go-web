package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func waht(c *gin.Context) {
	c.String(http.StatusOK, "waht")

}

func Web(e *gin.Engine) {
	//e := gin.Default()
	e.GET("/web", waht)
}
