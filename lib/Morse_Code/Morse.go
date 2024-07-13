package morsecode

import (
	"fmt"
	"strings"
)

var morsemap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
}

func Encode(words string) string {
	res := ""
	for _, word := range strings.Split(words, "   ") {
		for _, v := range strings.Split(word, " ") {
			for k, val := range morsemap {
				if val == v {
					res += k
				}
			}
		}
		res += "   "
	}
	return strings.TrimSuffix(res, "   ")
}

func Decode(words string) (string, error) {
	if words == "" {
		return "", nil
	}

	res := ""
	for _, word := range strings.Split(words, "   ") {
		if len(word) > 0 {
			for _, c := range word {
				val, ok := morsemap[string(c)]
				if !ok {
					return "", fmt.Errorf("unknown character: %s", string(c))
				}
				res += val
				if c != rune(word[len(word)-1]) {
					res += " "
				}
			}
		}
	}
	return strings.TrimSuffix(res, " "), nil
}
