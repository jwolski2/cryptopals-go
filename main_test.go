package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Set 1, Challenge 1
func TestConvertHexToBase64(t *testing.T) {
	str, err := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	assert.NoError(t, err, "an error occurred converting hex to base64")
	assert.Equal(t, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", str, "base64 value does not equal expected value")
}

// Set 1, Challenge 2
func TestFixedXOR(t *testing.T) {
	str, err := fixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	assert.NoError(t, err, "an error occurred xoring")
	assert.Equal(t, "746865206b696420646f6e277420706c6179", str, "xored value does not equal expected value")
}

// Set 1, Challenge 3
func TestSingleByteXORCipher(t *testing.T) {
	msg, err := singleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	assert.NoError(t, err, "an error occurred decrypting message")
	// Cryptopals does not tell you the expected output. You have to discover
	// that for yourself. But once discovered, you want to assert that it
	// remains true.
	assert.Equal(t, "Cooking MC's like a pound of bacon", msg, "decrypted message does not equal expected value")
}
