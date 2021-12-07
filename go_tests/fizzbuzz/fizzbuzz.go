package fizzbuzz

import "strconv"

func Fizzbuzz(numero int) string {
	var text string
	if numero%3 == 0 {
		text += "fizz"
	}
	if numero%5 == 0 {
		text += "buzz"
	}
	if text == "" {
		text = strconv.FormatInt(int64(numero), 10)
	}
	return text
}
