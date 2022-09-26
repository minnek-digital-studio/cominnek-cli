package security

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)


func generateKey() string {
	host, _ := os.Hostname()
	salt := fmt.Sprintf("%s%s", host, time.Now().Format("2006-01-02 15:04:05"))
	base64Salt := base64.StdEncoding.EncodeToString([]byte(salt))
	encrypted := Encrypt(base64Salt, config.Private.EncryptKey)
	files.Create(encrypted, config.Public.KeyPath)
	return base64Salt
}

func GetKey() string {
	currentKey := "";
	encryptedKey := files.Read(config.Public.KeyPath)

	if(encryptedKey != nil && len(string(encryptedKey)) > 0) {
		currentKey = Decrypt(encryptedKey, config.Private.EncryptKey)
	}

	if currentKey == "" {
		fmt.Println("No key found, generating new key")
		currentKey = generateKey()
	}

	if currentKey == "" {
		fmt.Println("Error generating key")
		os.Exit(1)
	}

	return currentKey
}
