package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("view/template/*")
	router.StaticFS("/assets", http.Dir("view/dist/assets"))

	router.GET("/", index)

	port := ":8081"
	log.Printf("Running on port %s", port)

	if err := router.Run(port); err != nil {
		log.Fatal(err)

	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Version": "1.0.0",
	})
}
