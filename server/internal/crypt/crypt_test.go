package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPasswordDecrypt(t *testing.T) {

	password := "qwerty123"
	key := []byte("asdkasdandjasdnjafnojfkjokpkjooo")
	encrypted, _ := Encrypt([]byte(password), key)

	t.Run("valid", func(t *testing.T) {

		ok, err := CheckPasswordDecrypt(password, encrypted, key)
		if err != nil {
			t.Errorf("TestCheckPasswordDecrypt: %v", err)
		}
		if !ok {
			t.Error("TestCheckPasswordDecrypt: password not valid")
		}
	})

	t.Run("not valid", func(t *testing.T) {

		ok, err := CheckPasswordDecrypt("asdfghj123", encrypted, key)
		if err != nil {
			t.Errorf("TestCheckPasswordDecrypt: %v", err)
		}
		if ok {
			t.Error("TestCheckPasswordDecrypt: password must be incorrect")
		}
	})
}

func TestDecrypt(t *testing.T) {

	val := []byte("some string")
	str, _ := Encrypt(val, []byte("asdkasdandjasdnjafnojfkjokpkjooo"))
	res, _ := Decrypt(str, []byte("asdkasdandjasdnjafnojfkjokpkjooo"))

	// check values are equal
	assert.Equal(t, val, res, "different data")
}

func Test_generateRandom(t *testing.T) {

	aesblock, _ := aes.NewCipher([]byte("asdkasdandjasdnjafnojfkjokpkjooo"))
	aesgcm, _ := cipher.NewGCM(aesblock)

	res1, _ := generateRandom(aesgcm.NonceSize())
	res2, _ := generateRandom(aesgcm.NonceSize())

	// check values are NOT equal
	assert.NotEqual(t, res1, res2, "the function generates the same values")
}
