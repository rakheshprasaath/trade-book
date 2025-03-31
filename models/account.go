package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

type Account struct{
	UserId string `json: "userId"`
	AccountKey string `json: "accountId"`
	AccountNumber string `json: "accountNumber"`
	Password string `json: "password"`
	Server string `json: "server"`
}

func EncryptAccountId(key, plaintext string) (string, error) {
	// Generate a key hash (32 bytes) using SHA-256
	hashedKey := sha256.Sum256([]byte(key))

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return "", err
	}

	// Use AES-GCM (Galois/Counter Mode)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := []byte("uniqueNonce1")  // 12-byte nonce (use a unique nonce in real cases)
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(ciphertext), nil
}
