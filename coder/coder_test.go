package coder

import (
	"testing"

	"github.com/gosuit/utils/generator"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	coder := getCoder()

	test := "some string"

	enc, err := coder.Encrypt(test)
	if err != nil {
		t.Errorf("Test err: %v", err)
	}

	dec, err := coder.Decrypt(enc)
	if err != nil {
		t.Errorf("Test err: %v", err)
	}

	assert.Equal(t, dec, test)
}

func TestHash(t *testing.T) {
	coder := getCoder()

	test := "some string"

	hash, err := coder.Hash(test)
	if err != nil {
		t.Errorf("Test error: %s", err)
	}

	err = coder.CompareHash(hash, test)
	if err != nil {
		t.Errorf("Test error: %s", err)
	}
}

func getCoder() *Coder {
	secret, err := generator.GetSecret(32)
	if err != nil {
		panic(err)
	}

	return &Coder{
		secret: secret,
		cost:   10,
	}
}
