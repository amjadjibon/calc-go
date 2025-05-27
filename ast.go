package calc

import "math"

type Node interface {
	Eval() float64
}

type NumberNode struct {
	Value float64
}

func (n *NumberNode) Eval() float64 {
	return n.Value
}

type UnaryOpNode struct {
	Op    TokenType
	Child Node
}

func (n *UnaryOpNode) Eval() float64 {
	switch n.Op {
	case TokenPlus:
		return n.Child.Eval()
	case TokenMinus:
		return -n.Child.Eval()
	default:
		panic("unhandled default case")
	}
}

type BinaryOpNode struct {
	Op    TokenType
	Left  Node
	Right Node
}

func (n *BinaryOpNode) Eval() float64 {
	switch n.Op {
	case TokenPlus:
		return n.Left.Eval() + n.Right.Eval()
	case TokenMinus:
		return n.Left.Eval() - n.Right.Eval()
	case TokenMul:
		return n.Left.Eval() * n.Right.Eval()
	case TokenDiv:
		return n.Left.Eval() / n.Right.Eval()
	case TokenPow:
		return math.Pow(n.Left.Eval(), n.Right.Eval())
	default:
		panic("unhandled default case")
	}
}

type FuncCallNode struct {
	Name string
	Args []Node
}

func (n *FuncCallNode) Eval() float64 {
	switch n.Name {
	case "sin":
		return math.Sin(n.Args[0].Eval())
	case "cos":
		return math.Cos(n.Args[0].Eval())
	case "tan":
		return math.Tan(n.Args[0].Eval())
	case "sinh":
		return math.Sinh(n.Args[0].Eval())
	case "cosh":
		return math.Cosh(n.Args[0].Eval())
	case "tanh":
		return math.Tanh(n.Args[0].Eval())
	case "asin":
		return math.Asin(n.Args[0].Eval())
	case "acos":
		return math.Acos(n.Args[0].Eval())
	case "atan":
		return math.Atan(n.Args[0].Eval())
	case "asinh":
		return math.Asinh(n.Args[0].Eval())
	case "acosh":
		return math.Acosh(n.Args[0].Eval())
	case "atanh":
		return math.Atanh(n.Args[0].Eval())
	case "ln":
		return math.Log(n.Args[0].Eval())
	case "log":
		return math.Log10(n.Args[0].Eval())
	case "sqr", "sqrt":
		return math.Sqrt(n.Args[0].Eval())
	case "abs":
		return math.Abs(n.Args[0].Eval())
	case "ceil":
		return math.Ceil(n.Args[0].Eval())
	case "floor":
		return math.Floor(n.Args[0].Eval())
	case "round":
		return math.Round(n.Args[0].Eval())
	case "exp":
		return math.Exp(n.Args[0].Eval())
	case "max":
		if len(n.Args) != 2 {
			panic("max function requires two arguments")
		}
		return math.Max(n.Args[0].Eval(), n.Args[1].Eval())
	case "min":
		if len(n.Args) != 2 {
			panic("min function requires two arguments")
		}
		return math.Min(n.Args[0].Eval(), n.Args[1].Eval())
	case "log2":
		return math.Log2(n.Args[0].Eval())
	case "log1p":
		return math.Log1p(n.Args[0].Eval())
	case "logb":
		if len(n.Args) != 2 {
			panic("logb function requires two arguments: base and value")
		}
		base := n.Args[0].Eval()
		if base <= 0 || base == 1 {
			panic("logarithm base must be > 0 and != 1")
		}
		return math.Log(n.Args[1].Eval()) / math.Log(base)
	case "gcd":
		if len(n.Args) != 2 {
			panic("gcd function requires two arguments: a and b")
		}
		a := int(n.Args[0].Eval())
		b := int(n.Args[1].Eval())
		return float64(GCD(a, b))
	case "lcm":
		if len(n.Args) != 2 {
			panic("lcm function requires two arguments: a and b")
		}
		a := int(n.Args[0].Eval())
		b := int(n.Args[1].Eval())
		return float64(LCM(a, b))
	case "factorial":
		if len(n.Args) != 1 {
			panic("factorial function requires one argument")
		}
		num := int(n.Args[0].Eval())
		return float64(Factorial(num))
	default:
		panic("unknown function: " + n.Name)
	}
}

type ConstNode struct {
	Name string
}

func (n *ConstNode) Eval() float64 {
	switch n.Name {
	case "pi":
		return math.Pi
	case "e":
		return math.E
	default:
		panic("unknown constant: " + n.Name)
	}
}
