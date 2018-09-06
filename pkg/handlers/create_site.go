package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type CreateSite struct {
	r repository.Repository
}

func NewCreateSite(r repository.Repository) *CreateSite {
	return &CreateSite{r}
}

type atr struct {
	URL string `json:"url" binding:"required"`
}

func (cs *CreateSite) Handle(c *gin.Context) {
	atr := atr{}

	c.BindJSON(&atr)
	if atr.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is nil"})
		return
	}

	site, err := cs.r.CreateSite(atr.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, site)
}
