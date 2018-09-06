package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type UpdateSite struct {
	r repository.Repository
}

func NewUpdateSite(r repository.Repository) *UpdateSite {
	return &UpdateSite{r}
}

func (us *UpdateSite) Handle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID can't be nil")
		return
	}

	atr := atr{}
	c.BindJSON(&atr)
	if atr.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is nil"})
		return
	}

	site, err := us.r.UpdateSite(id, atr.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, site)
}
