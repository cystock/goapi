package miapi

import (
	"../../services/miapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	paramCountryId = "countryId"
)


func GetCountry (c *gin.Context)  {
	countryId := c.Param(paramCountryId)
	country, err := miapi.GetCountryFromApi(countryId)
	if err != nil{
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusOK, country)
}
