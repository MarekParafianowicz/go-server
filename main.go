package main

import (
	"github.com/marekparafianowicz/go-server/pages"
	"github.com/marekparafianowicz/go-server/sites"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.GET("/", pages.Index)

	router.GET("/sites", sites.Index)
	router.GET("/sites/:id", sites.Show)
	router.POST("/sites", sites.Create)
	router.PUT("/sites/:id", sites.Update)
	router.DELETE("/sites/:id", sites.Delete)

	router.Run()
}
