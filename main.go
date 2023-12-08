package main

import (
	"net/http"
	"web/model"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.GET("/index", IndexController)

	r.Run(":10000")
}
func IndexController(c *gin.Context) {
	user := model.GetUserInfo("1")
	username := user[0].User_name
	c.HTML(http.StatusOK, "index.html", gin.H{"username": username})

}
