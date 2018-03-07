package main

import (
	"math"

	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"net/http"
)


func magicNum(x float64) float64 {
	for j:=0; j<=250; j++ {
		for i:=0; i<= 1000000; i++ {
			x += math.Sqrt(x)
		}
	}
	return x
}

func main() {
	r := gin.Default()
	r.GET("/v1/number/:num", func(c *gin.Context) {
		// parse number from GET parameter
		x, err := strconv.ParseFloat(c.Param("num"), 64)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// get magic number
		num := magicNum(x)

		// return value
		c.JSON(http.StatusOK, gin.H{
			"value": fmt.Sprintf("%.100f", num),
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
