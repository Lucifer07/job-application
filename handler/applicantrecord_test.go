package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewApplicantHandler(t *testing.T) {
	t.Run("should return applicant handler", func(t *testing.T) {
		// given
		targetHandler := handler.ApplicantHandler{}
		mockService := new(mocks.AppicantRecordService)
		// when
		handler := handler.NewApplicantHandler(mockService)
		assert.IsType(t, targetHandler, *handler)
	})

}

func TestApplicantHandler_AppliedJob(t *testing.T) {
	t.Run("should return 201 if success", func(t *testing.T) {
		// given
		appRecord := dto.ApplicantRecord{}
		data := map[string]string{"id": "1", "role": "user"}
		route := gin.Default()
		mockService := new(mocks.AppicantRecordService)
		handlerApp := handler.NewApplicantHandler(mockService)

		// when
		mockService.On("AppliedJob", mock.Anything, mock.Anything).Return(&appRecord, nil)
		route.POST("/jobs/:id", func(c *gin.Context) {
			c.Set("data", data)
			handlerApp.AppliedJob(c)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/jobs/1", nil)
		route.ServeHTTP(w, req)

		// then
		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})
}
