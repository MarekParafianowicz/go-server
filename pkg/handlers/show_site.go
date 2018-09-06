package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type ShowSite struct {
	r repository.Repository
}

func NewShowSite(r repository.Repository) *ShowSite {
	return &ShowSite{r}
}

func (ss *ShowSite) Handle(c *gin.Context) {
	id := c.Param("id")
	site, err := ss.r.FindSite(id)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record with this ID is not found"})
		return
	case err != nil:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, site)
}
