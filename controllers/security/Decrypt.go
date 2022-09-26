package security

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func Decrypt(ciphertext []byte, _key string) string {
	key := []byte(_key)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error decrypting run 'cominnek auth login'")
		os.Exit(1)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error decrypting run 'cominnek auth login'")
		os.Exit(1)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println("Error decrypting run 'cominnek auth login'")
		os.Exit(1)
	}

	return string(plaintext)
}
