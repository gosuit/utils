package coder

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Coder provides methods for encryption, decryption, and hashing.
type Coder interface {
	// Encrypt encrypts the given plaintext using AES encryption.
	//
	// Parameters:
	//   - text: The plaintext string to encrypt.
	//
	// Returns:
	//   - The encrypted text as a hexadecimal string or an error if the encryption fails.
	Encrypt(text string) (string, error)

	// Decrypt decrypts the given hexadecimal string back into plaintext using AES decryption.
	//
	// Parameters:
	//   - text: The encrypted text as a hexadecimal string to decrypt.
	//
	// Returns:
	//   - The decrypted plaintext string or an error if the decryption fails.
	Decrypt(text string) (string, error)

	// Hash generates a bcrypt hash of the given plaintext string.
	//
	// Parameters:
	//   - text: The plaintext string to hash.
	//
	// Returns:
	//   - The hashed string or an error if the hashing fails.
	Hash(text string) (string, error)

	// CompareHash compares a bcrypt hash with a plaintext string to verify if they match.
	//
	// Parameters:
	//   - hash: The bcrypt hash to compare against.
	//   - text: The plaintext string to compare with the hash.
	//
	// Returns:
	//   - An error if the hash does not match the plaintext.
	CompareHash(hash, text string) error
}

// Config holds the configuration for the Coder.
// It includes the encryption secret and the cost for hashing.
type Config struct {
	// Encryption secret must be 16, 24 or 32 bytes
	Secret string `env:"CODER_ENCRYPT_SECRET" env-default:""`

	// HashCost must be in range [4; 31]
	HashCost int `env:"CODER_HASH_COST" env-default:"10"`
}

// New creates a new Coder instance with the provided configuration.
// It validates the secret length and hash cost before returning the instance.
//
// Parameters:
//   - cfg: A pointer to a Config struct containing the necessary settings.
//
// Returns:
//   - A pointer to a Coder instance or an error if the configuration is invalid.
func New(cfg *Config) (Coder, error) {
	if len(cfg.Secret) != 16 && len(cfg.Secret) != 24 && len(cfg.Secret) != 32 {
		return nil, errors.New("secret must be 16, 24 or 32 bytes")
	}

	if cfg.HashCost < bcrypt.MinCost || cfg.HashCost > bcrypt.MaxCost {
		return nil, errors.New("hash cost must be in range [4; 31]")
	}

	return &coder{
		secret: cfg.Secret,
		cost:   cfg.HashCost,
	}, nil
}

type coder struct {
	secret string
	cost   int
}

func (c *coder) Encrypt(text string) (string, error) {
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

func (c *coder) Decrypt(text string) (string, error) {
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

func (c *coder) Hash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), c.cost)
	if err != nil {
		return "", errors.New("invalid text. text length must be less or equal to 72")
	}

	return string(hash), nil
}

func (c *coder) CompareHash(hash, text string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return errors.New("invalid hash")
	}

	return nil
}
