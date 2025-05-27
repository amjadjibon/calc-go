package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		expr := strings.Join(os.Args[1:], " ")
		fmt.Println(expr)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter expressions or 'q'/'exit' to quit.")
	for {
		fmt.Print("» ")
		if !scanner.Scan() {
			break
		}

		expr := scanner.Text()
		if strings.TrimSpace(expr) == "" {
			continue
		}

		if expr == "q" || expr == "exit" {
			fmt.Println("Exiting...")
			break
		}

		fmt.Println(expr)
	}
}
