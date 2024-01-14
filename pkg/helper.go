package pkg

import (
	"strconv"
	"strings"
	"errors"
)

// (xxx, <number>) convert the number from string to int
func Digit(s string) (int, error) {
	ch := false
	sLen := len(s) - 1
	var snum string
	for i, v := range s {
		if v >= '0' && v <= '9' && ch {
			snum += string(v)
		} else if ch && i != sLen {
			return 0, errors.New("Error - Number given to repeat the method can't convert/used")
		} else if v == ' ' {
			ch = true
		} else {
			continue
		}
	}
	num, _ := strconv.Atoi(snum)
	return num, nil
}

// ., ,, !, ?, : and ;
func SpacialSign(r rune) bool {
	s := ".,!?:;"
	for _, v := range s {
		if r == v {
			return true
		}
	}
	return false
}

// vowel (a, e, i, o, u) or h
func VowelChar(r string) bool {
	s := "aeiouh"
	for _, v := range s {
		if r == string(v) {
			return true
		}
	}
	return false
}

// Check if all text is Spacial Sign
func CheckText(s string) bool {
	l := len(s)
	count := 0
	for _, r := range s {
		if SpacialSign(r) {
			count++
		}
	}
	if l == count && l != 0 {
		return true
	}
	return false
}

// function to cnvert the text to the
func GoReloaded(fileText string) string {
	var textList []string
	word := ""
	for i, r := range fileText {
		if r == ')' {
			word += ")"
			textList = append(textList, word)
			word = ""
		} else if strings.Contains(word, "(") {
			word += string(r)
		} else if i == len(fileText)-1 {
			word += string(r)
			textList = append(textList, word)
		} else if !SpacialSign(r) && CheckText(word) {
			textList = append(textList, word)
			word = ""
			if r != ' ' {
				word += string(r)
			}
		} else if !SpacialSign(rune(fileText[i+1])) && SpacialSign(r) {
			word += string(r)
			textList = append(textList, word)
			word = ""
		} else if SpacialSign(rune(fileText[i+1])) && r == ' ' {
			continue
		} else if r == ' ' && word != "" {
			textList = append(textList, word)
			word = ""
		} else if r != ' ' {
			word += string(r)
		}
	}
	var newList []string
	for i, v := range textList {
		switch {
		// (hex) should replace the word before with the decimal version of the word
		case strings.Contains(v, "(hex)"):
			k, err := strconv.ParseInt(textList[i-1], 16, 64)
			if err != nil {
				return "Error - on Hex Number"
			}
			newList[len(newList)-1] = strconv.Itoa(int(k))
		// (bin) should replace the word before with the decimal version of the word
		case strings.Contains(v, "(bin)"):
			k, err := strconv.ParseInt(textList[i-1], 2, 64)
			if err != nil {
				return "Error - on Bin Number"
			}
			newList[len(newList)-1] = strconv.Itoa(int(k))
		// (up) converts the word before with the Uppercase version of it
		case strings.Contains(v, "(up") && strings.Contains(v, ")"):
			if strings.Contains(v, ",") {
				r, err := Digit(v)
				if err != nil {
					return err.Error()
				}
				if len(newList) >= r {
					for j := r; j > 0; j-- {
						newList[len(newList)-j] = strings.ToUpper(textList[i-j])
					}
				} else {
					return "Error - in (up): the given number is out of range"
				}
			} else {
				newList[len(newList)-1] = strings.ToUpper(textList[i-1])
			}
		// (low) converts the word before with the Lowercase version of it
		case strings.Contains(v, "(low") && strings.Contains(v, ")"):
			if strings.Contains(v, ",") {
				r, err := Digit(v)
				if err != nil {
					return err.Error()
				}
				if len(newList) >= r {
					for j := r; j > 0; j-- {
						newList[len(newList)-j] = strings.ToLower(textList[i-j])
					}
				} else {
					return "Error - in (low): the given number is out of range"
				}
			} else {
				newList[len(newList)-1] = strings.ToLower(textList[i-1])
			}
		// (cap) converts the word before with the capitalized version of it
		case strings.Contains(v, "(cap") && strings.Contains(v, ")"):
			if strings.Contains(v, ",") {
				r, err := Digit(v)
				if err != nil {
					return err.Error()
				}
				if len(newList) >= r {
					for j := r; j > 0; j-- {
						newList[len(newList)-j] = strings.Title(textList[i-j])
					}
				} else {
					return "Error - in (cap): the given number is out of range"
				}
			} else {
				newList[len(newList)-1] = strings.Title(textList[i-1])
			}
		// a should be turned into an if the next is vowel (a, e, i, o, u) or h
		case v == "a" || v == "A":
			if len(textList) > i {
				w := textList[i+1]
				var char byte = w[0]
				if string(w[0]) >= "A" && string(w[0]) <= "Z" {
					char = w[0] + 32
				}
				if VowelChar(string(char)) {
					v += "n"
				}
			}
			newList = append(newList, v)
		default:
			newList = append(newList, v)
		}
	}
	var newText string
	pos := true
	for i, v := range newList {
		if v != " " {
			if len(newList)-1 == i {
				newText += v
			} else if CheckText(newList[i+1]) {
				newText += v
				// ' right and left of the word in the middle of them, without any spaces
			} else if v == "'" && pos {
				newText += v
				pos = false
			} else if !pos && newList[i+1] == "'" {
				newText += v
				pos = true
			} else {
				newText += v + " "
			}
		}
	}
	return newText
}