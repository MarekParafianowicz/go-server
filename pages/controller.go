package pages

import "github.com/gin-gonic/gin"

// Index is root path in this app
func Index(c *gin.Context) {
	m := message{"Hello in API"}
	c.JSON(200, m)
}
