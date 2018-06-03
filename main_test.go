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

// Set 1, Challenge 4
func TestDetectSingleCharacterXOR(t *testing.T) {
	msg, err := detectSingleCharacterXOR("./4.txt")
	assert.NoError(t, err, "an error occurred detecting key")
	// Again, Cryptopals does not you the expected output. But it's pretty clear
	// what the correct answer is once you decrypt all lines in the file.
	assert.Equal(t, "Now that the party is jumping\n", msg, "detected msesage does not equal expected value")
}

// Set 1, Challenge 5
func TestRepeatingKeyXOR(t *testing.T) {
	encrypted, err := encryptUsingRepeatingKeyXOR(
		`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`, "ICE")
	assert.NoError(t, err, "an error occurring during encryption")
	assert.Equal(t, "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f", encrypted, "encrypted message does not equal expected value")
}

// Set 1, Challenge 6
func TestBreakingRepeatingKeyXOR(t *testing.T) {
	distance, err := computeHammingDistance("this is a test", "wokka wokka!!!")
	assert.NoError(t, err, "an error occurred computing hamming distance")
	assert.Equal(t, 37, distance, "distance is not equal to expected value")
}
