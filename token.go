package calc

type TokenType int

const (
	TokenEof TokenType = iota
	TokenNumber
	TokenPlus
	TokenMinus
	TokenMul
	TokenDiv
	TokenPow
	TokenLParen
	TokenRParen
	TokenLBracket
	TokenRBracket
	TokenLBrace
	TokenRBrace
	TokenIdent
	TokenComma
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	values := map[TokenType]string{
		TokenEof:      "EOF",
		TokenNumber:   "NUMBER",
		TokenPlus:     "PLUS",
		TokenMinus:    "MINUS",
		TokenMul:      "MUL",
		TokenDiv:      "DIV",
		TokenPow:      "POW",
		TokenLParen:   "LPAREN",
		TokenRParen:   "RPAREN",
		TokenLBracket: "LBRACKET",
		TokenRBracket: "RBRACKET",
		TokenLBrace:   "LBRACE",
		TokenRBrace:   "RBRACE",
		TokenIdent:    "IDENT",
		TokenComma:    "COMMA",
	}
	return values[t.Type] + " (" + t.Value + ")"
}
