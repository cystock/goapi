package miapi

import (
	"../../services/miapi"
	"../../utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	paramUserId = "userId"
)

func GetUser( c *gin.Context)  {
	userId := c.Param(paramUserId)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil{
		apiError := &apierrors.ApiError{
			err.Error(),
			http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)
		return
	}

	user, apiError := miapi.GetUserFromApi(id)
	if apiError != nil{
		c.JSON(apiError.Status, apiError.Message)
		return
	}
	c.JSON(http.StatusOK, user)
}
