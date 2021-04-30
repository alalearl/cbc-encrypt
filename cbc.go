package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"errors"
)


var base64EncodedKeyBytes = "kV9Ld-X4rKlTQF4ZJwyn9A"
var base64EncodedIvBytes = "PCb_WQYrUgbahQeqDEkuUw"

func Encrypt(text []byte, usekey string) (string, error) {

	if usekey != "" {
		base64EncodedKeyBytes = usekey
	}



	key, err1 := b64.RawURLEncoding.DecodeString(base64EncodedKeyBytes)
	iv, err2 := b64.RawURLEncoding.DecodeString(base64EncodedIvBytes)

	if err1 != nil || err2 != nil {
		panic("Error decoding key/iv")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := text

	b = PKCS5Padding(b, aes.BlockSize, len(text))
	ciphertext := make([]byte, len(b))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, b)

	return b64.RawURLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt( encText string, usekey string) ([]byte, error) {

	if usekey != "" {
		base64EncodedKeyBytes = usekey
	}


	key, err1 := b64.RawURLEncoding.DecodeString(base64EncodedKeyBytes)
	iv, err2 := b64.RawURLEncoding.DecodeString(base64EncodedIvBytes)

	if err1 != nil || err2 != nil {
		panic("Error decoding key/iv")
	}
	
	text, _ := b64.RawURLEncoding.DecodeString(encText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, text)

	return PKCS5UnPadding(decrypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}