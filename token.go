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
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	values := map[TokenType]string{
		TokenEof:    "EOF",
		TokenNumber: "NUMBER",
		TokenPlus:   "PLUS",
		TokenMinus:  "MINUS",
		TokenMul:    "MUL",
		TokenDiv:    "DIV",
		TokenPow:    "POW", // Added for exponentiation
	}
	return values[t.Type] + " (" + t.Value + ")"
}
