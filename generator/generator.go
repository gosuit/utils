package generator

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	mrand "math/rand"
	"strings"
)

// GetRandomNum generates a string of random digits of the specified length.
//
// Parameters:
//   - length: The length of the string to be generated.
//
// Returns:
//   - A string of random digits of the specified length.
func GetRandomNum(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = numbs[mrand.Intn(lenOfNumbs)]
	}

	return string(res)
}

// GetSecret generates a string of random characters of the specified length.
//
// Parameters:
//   - length: The length of the string to be generated.
//
// Returns:
//   - A string of random characters of the specified length and an error if one occurred.
func GetSecret(length int) (string, error) {
	b := make([]byte, length)

	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}

	key := hex.EncodeToString(b)

	var sb strings.Builder
	for _, r := range key {
		if strings.ContainsRune(charset, r) {
			sb.WriteRune(r)
		}
	}

	for len(sb.String()) < length {
		sb.WriteRune(symbols[mrand.Intn(lenOfSymbols)])
	}

	return sb.String()[:length], nil
}
