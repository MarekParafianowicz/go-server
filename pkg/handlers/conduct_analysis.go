package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/pkg/analysis"
	"github.com/marekparafianowicz/go-server/pkg/repository"
)

type ConductAnalysis struct {
	r repository.Repository
}

func NewConductAnalysis(r repository.Repository) *ConductAnalysis {
	return &ConductAnalysis{r}
}

func (ca *ConductAnalysis) Handle(c *gin.Context) {
	id := c.Param("id")
	site, err := ca.r.FindSite(id)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record with this ID is not found"})
		return
	case err != nil:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags := analysis.Downloader(site.URL)
	fmt.Println(tags)

	c.JSON(200, tags)
}
