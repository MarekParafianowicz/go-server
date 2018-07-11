package sites

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sitesMessage struct {
	Sites []site `json:"sites_attr"`
}

// Index renders all sites (TODO: pagination!)
func Index(c *gin.Context) {

	sites, err := allSites()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, sites)
}

// Show renders single site
func Show(c *gin.Context) {
	id := c.Param("id")
	site, err := findSite(id)

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

type atr struct {
	URL string `json:"url" binding:"required"`
}

// Create adds new site to DB
func Create(c *gin.Context) {
	atr := atr{}

	c.BindJSON(&atr)
	if atr.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is nil"})
		return
	}

	site, err := createSite(atr.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, site)
}

// Update edits single site
func Update(c *gin.Context) {
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

	site, err := updateSite(id, atr.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, site)
}

// Delete destroys row in DB
func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID can't be nil")
		return
	}

	err := deleteSite(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
