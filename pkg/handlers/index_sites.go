package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type IndexSites struct {
	r repository.Repository
}

func NewIndexSites(r repository.Repository) *IndexSites {
	return &IndexSites{r}
}

func (is *IndexSites) Handle(c *gin.Context) {
	sites, err := is.r.AllSites()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, sites)
}
