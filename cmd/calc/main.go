package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amjadjibon/calc-go"
)

func formatResult(result float64) string {
	str := fmt.Sprintf("%.6f", result)
	str = strings.TrimRight(str, "0")
	str = strings.TrimSuffix(str, ".")
	if str == "" {
		return "0"
	}
	return str
}

func safeEvalPrint(expr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
		}
	}()
	root := calc.NewParser(strings.TrimSpace(expr)).Parse()
	result := root.Eval()
	switch v := result.(type) {
	case string:
		fmt.Println(v)
	case float64:
		fmt.Println(formatResult(v))
	default:
		fmt.Println(v)
	}
}

func main() {
	if len(os.Args) > 1 {
		expr := strings.Join(os.Args[1:], " ")
		if expr == "version" || expr == "-v" || expr == "--version" {
			fmt.Printf("calc-go version %s\n", calc.Version)
			return
		}
		safeEvalPrint(expr)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter expressions or 'q'/'exit' to quit.")
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
			fmt.Println("sayonara...")
			break
		}

		safeEvalPrint(expr)
	}
}
