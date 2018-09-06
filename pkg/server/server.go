package server

import (
	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/handlers"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type server struct {
	r repository.Repository
}

func New(r repository.Repository) server {
	return server{r}
}

func (s server) Run() {
	router := gin.Default()

	router.GET("/sites", handlers.NewIndexSites(s.r).Handle)
	router.GET("/sites/:id", handlers.NewShowSite(s.r).Handle)
	router.POST("/sites", handlers.NewCreateSite(s.r).Handle)
	router.PUT("/sites/:id", handlers.NewUpdateSite(s.r).Handle)
	router.DELETE("/sites/:id", handlers.NewDeleteSite(s.r).Handle)

	router.POST("/sites/:id/analysis", handlers.NewConductAnalysis(s.r).Handle)

	router.Run()
}
