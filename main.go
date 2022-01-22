// go-plate-calc provides an API for calculating how much weight you need
// to load on each side of a weight bar to get to your working weight. Your
// working weight is the weight you are lifting.

// For example, if you want to lift 100lbs (this is your working weight) and
// you are using a standard 45lb olympic bar, then you need to load
// (100 - 45) / 2 = 27.5lbs on both sides of the bar.

// Assuming the following plate weights are available to you: 2.5, 5, 10, 20
// go-plate-calc tells you the plates that make up 27.5lbs, from heaviest to
// lightest: 20, 5, 2.5

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/plates", calculatePlates)

	router.Run(":8080")
}

func calculatePlates(c *gin.Context) {
	plates := [...]float32{20, 5, 2.5}
	c.JSON(http.StatusOK, gin.H{"plates": plates})
}
