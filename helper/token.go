package helper

import (
	"go_confess_space-project/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateToken(claims jwt.MapClaims, secret string, duration time.Duration) string {
	now := time.Now()
	claims["exp"] = now.Add(duration).Unix()
	claims["iat"] = now.Unix()

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	return token
}

func CreateAccessToken(user dto.UserResponse) string {
	secret := os.Getenv("GENERATE_TOKEN_SECRET")
	claims := jwt.MapClaims{
		"id":          user.Id,
		"name":        user.Name,
		"username":    user.Username,
		"email":       user.Email,
		"avatar_type": user.AvatarType,
	}
	return generateToken(claims, secret, time.Minute*30)
}

func CreateRefreshToken(user dto.UserResponse) string {
	secret := os.Getenv("GENERATE_REFRESH_TOKEN_SECRET")
	claims := jwt.MapClaims{
		"id":          user.Id,
		"name":        user.Name,
		"username":    user.Username,
		"email":       user.Email,
		"avatar_type": user.AvatarType,
	}
	return generateToken(claims, secret, time.Hour*24*30)
}
