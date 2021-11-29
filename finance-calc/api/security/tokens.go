package security

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"financeCalc/api/models"
	"financeCalc/api/utils"
)

func getKey() []byte {
	if key, ok := os.LookupEnv("JWT_KEY"); ok {
		return []byte(key)
	}
	panic(errors.New("JWT_KEY is not set as an environment variable"))
}

func GenerateUserToken(user *models.User) string {
	now := time.Now()
	expiration := now.Add(time.Hour * 4)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usr": user,
		"exp": expiration.Unix(),
		"iat": now.Unix(),
		"iss": "finance-calc",
		"sub": user.Email,
	})
	signedToken, err := token.SignedString(getKey())
	utils.CheckError(err)
	return signedToken
}
