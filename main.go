package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"testing"
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

func hextobase64(hexString string) (string, error) {
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

func TestConvertHexToBase64(t *testing.T) {

}

func exer1() {
	b64, err := hextobase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		fmt.Println("failed to base64 encode hex string", err.Error())
		os.Exit(1)
	}

	fmt.Println("=== Set 1, Challenge 1: Convert hex to base64")
	fmt.Println(b64)
	fmt.Println()
}

func exer2() {
	str, err := fixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		fmt.Println("failed to XOR strings", err.Error())
		os.Exit(1)
	}

	fmt.Println("=== Set 1, Challenge 2: Fixed XOR")
	fmt.Println(str)
	fmt.Println()
}

func exer3() {
	textBytes, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Println("failed to decode string")
		os.Exit(1)
	}

	score := scoreText(textBytes)
	fmt.Println("=== Set 1, Challenge 3: Single-byte XOR cipher")
	fmt.Println(score)
	fmt.Println()
}

func exer4() {
	file, err := os.Open("./4.txt")
	if err != nil {
		return
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
			fmt.Println("failed to decode string")
			os.Exit(1)
		}

		key := findKey(textBytes)
		message := decryptMessage(byte(key), textBytes)
		score := computeScore(textBytes)
		if score > maxScore {
			maxScore = score
			maxLine = message
		}
	}
	fmt.Println("=== Set 1, Challenge 4: Detect single-character XOR")
	fmt.Println(maxLine)
	fmt.Println()
}

func main() {
	exer1()
	exer2()
	exer3()
	exer4()
}
