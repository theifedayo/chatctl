package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

func decrypt(key []byte, ciphertext string) (string, error) {
	parts := strings.Split(ciphertext, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid ciphertext format")
	}

	nonce, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, []byte(parts[1]), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}


