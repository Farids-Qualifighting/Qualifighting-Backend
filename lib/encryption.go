package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"reflect"
)

const key = "eThWmZq4t7w!z%C*F-J@NcRfUjXn2r5u"

func EncryptByteArray(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func DecryptByteArray(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func encryptString(plainText string, key string) (string, error) {
	encryptedByteArray, err := EncryptByteArray([]byte(plainText), []byte(key))

	if err != nil {
		return "", err
	}

	return string(encryptedByteArray), nil
}

func DecryptString(cipherText string, key string) (string, error) {
	decryptedByteArray, err := DecryptByteArray([]byte(cipherText), []byte(key))

	if err != nil {
		return "", err
	}

	return string(decryptedByteArray), nil
}

func Encrypt[T any](obj T) (T, error) {
	v := reflect.ValueOf(&obj).Elem()
	if v.Kind() != reflect.Struct {
		return obj, fmt.Errorf("argument must be a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() == reflect.String {
			if v.Type().Field(i).Tag.Get("encryption") == "true" {
				plainText := f.String()
				cipherText, errEncryption := encryptString(plainText, key)
				if errEncryption != nil {
					return obj, errEncryption
				}
				f.SetString(cipherText)
			}
		}
	}

	return obj, nil
}
