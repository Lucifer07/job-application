package payload

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
)
func LoginToUser(userLogin dto.Login)entity.User{
	var user entity.User
	user.Email=userLogin.Email
	user.Password=userLogin.Password
	return user
}