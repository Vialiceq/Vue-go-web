package main

import (
	"server/db"

	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	//"net/http"
	"server/routers/admin"
	"server/routers/web"
)

type Option func(*gin.Engine)

var options = []Option{}

//include routers
func include(opts ...Option) {
	options = append(options, opts...)
}

func arrange() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}

func main() {

	include(admin.Admin, web.Web, db.Main)
	r := arrange()

	err := r.Run(":7777")
	if err != nil {
		return
	}
}
