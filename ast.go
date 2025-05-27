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
