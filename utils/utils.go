package utils

import (
	"crypto/rand"
	"fmt"
)

func IPAndEncrypt() {
	ip, err := getIPAddress()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	//key := make([]byte, 32)
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

func Decrypt(theencrypted string) {
	//key := make([]byte, 32)
	encryptedIP := theencrypted

	ip, err := decrypt(key, encryptedIP)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error decrypting IP address:", err)
		return
	}

	fmt.Println("Decrypted IP address:", ip)
}
