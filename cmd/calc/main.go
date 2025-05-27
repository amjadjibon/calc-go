package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amjadjibon/calc-go"
)

func main() {
	if len(os.Args) > 1 {
		expr := strings.Join(os.Args[1:], " ")
		root := calc.NewParser(expr).Parse()
		fmt.Println(root.Eval())
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

		root := calc.NewParser(expr).Parse()
		fmt.Println(root.Eval())
	}
}
