package analysis

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marekparafianowicz/go-server/sites"
)

// Create crawls Site, fetch all H1 tags and render them
func Create(c *gin.Context) {
	id := c.Param("id")
	site, err := sites.FindSite(id)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record with this ID is not found"})
		return
	case err != nil:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tags := downloader(site.URL)
	fmt.Println(tags)

	c.JSON(200, tags)
}
