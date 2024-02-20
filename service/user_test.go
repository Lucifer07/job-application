package service_test

import (
	"context"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/service"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceImp_Register(t *testing.T) {
	t.Run("should return id if success", func(t *testing.T) {
		// given
		password := "test"
		ctx := new(context.Context)
		idTarget := 1
		user := entity.User{}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockHelper.On("HashPassword", mock.Anything, mock.Anything).Return([]byte(password), nil)
		mockUser.On("Register", mock.Anything, mock.Anything).Return(&idTarget, nil)
		id, _ := userService.Register(*ctx, user)
		// then
		assert.Equal(t, idTarget, *id)
	})
	t.Run("should return err internal server error if failed hash password", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockHelper.On("HashPassword", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_, err := userService.Register(*ctx, user)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
	t.Run("should return err internal server error if failed hash register", func(t *testing.T) {
		// given
		password := "test"
		ctx := new(context.Context)
		user := entity.User{}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockHelper.On("HashPassword", mock.Anything, mock.Anything).Return([]byte(password), nil)
		mockUser.On("Register", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_, err := userService.Register(*ctx, user)
		assert.Equal(t, util.ErrorInternal, err)
	})
}

func TestUserServiceImp_Login(t *testing.T) {
	t.Run("should return jwt if success", func(t *testing.T) {
		// given
		jwtTarget := "test"
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(&user, nil)
		mockHelper.On("CheckPassword", mock.Anything, mock.Anything).Return(true, nil)
		mockHelper.On("CreateAndSign", mock.Anything).Return(jwtTarget, nil)
		jwt, _ := userService.Login(*ctx, user)
		// then
		assert.Equal(t, jwtTarget, *jwt)
	})
	t.Run("should return err internal server error if failed check login", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(nil, util.ErrorInternal)
		_, err := userService.Login(*ctx, user)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
	t.Run("should return error user not found if data nil", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(nil, nil)
		_, err := userService.Login(*ctx, user)
		// then
		assert.Equal(t, util.ErrorUserNotFound, err)
	})
	t.Run("should return interal server error if failed check password", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(&user, nil)
		mockHelper.On("CheckPassword", mock.Anything, mock.Anything).Return(false, util.ErrorInternal)
		_, err := userService.Login(*ctx, user)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
	t.Run("should return error wrong password if result not true", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(&user, nil)
		mockHelper.On("CheckPassword", mock.Anything, mock.Anything).Return(false, nil)
		_, err := userService.Login(*ctx, user)
		// then
		assert.Equal(t, util.ErrorWrongPassword, err)
	})
	t.Run("should return internal server error if error when create sign jwt", func(t *testing.T) {
		// given
		ctx := new(context.Context)
		user := entity.User{Password: "test"}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		userService := service.NewUserService(mockUser, mockHelper)
		// when
		mockUser.On("Login", mock.Anything, mock.Anything).Return(&user, nil)
		mockHelper.On("CheckPassword", mock.Anything, mock.Anything).Return(true, nil)
		mockHelper.On("CreateAndSign", mock.Anything).Return("", util.ErrorInternal)
		_, err := userService.Login(*ctx, user)
		// then
		assert.Equal(t, util.ErrorInternal, err)
	})
}

func TestNewUserService(t *testing.T) {
	t.Run("should return jwt if success", func(t *testing.T) {
		// given
		targetService := service.UserServiceImp{}
		mockUser := new(mocks.UserRepository)
		mockHelper := new(mocks.HelperInf)
		// when
		userService := service.NewUserService(mockUser, mockHelper)
		// then
		assert.IsType(t, targetService, *userService)
	})
}
