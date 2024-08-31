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

func fibonacci (num int) []int {
	var fibos = []int{}

	prev1, prev2 := 0, 1;

	fib := func () int {
		prev1, prev2 = prev2, prev1 + prev2

		return prev1;
	}

	for i:=0; i<num; i++ {
		fibos = append(fibos, fib());
	}

	fmt.Printf("The First %d Fibonaci are: %v \n", num, fibos)
	return fibos
}

func main () {
	checkPalindrome("?????KK;;;;madam,,,..Kk");
	fibonacci(20);
}