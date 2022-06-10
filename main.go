package main

import (
	"log"
	"net/http"
	"os"
	//"net/url"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
//	router.LoadHTMLGlob("templates/**/**/*.tmpl.html")
	router.Static("/static", "static")
	
	//router.GET("/login", ShowLogin) // login page
	
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		})
	
	/*
	router.NoRoute(func(c *gin.Context) {
		// set location
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusFound, location.RequestURI())

	})
	*/
	router.Run(":" + port)
}
/*
func ShowLogin(c *gin.Context) {
	cc := gin.H{"title": "หน้าแรก"}
	c.HTML(http.StatusOK, "teml/teml/login.tmpl.html", cc)

	c.Abort()

}
*/
