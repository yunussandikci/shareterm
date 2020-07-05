package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type EncryptionService interface {
	Encrypt(plain []byte, key []byte) ([]byte, error)
	Decrypt(cipherText []byte, key []byte) ([]byte, error)
}

type encryptionService struct {
}

func NewEncryptionService() EncryptionService {
	return encryptionService{}
}

func (e encryptionService) Encrypt(plain []byte, key []byte) ([]byte, error) {
	aesCipher, aesCipherErr := aes.NewCipher(key)
	if aesCipherErr != nil {
		return nil, aesCipherErr
	}

	gcm, gcmErr := cipher.NewGCM(aesCipher)
	if gcmErr != nil {
		return nil, gcmErr
	}

	nonce := make([]byte, gcm.NonceSize())
	_, nonceErr := io.ReadFull(rand.Reader, nonce)
	if nonceErr != nil {
		return nil, nonceErr
	}

	return gcm.Seal(nonce, nonce, plain, nil), nil
}

func (e encryptionService) Decrypt(cipherText []byte, key []byte) ([]byte, error) {
	aesCipher, aesCipherErr := aes.NewCipher(key)
	if aesCipherErr != nil {
		return nil, aesCipherErr
	}

	gcm, gcmErr := cipher.NewGCM(aesCipher)
	if gcmErr != nil {
		return nil, gcmErr
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("err short cipher text")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}
