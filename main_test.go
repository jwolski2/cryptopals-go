package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHexToBase64(t *testing.T) {
	str, err := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	assert.NoError(t, err, "an error occurred converting hex to base64")
	assert.Equal(t, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", str, "base64 value is not equal to expected value")
}
