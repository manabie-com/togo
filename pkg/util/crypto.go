package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Encrypt(text string) (string, string) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString("bb64719a7a0f14cdcceda03541bfbf81054d7360f37a149900665a67d2b89f36")
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	// fmt.Printf("%x\n", nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	// fmt.Printf("%x\n", ciphertext)
	return hex.EncodeToString(nonce[:]), hex.EncodeToString(ciphertext[:])
}

func Decrypt(hash, nonceInput string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString("bb64719a7a0f14cdcceda03541bfbf81054d7360f37a149900665a67d2b89f36")
	ciphertext, _ := hex.DecodeString(hash)
	nonce, _ := hex.DecodeString(nonceInput)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("%s\n", plaintext)
	return string(plaintext)
}

func DecryptUseKey(message string, privateKey *rsa.PrivateKey) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(message)
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		//log.Error(err)
	}
	return string(plaintext)
}