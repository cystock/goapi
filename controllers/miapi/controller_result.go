package miapi

// aca va el wg de go routines


import (
	"../../services/miapi"
	"../../utils/apierrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	paramUser = "userId"
)

func GetResult( c *gin.Context)  {
	userId := c.Param(paramUser)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil{
		apiError := &apierrors.ApiError{
			err.Error(),
			http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)
		return
	}

	result, apiError := miapi.GetResultFromApi(id)
	if apiError != nil{
		c.JSON(apiError.Status, apiError.Message)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetResultGoroutine( c *gin.Context)  {
	userId := c.Param(paramUser)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil{
		apiError := &apierrors.ApiError{
			err.Error(),
			http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)
		return
	}

	result, apiError := miapi.GetResultGoroutineFromApi(id)
	if apiError != nil{
		c.JSON(apiError.Status, apiError.Message)
		return
	}
	c.JSON(http.StatusOK, result)
}


func GetResultChannel( c *gin.Context)  {
	userId := c.Param(paramUser)
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil{
		apiError := &apierrors.ApiError{
			err.Error(),
			http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)
		return
	}

	result := miapi.GetResultGChannelFromApi(id)
	c.JSON(http.StatusOK, result)
}