package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//format code + replace print with fmt + check error handling + make applicable for web + remove panic

func main() {
	textfile := "templates/standard.txt" //access ascii-art letters and store in []string
	file, err := ioutil.ReadFile(textfile)
	if err != nil {
		panic(err)
	}
	letters := strings.Split(string(file), "\n")

	a := os.Args //access arguments
	if len(a) != 2 {
		fmt.Println("ERROR: incorrect number of arguments")
		os.Exit(1)
	}

	text := string(a[1]) //format text
	lines := strings.Split(text, "\n")
	lines = addNewlines(lines)

	for _, line := range lines { //print text as ascii-art
		if line == "\n" {
			fmt.Print("\n")
		} else {
			slice := []rune(line)
			printLine(slice, letters)
		}
	}
}

func printLine(text []rune, letters []string) {
	for i := 0; i < 8; {
		for _, ch := range text {
			if int(ch) > 126 || int(ch) < 32 {
				fmt.Println("ERROR: unallowed characters in text")
				os.Exit(1)
			}

			nr := (int(ch)-32)*9 + 1 + i
			line := letters[nr]
			fmt.Print(line) //???
		}
		if i != 7 {
			fmt.Println("")
		}
		i++
	}
}

func addNewlines(slice []string) []string {
	var newSlice []string
	for i, s := range slice {
		if i == len(slice)-1 {
			newSlice = append(newSlice, s)
		} else if s == "" {
			newSlice = append(newSlice, "\n")
		} else {
			newSlice = append(newSlice, s, "\n")
		}
	}
	return newSlice
}
