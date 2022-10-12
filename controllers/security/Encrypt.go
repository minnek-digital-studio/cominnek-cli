package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func Encrypt(base string, _key string) []byte {
	plaintext := []byte(base)
	key := []byte(_key)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext
}
