package main
// import "bufio"
import (
        "fmt"
        "bufio"
        "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("user> ")

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(rep(line));
    fmt.Print("user> ")
  }

  return;
}

func rep(input string) string {
  return PRINT(EVAL(READ(input)));
}

func READ(input string) string {
  return input;
}

func EVAL(input string) string {
  return input;
}

func PRINT(input string) string {
  return input;
}
