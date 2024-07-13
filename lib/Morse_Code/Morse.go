package morsecode

import (
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

func Encode(words string) (string) {
	res := ""
	for _, words := range strings.Split(words, "   ") {
		for  _, v := range strings.Split(words, " ") {
			for  k, c := range morsemap {
				if c == v { 
					res += string(k)
				}
			}
		}
	}
	return res
}

func Decode(words string) (string) {
	res := ""
	for _, v := range strings.Split(words, " ") {
		for k, c := range v {
			if val, ok := morsemap[string(c)]; ok {
				res += val
				if k < len(v) - 1 {
					res += " "
				}
			}
		}
	}
	return res
}

