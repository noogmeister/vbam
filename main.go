package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text := bufio.NewScanner(reader)
	fmt.Printf("hello %s. You are a pig.", text.Text())

}
