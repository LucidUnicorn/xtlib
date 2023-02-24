package random

import (
	"crypto/rand"
	"math/big"
)

// String generates a cryptographically secure random string.
type String struct {
	// CharSet is the set of characters used to generate the string.
	CharSet []rune
}

// DefaultStringCharSet is used to generate the string unless overridden in a
// custom instance of String.
var DefaultStringCharSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// NewRandomString creates a new String with default options for quick
// generation of random strings.
func NewRandomString() String {
	return String{
		CharSet: DefaultStringCharSet,
	}
}

// Generate creates the random string according to the configured options.
func (r *String) Generate(length int) string {
	str := make([]rune, length)
	charSetLen64 := int64(len(r.CharSet))

	for i := range str {
		randInt, _ := rand.Int(rand.Reader, big.NewInt(charSetLen64))
		str[i] = r.CharSet[randInt.Int64()]
	}

	return string(str)
}
