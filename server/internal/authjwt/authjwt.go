package authjwt

import (
	"github.com/RomanIkonnikov93/keeper/server/internal/config"

	"github.com/golang-jwt/jwt/v4"
)

func EncodeJWT(ID, key string) (string, error) {

	var claims = jwt.RegisteredClaims{
		ID: ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseJWTWithClaims(token string, cfg config.Config) (string, error) {

	tkn, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.JWTSecretKey, nil
	})
	if claims, ok := tkn.Claims.(*jwt.RegisteredClaims); ok {
		return claims.ID, nil
	} else {
		return "", err
	}
}
