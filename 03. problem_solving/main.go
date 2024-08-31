package main

import (
	"fmt"
	"unicode"
)

func checkPalindrome (str string) {
	chars := ""

	for _, val := range str {
		if unicode.IsLetter(val) {
			chars += string(unicode.ToLower(val))
		} 
	}

	for i:=0; i<=len(chars)/2; i++ {
		if(chars[i] != chars[len(chars)-1-i]){
			fmt.Println("It Is Not a Palindrome")
		}
	}

	fmt.Println("It Is a Palindrome")
}

func main () {
	checkPalindrome("?????KK;;;;madam,,,..Kk");
}