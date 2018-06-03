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
