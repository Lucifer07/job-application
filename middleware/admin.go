package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

func AuthorizationAdmin(c *gin.Context) {
	data := c.Value("data").(map[string]string)
	if data["role"] != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, response.ResponseMsgErr{Message: util.ErrorForbidden.Error()})
		return
	}
	c.Next()
}
