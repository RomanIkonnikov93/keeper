package authjwt

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestEncodeJWT(t *testing.T) {

	key := "secret"
	ID := ksuid.New().String()

	t.Run("valid", func(t *testing.T) {

		token, err := EncodeJWT(ID, key)

		if err != nil {
			t.Errorf("TestEncodeJWT: %v", err)
		}

		tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil || !tkn.Valid {
			t.Error("TestEncodeJWT: token not valid")
		}
	})

	t.Run("not valid", func(t *testing.T) {

		token, err := EncodeJWT(ID, key)

		if err != nil {
			t.Errorf("TestEncodeJWT: %v", err)
		}

		tkn, err := jwt.Parse(token+"string", func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err == nil || tkn.Valid {
			t.Error("TestEncodeJWT: token must not be valid")
		}
	})
}

func TestUserTokenValidation(t *testing.T) {

	key := "secret"
	ID := ksuid.New().String()
	token, err := EncodeJWT(ID, key)
	if err != nil {
		t.Errorf("TestUserTokenValidation: %v", err)
	}
	md := metadata.New(map[string]string{"usertoken": token})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	t.Run("valid", func(t *testing.T) {

		user, err := UserTokenValidation(ctx, key)
		if err != nil {
			t.Errorf("TestUserTokenValidation: %v", err)
		}

		// check values are equal
		assert.Equal(t, user, ID, "TestUserTokenValidation: different data")
	})

	t.Run("not valid", func(t *testing.T) {

		_, err = UserTokenValidation(ctx, key+"!")
		if err == nil {
			t.Error("TestUserTokenValidation: token must not be valid")
		}
	})
}
