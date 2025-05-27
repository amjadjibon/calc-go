package calc

type TokenType int

const (
	TokenEOF TokenType = iota
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
	TokenMod
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	values := map[TokenType]string{
		TokenEOF:      "EOF",
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
		TokenMod:      "MOD",
	}
	return values[t.Type] + " (" + t.Value + ")"
}
