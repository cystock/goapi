package miapi

import (
	"../../services/miapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramSiteId = "siteId"
)

func GetSite (c *gin.Context)  {
	siteId := c.Param(paramSiteId)
	site, err := miapi.GetSiteFromApi(siteId)
	if err != nil{
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusOK, site)
}

/*func GetSites (c *gin.Context)  {
	sites, err := miapi.GetSitesFromApi()
	if err != nil{
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusOK, sites)
}*/


