package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type DeleteSite struct {
	r repository.Repository
}

func NewDeleteSite(r repository.Repository) *DeleteSite {
	return &DeleteSite{r}
}

func (ds *DeleteSite) Handle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID can't be nil")
		return
	}

	err := ds.r.DeleteSite(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
