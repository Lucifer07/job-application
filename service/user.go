package service

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/util"
)

type UserServiceImp struct {
	userRepo   repository.UserRepository
	helperTool util.HelperInf
}
type UserService interface {
	Register(ctx context.Context, user entity.User) (*int, error)
	Login(ctx context.Context, user entity.User) (*string, error)
}

func NewUserService(userRepo repository.UserRepository, helperRepo util.HelperInf) *UserServiceImp {
	return &UserServiceImp{
		userRepo:   userRepo,
		helperTool: helperRepo,
	}
}
func (s *UserServiceImp) Register(ctx context.Context, user entity.User) (*int, error) {
	passwordHash, err := s.helperTool.HashPassword(user.Password, 12)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHash)
	id, err := s.userRepo.Register(ctx, user)

	if err != nil {
		return nil, err
	}
	return id, nil
}
func (s *UserServiceImp) Login(ctx context.Context, user entity.User) (*string, error) {
	userData, err := s.userRepo.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	if userData != nil {
		passwordHash := []byte(userData.Password)
		result, err := s.helperTool.CheckPassword(user.Password, passwordHash)
		if err != nil {
			return nil, err
		}
		if !result {
			return nil, util.ErrorWrongPassword
		}
		jwt, err := s.helperTool.CreateAndSign(*userData)
		if err != nil {
			return nil, err
		}
		return &jwt, nil
	}
	return nil, util.ErrorUserNotFound
}
