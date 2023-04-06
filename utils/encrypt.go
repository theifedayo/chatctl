package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net"
)

func getIPAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range interfaces {
		if (i.Flags&net.FlagUp) != 0 && (i.Flags&net.FlagLoopback) == 0 {
			addrs, err := i.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("unable to determine IP address")
}

func encrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return base64.URLEncoding.EncodeToString(nonce) + "." + base64.URLEncoding.EncodeToString(ciphertext), nil
}

func main() {
	ip, err := getIPAddress()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	encryptedIP, err := encrypt(key, ip)
	if err != nil {
		fmt.Println("Error encrypting IP address:", err)
		return
	}

	fmt.Println("IP address:", ip)
	fmt.Println("Encrypted IP address:", encryptedIP)
}
