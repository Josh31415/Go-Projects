package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

var words = make([]string, 100)
var engWords = make([]string, 100)

func search(charArray []string, bInt, eInt int) {

	for i := bInt; i < eInt; i++ {
		
	}
}

func main() {

	args := os.Args[1:]
	chars := strings.Split(args[0], "");

	file := os.Open("~/english-words/words.txt")
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		s := make([]string, 1)
		engWords = append(s)
		fmt.Println(s)
	}	

	for i := 0; i < len(chars); i++ {
		
	}
	
}
