package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHexToBase64(t *testing.T) {
	_, err := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	assert.NoError(t, err, "an error occurred converting hex to base64")
}
