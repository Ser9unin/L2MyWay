package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.
	runeStr := []rune(s)
	runeStrLen := len(runeStr) - 1

	var (
		resultString, prevSymbolString string
		n, backslashCount              int
		prevSymbolDigit                bool
	)

	for i, symbol := range runeStr {
		switch {
		case symbol == '\\' && i == runeStrLen:
			return "", ErrInvalidString
		case symbol == '\\':
			backslashCount++
			fallthrough
		case i > 0:
			prevSymbolString = string(runeStr[i-1])
			prevSymbolDigit = unicode.IsDigit(runeStr[i-1])
		}

		if unicode.IsDigit(symbol) {
			n = int(symbol - '0')
			switch {
			case i == 0:
				return "", ErrInvalidString
			case prevSymbolDigit && runeStr[i-2] != '\\':
				return "", ErrInvalidString
			case prevSymbolDigit && backslashCount%2 == 0 && backslashCount > 0:
				return "", ErrInvalidString
			case backslashCount%2 == 0 && backslashCount > 0:
				resultString += strings.Repeat(`\`, n)
			case n == 0:
				resultString = strings.TrimSuffix(resultString, prevSymbolString)
			case backslashCount%2 != 0:
				if backslashCount > 1 {
					resultString += `\`
				}
				resultString += string(symbol)
				backslashCount = 0
			default:
				resultString += strings.Repeat(prevSymbolString, n-1)
			}
		} else if symbol != '\\' {
			resultString += string(symbol)
		}
	}
	return resultString, nil
}
