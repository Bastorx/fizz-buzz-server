package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FizzBuzzParams struct {
	Int1	int `form:"int1" json:"int1" xml:"int1"  binding:"required"`
	Int2	int `form:"int2" json:"int2" xml:"int2" binding:"required"`
	Limit 	int `form:"limit" json:"limit" xml:"limit" binding:"required"`
	Str1	string `form:"str1" json:"str1" xml:"str1" binding:"required"`
	Str2	string `form:"str2" json:"str2" xml:"str2" binding:"required"`
}

func FizzBuzz(c *gin.Context) {
	var json FizzBuzzParams
	var res string

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 1; i <= json.Limit; i++ {
		if i % (json.Int1 * json.Int2) == 0 {
			res = res + json.Str1 + json.Str2
		} else if i % json.Int1 == 0 {
			res = res + json.Str1
		} else if i % json.Int2 == 0 {
			res = res + json.Str2
		} else {
			res = res + strconv.Itoa(i)
		}
		if i < json.Limit {
			res = res + ","
		}
	}

	fmt.Println(res)
	c.String(200, res)
}