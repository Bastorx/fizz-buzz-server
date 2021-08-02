package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	initRouter(r)
	err := r.Run(":4000")
	if err != nil {
		return
	}
}
