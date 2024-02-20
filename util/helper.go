package util

import (
	"fmt"
	"os"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/maulana-jaelani/exercise-job-application-rest-api/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type HelperInf interface {
	HashPassword(pwd string, cost int) ([]byte, error)
	CheckPassword(pwd string, hash []byte) (bool, error)
	CreateAndSign(data entity.User) (string, error)
}
type HelperImpl struct {
}

func (h *HelperImpl) HashPassword(pwd string, cost int) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	if err != nil {

		return nil, err

	}

	return hash, nil

}
func (h *HelperImpl) CheckPassword(pwd string, hash []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))

	if err != nil {

		return false, err

	}

	return true, nil

}

func (h *HelperImpl) CreateAndSign(data entity.User) (string, error) {
	godotenv.Load()
	app := os.Getenv("APP_NAME")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": data.Id,
		"role":data.Role,
		"iss": app,
		"exp": time.Now().Add(1 * time.Hour).Unix(),
		// and other claims, alternatively, you may want to explore on how to create custome claims

	})
	signatur := os.Getenv("JWT_SIGNATURE_KEY")
	signed, err := token.SignedString([]byte(signatur))

	if err != nil {

		return "", err

	}

	return signed, nil

}

func ParseAndVerify(signed string) (jwt.MapClaims, error) {
	app := os.Getenv("APP_NAME")
	signatur := os.Getenv("JWT_SIGNATURE_KEY")
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(signatur), nil

	}, jwt.WithIssuer(app),

		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),

		jwt.WithExpirationRequired(),
	)

	if err != nil {

		return nil, err

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		return claims, nil

	} else {

		return nil, fmt.Errorf("unknown claims")

	}

}
