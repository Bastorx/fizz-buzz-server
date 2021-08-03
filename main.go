package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := setupServer()
	err := r.Run(":4000")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func setupServer()  *gin.Engine {
	r := gin.Default()
	initRouter(r)
	return r
}