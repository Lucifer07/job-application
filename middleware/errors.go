package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

func CustomMiddlewareError(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {

		firstError := c.Errors[0].Err
		errResponse := checkError(firstError)
		if errResponse.StatusCode != http.StatusInternalServerError {
			c.JSON(errResponse.StatusCode, errResponse)
			return
		}
		c.AbortWithStatusJSON(errResponse.StatusCode, errResponse)
		return
	}

}

func checkError(err error) response.ResponseMsgErr {
	switch err {
	case util.ErrorUserNotFound:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorUserNotFound.Error()}
	case util.ErrorWrongPassword:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorWrongPassword.Error()}
	case util.ErrorBadRequest:
		return response.ResponseMsgErr{StatusCode: http.StatusBadRequest, Message: util.ErrorBadRequest.Error()}
	default:
		return response.ResponseMsgErr{StatusCode: http.StatusInternalServerError}
	}
}
