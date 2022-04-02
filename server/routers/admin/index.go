package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*


 */

//handel
func waht(c *gin.Context) {
	c.String(http.StatusOK, "waht")

}

type create_json struct {
	Url    string `json:"url"`
	Detail string
	Rank   int
}

func hjson(c *gin.Context) {

	/*
		// 1.go语言封装一个json数据
		map[string]interface{}{
			"name": "123",
			"message": "谦虚、温柔、长得好看。",
			"age": 23
		}
	*/
	data := gin.H{
		"url":    "google.com",
		"detail": "nice search engin",
		"rank":   2,
	}
	c.JSON(http.StatusOK, data)
}

func create_struct(c *gin.Context) {

	data := create_json{
		Url:    "google.com",
		Detail: "nice search engin",
		Rank:   2,
	}
	c.JSON(http.StatusOK, data)
}
func Admin(e *gin.Engine) {
	//e := gin.Default()
	e.GET("/admin", waht)
	e.GET("/admin/json", hjson)
	e.GET("/admin/create_struct", create_struct)

	//return r

}
