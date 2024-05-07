package utils

import "github.com/golang-jwt/jwt/v4"

func GenerateJwtToken(secretKey string, iat int64, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetJwtToken(secretKey string, tokenString string) (jwt.Claims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}
