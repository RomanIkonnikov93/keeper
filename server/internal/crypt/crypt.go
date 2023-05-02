package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// generateRandom
func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Encrypt in AES.
func Encrypt(src, key []byte) (string, error) {

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		return "", err
	}

	dst := aesgcm.Seal(nil, nonce, src, nil) // зашифровываем

	encryptedStr := hex.EncodeToString(dst) + "," + hex.EncodeToString(nonce)

	return encryptedStr, nil
}

// CheckPasswordDecrypt
func CheckPasswordDecrypt(incomingPass, encryptedPass string, key []byte) (bool, error) {

	arr := strings.Split(encryptedPass, ",")

	password, _ := hex.DecodeString(arr[0])
	nonce, _ := hex.DecodeString(arr[1])

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return false, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return false, err
	}

	result, err := aesgcm.Open(nil, nonce, password, nil)
	if err != nil {
		return false, err
	}

	if incomingPass == string(result) {
		return true, nil
	} else {
		return false, nil
	}
}
