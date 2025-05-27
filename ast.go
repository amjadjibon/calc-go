package calc

import (
	"fmt"
	"math"
)

type Node interface {
	Eval() any
}

type NumberNode struct {
	Value float64
}

func (n *NumberNode) Eval() any {
	return n.Value
}

type UnaryOpNode struct {
	Op    TokenType
	Child Node
}

func (n *UnaryOpNode) Eval() any {
	child := n.Child.Eval()
	if v, ok := child.(float64); ok {
		switch n.Op {
		case TokenPlus:
			return v
		case TokenMinus:
			return -v
		default:
			panic("unhandled default case")
		}
	}
	panic("unhandled default case")
}

type BinaryOpNode struct {
	Op    TokenType
	Left  Node
	Right Node
}

func (n *BinaryOpNode) Eval() any {
	left := n.Left.Eval()
	right := n.Right.Eval()
	lf, lok := left.(float64)
	rf, rok := right.(float64)
	if lok && rok {
		switch n.Op {
		case TokenPlus:
			return lf + rf
		case TokenMinus:
			return lf - rf
		case TokenMul:
			return lf * rf
		case TokenDiv:
			return lf / rf
		case TokenPow:
			return math.Pow(lf, rf)
		case TokenMod:
			return math.Mod(lf, rf)
		default:
			panic("unhandled default case")
		}
	}
	panic("unhandled default case")
}

type FuncCallNode struct {
	Name string
	Args []Node
}

func (n *FuncCallNode) Eval() any {
	switch n.Name {
	case "sin":
		return math.Sin(asFloat(n.Args[0].Eval()))
	case "cos":
		return math.Cos(asFloat(n.Args[0].Eval()))
	case "tan":
		return math.Tan(asFloat(n.Args[0].Eval()))
	case "sinh":
		return math.Sinh(asFloat(n.Args[0].Eval()))
	case "cosh":
		return math.Cosh(asFloat(n.Args[0].Eval()))
	case "tanh":
		return math.Tanh(asFloat(n.Args[0].Eval()))
	case "asin":
		return math.Asin(asFloat(n.Args[0].Eval()))
	case "acos":
		return math.Acos(asFloat(n.Args[0].Eval()))
	case "atan":
		return math.Atan(asFloat(n.Args[0].Eval()))
	case "asinh":
		return math.Asinh(asFloat(n.Args[0].Eval()))
	case "acosh":
		return math.Acosh(asFloat(n.Args[0].Eval()))
	case "atanh":
		return math.Atanh(asFloat(n.Args[0].Eval()))
	case "ln":
		return math.Log(asFloat(n.Args[0].Eval()))
	case "log":
		return math.Log10(asFloat(n.Args[0].Eval()))
	case "sqr", "sqrt":
		return math.Sqrt(asFloat(n.Args[0].Eval()))
	case "abs":
		return math.Abs(asFloat(n.Args[0].Eval()))
	case "ceil":
		return math.Ceil(asFloat(n.Args[0].Eval()))
	case "floor":
		return math.Floor(asFloat(n.Args[0].Eval()))
	case "round":
		return math.Round(asFloat(n.Args[0].Eval()))
	case "exp":
		return math.Exp(asFloat(n.Args[0].Eval()))
	case "max":
		if len(n.Args) != 2 {
			panic("max function requires two arguments")
		}
		return math.Max(asFloat(n.Args[0].Eval()), asFloat(n.Args[1].Eval()))
	case "min":
		if len(n.Args) != 2 {
			panic("min function requires two arguments")
		}
		return math.Min(asFloat(n.Args[0].Eval()), asFloat(n.Args[1].Eval()))
	case "log2":
		return math.Log2(asFloat(n.Args[0].Eval()))
	case "log1p":
		return math.Log1p(asFloat(n.Args[0].Eval()))
	case "logb":
		if len(n.Args) != 2 {
			panic("logb function requires two arguments: base and value")
		}
		base := asFloat(n.Args[0].Eval())
		if base <= 0 || base == 1 {
			panic("logarithm base must be > 0 and != 1")
		}
		return math.Log(asFloat(n.Args[1].Eval())) / math.Log(base)
	case "gcd":
		if len(n.Args) != 2 {
			panic("gcd function requires two arguments: a and b")
		}
		a := int(asFloat(n.Args[0].Eval()))
		b := int(asFloat(n.Args[1].Eval()))
		return float64(GCD(a, b))
	case "lcm":
		if len(n.Args) != 2 {
			panic("lcm function requires two arguments: a and b")
		}
		a := int(asFloat(n.Args[0].Eval()))
		b := int(asFloat(n.Args[1].Eval()))
		return float64(LCM(a, b))
	case "factorial":
		if len(n.Args) != 1 {
			panic("factorial function requires one argument")
		}
		num := int(asFloat(n.Args[0].Eval()))
		return float64(Factorial(num))
	case "bin":
		if len(n.Args) != 1 {
			panic("bin function requires one argument")
		}
		return fmt.Sprintf("0b%b", int(asFloat(n.Args[0].Eval())))
	case "oct":
		if len(n.Args) != 1 {
			panic("oct function requires one argument")
		}
		return fmt.Sprintf("0o%o", int(asFloat(n.Args[0].Eval())))
	case "hex":
		if len(n.Args) != 1 {
			panic("hex function requires one argument")
		}
		return fmt.Sprintf("0x%X", int(asFloat(n.Args[0].Eval())))
	default:
		panic("unknown function: " + n.Name)
	}
}

type ConstNode struct {
	Name string
}

func (n *ConstNode) Eval() any {
	switch n.Name {
	case "pi":
		return math.Pi
	case "e":
		return math.E
	default:
		panic("unknown constant: " + n.Name)
	}
}

func asFloat(v any) float64 {
	if f, ok := v.(float64); ok {
		return f
	}
	panic("value is not a float64")
}
