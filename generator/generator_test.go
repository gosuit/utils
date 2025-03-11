package generator

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomNum(t *testing.T) {
	length := 6

	num := GetRandomNum(length)

	assert.Equal(t, len(num), length)

	for _, s := range num {
		if !slices.Contains(numbs, s) {
			t.Errorf("Num must include only numbers.")
		}
	}
}

func TestGetSecret(t *testing.T) {
	length := 10

	secret, err := GetSecret(length)
	if err != nil {
		t.Errorf("Test error: %v", err)
	}

	assert.Equal(t, len(secret), length)

	for _, s := range secret {
		if !slices.Contains(symbols, s) {
			t.Errorf("Secret must include only symbols from expected charset.")
		}
	}
}

func TestSymbols(t *testing.T) {
	expected := []rune(charset)

	assert.Equal(t, expected, symbols)
}

func TestNumbs(t *testing.T) {
	expected := []rune("1234567890")

	assert.Equal(t, expected, numbs)
}

func TestCharset(t *testing.T) {
	expected := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	assert.Equal(t, expected, charset)
}

func TestLenOfSymbols(t *testing.T) {
	assert.Equal(t, lenOfSymbols, 62)
}

func TestLenOfNumbs(t *testing.T) {
	assert.Equal(t, lenOfNumbs, 10)
}
