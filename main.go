package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
)

var scores = map[string]float64{
	"a": 0.0651738,
	"b": 0.0124248,
	"c": 0.0217339,
	"d": 0.0349835,
	"e": 0.1041442,
	"f": 0.0197881,
	"g": 0.0158610,
	"h": 0.0492888,
	"i": 0.0558094,
	"j": 0.0009033,
	"k": 0.0050529,
	"l": 0.0331490,
	"m": 0.0202124,
	"n": 0.0564513,
	"o": 0.0596302,
	"p": 0.0137645,
	"q": 0.0008606,
	"r": 0.0497563,
	"s": 0.0515760,
	"t": 0.0729357,
	"u": 0.0225134,
	"v": 0.0082903,
	"w": 0.0171272,
	"x": 0.0013692,
	"y": 0.0145984,
	"z": 0.0007836,
	" ": 0.1918182,
}

func fixedXOR(a, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("failed to XOR: buffers are not of equal length")
	}

	aBytes, err := hex.DecodeString(a)
	if err != nil {
		return "", err
	}

	bBytes, err := hex.DecodeString(b)
	if err != nil {
		return "", err
	}

	dest := make([]byte, len(aBytes), len(aBytes))
	for i, e := range aBytes {
		dest[i] = e ^ bBytes[i]
	}

	return hex.EncodeToString(dest), nil
}

func hexToBase64(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func computeScore(textBytes []byte) float64 {
	charScores := make([]float64, 256)
	for i := 0; i < 256; i++ {
		for _, b := range textBytes {
			xored := b ^ byte(i)
			if val, ok := scores[string(xored)]; ok {
				charScores[i] += val
			}
		}
	}

	maxScore := float64(-1)
	for _, s := range charScores {
		if maxScore < s {
			maxScore = s
		}
	}

	return maxScore
}

func findKey(textBytes []byte) int {
	charScores := make([]float64, 256)
	for i := 0; i < 256; i++ {
		for _, b := range textBytes {
			xored := b ^ byte(i)
			if val, ok := scores[string(xored)]; ok {
				charScores[i] += val
			}
		}
	}

	maxScore := float64(-1)
	theIndex := -1
	for i, s := range charScores {
		if maxScore < s {
			maxScore = s
			theIndex = i
		}
	}

	return theIndex
}

func decryptMessage(key byte, textBytes []byte) string {
	message := make([]byte, len(textBytes))
	for i, b := range textBytes {
		message[i] = b ^ key
	}

	return string(message)
}

func scoreText(textBytes []byte) string {
	key := findKey(textBytes)
	return decryptMessage(byte(key), textBytes)
}

func singleByteXORCipher(str string) (string, error) {
	textBytes, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}

	return scoreText(textBytes), nil
}

func detectSingleCharacterXOR(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	maxScore := float64(-1)
	maxLine := ""
	for _, line := range lines {
		textBytes, err := hex.DecodeString(line)
		if err != nil {
			return "", err
		}

		key := findKey(textBytes)
		message := decryptMessage(byte(key), textBytes)
		score := computeScore(textBytes)
		if score > maxScore {
			maxScore = score
			maxLine = message
		}
	}
	return maxLine, nil
}
