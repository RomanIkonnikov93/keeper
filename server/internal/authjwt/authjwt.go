package authjwt

import (
	"context"

	"github.com/RomanIkonnikov93/keeper/server/internal/models"
	"google.golang.org/grpc/metadata"

	"github.com/golang-jwt/jwt/v4"
)

// EncodeJWT creates a new token including the user ID.
func EncodeJWT(ID, key string) (string, error) {

	var claims = jwt.RegisteredClaims{
		Subject: ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// UserTokenValidation checks the validity of the token and returns the user ID.
func UserTokenValidation(ctx context.Context, key string) (string, error) {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok || len(md.Get("usertoken")) == 0 {
		return "", models.ErrNotExist
	}
	t, err := jwt.Parse(md["usertoken"][0], func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil || !t.Valid {
		return "", models.ErrNotValid
	}

	tkn, err := jwt.ParseWithClaims(md["usertoken"][0], &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tkn.Claims.(*jwt.RegisteredClaims); ok {
		return claims.Subject, nil
	} else {
		return "", err
	}
}
