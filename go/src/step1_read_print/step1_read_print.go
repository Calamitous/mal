package main

// import "bufio"
import (
	"fmt"
	"reader"
)

func main() {
	fmt.Print("user> ")

	for {
		line := "( 1 2 3)"
		fmt.Println(rep(line))
		fmt.Print("user> ")
	}

	return
}

func rep(input string) string {
	return PRINT(EVAL(READ(input)))
}

func READ(input string) string {
	return reader.Read_str(input)
}

func EVAL(input string) string {
	return input
}

func PRINT(input string) string {
	return input
}
