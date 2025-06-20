package calc

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lex *Lexer
	cur Token
}

func NewParser(input string) *Parser {
	l := NewLexer(input)
	p := &Parser{lex: l}
	p.cur = l.NextToken()
	return p
}

func (p *Parser) next(t TokenType) {
	if p.cur.Type == t {
		p.cur = p.lex.NextToken()
	} else {
		panic(fmt.Sprintf("Expected %v, got %v", t, p.cur.Type))
	}
}

func (p *Parser) Parse() Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() Node {
	n := p.parseTerm()
	for p.cur.Type == TokenPlus || p.cur.Type == TokenMinus {
		op := p.cur.Type
		p.next(op)
		n = &BinaryOpNode{op, n, p.parseTerm()}
	}
	return n
}

func (p *Parser) parseTerm() Node {
	n := p.parseFactor()
	for p.cur.Type == TokenMul || p.cur.Type == TokenDiv || p.cur.Type == TokenMod {
		op := p.cur.Type
		p.next(op)
		n = &BinaryOpNode{op, n, p.parseFactor()}
	}
	return n
}

func (p *Parser) parseFactor() Node {
	n := p.parseUnary()
	for p.cur.Type == TokenPow {
		op := p.cur.Type
		p.next(TokenPow)
		n = &BinaryOpNode{op, n, p.parseUnary()}
	}
	return n
}

func (p *Parser) parseUnary() Node {
	if p.cur.Type == TokenPlus || p.cur.Type == TokenMinus {
		op := p.cur.Type
		p.next(op)
		return &UnaryOpNode{op, p.parseUnary()}
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Node {
	switch p.cur.Type {
	case TokenNumber:
		n, _ := strconv.ParseFloat(p.cur.Value, 64)
		t := NumberNode{n}
		p.next(TokenNumber)
		return &t
	case TokenIdent:
		name := p.cur.Value
		p.next(TokenIdent)
		if p.cur.Type == TokenLParen || p.cur.Type == TokenLBracket || p.cur.Type == TokenLBrace {
			open := p.cur.Type
			p.next(open)
			var args []Node
			if p.cur.Type != TokenRParen && p.cur.Type != TokenRBracket && p.cur.Type != TokenRBrace {
				for {
					args = append(args, p.parseExpression())
					if p.cur.Type == TokenComma {
						p.next(TokenComma)
					} else {
						break
					}
				}
			}
			if p.cur.Type != TokenRParen && p.cur.Type != TokenRBracket && p.cur.Type != TokenRBrace {
				panic("expected closing bracket for function argument, got " + p.cur.String())
			}
			p.next(p.cur.Type)
			return &FuncCallNode{Name: name, Args: args}
		}
		return &ConstNode{Name: name}
	case TokenLParen, TokenLBracket, TokenLBrace:
		p.next(p.cur.Type)
		node := p.parseExpression()
		if p.cur.Type != TokenRParen && p.cur.Type != TokenRBracket && p.cur.Type != TokenRBrace {
			panic("expected closing bracket, got " + p.cur.String())
		}
		p.next(p.cur.Type)
		return node
	case TokenPlus, TokenMinus:
		return p.parseUnary()
	case TokenEOF:
		panic("unexpected end of input")
	default:
		panic("unhandled default case")
	}
}
