package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/gin-gonic/gin"
)

func MiddlewareJWTAuthorization(c *gin.Context) {
	const bearer = "Bearer "
	token := ""
	header := c.GetHeader("Authorization")
	if header != "" {
		token = header[len(bearer):]
	}
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseMsgErr{Message: util.ErrorUnauthorized.Error()})
		return
	}
	claims, err := util.ParseAndVerify(token)
	data:=make(map[string]string, 0)
	data["id"]=claims["id"].(string)
	data["role"]=claims["role"].(string)
	c.Set("data", data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseMsgErr{Message: util.ErrorUnauthorized.Error()})
		return
	}
	c.Next()
}
