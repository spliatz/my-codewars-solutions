package kata

import "strings"

const (
	dash          = "111111"
	dot           = "11"
	badDot        = "111"
	badDot2       = "1111"
	badDot3       = "11111"
	pauseInLetter = "00"
	pauseInWord   = "000000"
	pauseInText   = "00000000000000"
)

// You may get original char by morse code like this: MORSE_CODE[char]

func cutNulls(s string) string {
	if s[0] != '0' {
		return s
	}

	nulls := 0

	for _, elem := range s {
		if elem != '0' {
			break
		}
		nulls++
	}

	trims := strings.Repeat("0", nulls)
	return strings.Trim(s, trims)
}

func DecodeBits(bits string) string {
	bits = cutNulls(bits)
	words := strings.Split(bits, pauseInText)
	str := ""

	for _, word := range words {
		letters := strings.Split(word, pauseInWord)
		symbolWord := ""

		for _, letter := range letters {
			symbols := strings.Split(letter, pauseInLetter)
			// sym := ""
			for _, symbol := range symbols {
				if string(symbol) == dash {
					symbolWord += "-"
					continue
				}
				if string(symbol) == dot || string(symbol) == badDot || string(symbol) == badDot2 || string(symbol) == badDot3 {
					symbolWord += "."
					continue
				}
			}
			symbolWord += " "
		}
		str += symbolWord + "   "
	}

	return strings.TrimSpace(str)
}

func DecodeMorse(morseCode string) string {

	str := ""

	word := strings.Split(morseCode, "   ")
	for _, w := range word {
		letters := strings.Split(w, " ")
		for _, letter := range letters {
			human, ok := MORSE_CODE[letter]
			if !ok {
				str += ""
			} else {
				str += human
			}
		}
		str += " "
	}

	return strings.ToUpper(strings.TrimSpace(str))
}
