package coder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Secret   string `env:"CODER_ENCRYPT_SECRET" env-default:""`
	HashCost int    `env:"CODER_HASH_COST" env-default:"10"`
}

type Coder struct {
	secret string
	cost   int
}

func New(cfg *Config) (*Coder, error) {
	if len(cfg.Secret) != 16 && len(cfg.Secret) != 24 && len(cfg.Secret) != 32 {
		return nil, errors.New("secret must be 16, 24 or 32 bytes")
	}

	if cfg.HashCost < bcrypt.MinCost || cfg.HashCost > bcrypt.MaxCost {
		return nil, errors.New("hash cost must be in range [4; 31]")
	}

	return &Coder{
		secret: cfg.Secret,
		cost:   cfg.HashCost,
	}, nil
}

func (c *Coder) Encrypt(text string) (string, error) {
	aes, err := aes.NewCipher([]byte(c.secret))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)

	return hex.EncodeToString(ciphertext), nil
}

func (c *Coder) Decrypt(text string) (string, error) {
	b, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	text = string(b)

	aes, err := aes.NewCipher([]byte(c.secret))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := text[:nonceSize], text[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (c *Coder) Hash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), c.cost)
	if err != nil {
		return "", errors.New("invalid text. text length must be less or equal to 72")
	}

	return string(hash), nil
}

func (c *Coder) CompareHash(hash, text string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return errors.New("invalid hash")
	}

	return nil
}
