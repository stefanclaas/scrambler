package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var decodeFlag = flag.Bool("d", false, "-d decode input")

func invertSubstitution(substitution map[rune]rune) map[rune]rune {
	inverted := make(map[rune]rune)
	for key, value := range substitution {
		inverted[value] = key
	}
	return inverted
}

func decodeLetters(word string, substitution map[rune]rune) string {
	result := make([]rune, len(word))
	for i, char := range word {
		result[i] = substitution[char]
	}
	return string(result)
}

func scrambleSentence(sentence string, substitution map[rune]rune) string {
	var scrambledWords []string
	scanner := bufio.NewScanner(strings.NewReader(sentence))
	scanner.Split(bufio.ScanRunes)

	var currentWord string

	for scanner.Scan() {
		char := scanner.Text()
		if unicode.IsSpace(rune(char[0])) {
			scrambledWords = append(scrambledWords, decodeLetters(currentWord, substitution))
			scrambledWords = append(scrambledWords, char)
			currentWord = ""
		} else {
			currentWord += char
		}
	}

	scrambledWords = append(scrambledWords, decodeLetters(currentWord, substitution))

	return strings.Join(scrambledWords, "")
}

func main() {
	flag.Parse()

	substitution := map[rune]rune{
		'A': 'D',
		'B': 'F',
		'C': 'E',
		'D': 'A',
		'E': 'B',
		'F': 'C',
		'a': 'd',
		'b': 'f',
		'c': 'e',
		'd': 'a',
		'e': 'b',
		'f': 'c',
		'0': '3',
		'1': '6',
		'2': '5',
		'3': '0',
		'4': '8',
		'5': '7',
		'6': '4',
		'7': '9',
		'8': '1',
		'9': '2',
		'!': 'x',
		'"': '}',
		'#': 'N',
		'$': 'Ü',
		'%': '%',
		'&': 'Ä',
		'\'': 'I',
		'(': '$',
		')': 'U',
		'*': 'Z',
		'+': '[',
		',': 'y',
		'-': 'w',
		'.': 'ä',
		'/': 'J',
		':': '|',
		';': '-',
		'<': '@',
		'=': '€',
		'>': 'Q',
		'?': 'Y',
		'@': 'o',
		'G': ']',
		'H': 'l',
		'I': '.',
		'J': 'V',
		'K': 'ß',
		'L': '(',
		'M': '`',
		'N': 'g',
		'O': 'v',
		'P': 'T',
		'Q': 't',
		'R': 'G',
		'S': '>',
		'T': 'Ö',
		'U': '{',
		'V': 'P',
		'W': '?',
		'X': 'j',
		'Y': '\\',
		'Z': '~',
		'[': '*',
		'\\': '^',
		']': 'i',
		'^': '+',
		'_': 'ẞ',
		'`': '_',
		'g': 'X',
		'h': ')',
		'i': '"',
		'j': 'S',
		'k': 'r',
		'l': 'K',
		'm': 'M',
		'n': 'W',
		'o': '\'',
		'p': 'R',
		'q': 'H',
		'r': ':',
		's': 'q',
		't': 'm',
		'u': '!',
		'v': '=',
		'w': 'k',
		'x': ';',
		'y': 'h',
		'z': 's',
		'{': 'u',
		'|': 'ü',
		'}': 'z',
		'~': 'O',
		'Ä': '<',
		'Ö': ',',
		'Ü': 'L',
		'ß': 'ö',
		'ä': '#',
		'ö': '/',
		'ü': 'n',
		'ẞ': '&',
		'€': 'p',
	}

	if *decodeFlag {
		substitution = invertSubstitution(substitution)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scrambleSentence(scanner.Text(), substitution))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

