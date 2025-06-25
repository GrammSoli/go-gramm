package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encryptor struct {
	Key string
}

func NewEncrypter() *Encryptor {
	key := os.Getenv("ENCRYPTION_KEY")
	if key == "" {
		panic("ENCRYPTION_KEY не передан")
	}
	return &Encryptor{
		Key: key,
	}
}

func (e *Encryptor) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic("Ошибка создания шифра: " + err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic("Ошибка создания GCM: " + err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic("Ошибка генерации nonce: " + err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

func (e *Encryptor) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		panic("Ошибка создания шифра: " + err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic("Ошибка создания GCM: " + err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := encryptedStr[:nonceSize], encryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic("Ошибка расшифровки: " + err.Error())
	}
	return plainText
}
strings.NewReplacer("http://", "", "https://", "").Replace(url)