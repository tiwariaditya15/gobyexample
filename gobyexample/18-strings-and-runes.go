package main

import (
	"fmt"
	"unicode/utf8"
)

func conditionals(r rune) {
	// 't' rune literal here & literals are supposed to be written using single quotes
	if r == 't' {
		fmt.Println("Found t")
	} else if r == 'ด' {
		fmt.Println("Found rune ดี in ")
	}
}

func main() {
	// greet := "hello"
	thaiGreet := "สวัสดี"
	fmt.Println(utf8.RuneCountInString(thaiGreet))

	for _, char := range thaiGreet {
		fmt.Printf("%#U\n", char)
	}
	
	for i, w := 0, 0;i < len(thaiGreet); i = i + w {
		// access the first utf8 code pass it to this utility and it returns rune itself with width which can be used to hop to next rune instead of hopping to next utf8 encoded code
		runeValue, width := utf8.DecodeRuneInString(thaiGreet[i:])
		fmt.Printf("%#U at %d\n", runeValue, i)		
		w = width
		conditionals(runeValue)
	}
}