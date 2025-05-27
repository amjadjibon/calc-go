package calc

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos:   0,
	}
}

func (l *Lexer) read() rune {
	if l.pos >= len(l.input) {
		return 0 // EOF
	}
	r := rune(l.input[l.pos])
	l.pos++
	return r
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0 // EOF
	}
	return rune(l.input[l.pos])
}

func (l *Lexer) NextToken() Token {
	for {
		r := l.read()
		if r == ' ' || r == '\t' || r == '\n' {
			continue
		} else if r == '+' {
			return Token{Type: TokenPlus, Value: string(r)}
		} else if r == '-' {
			return Token{Type: TokenMinus, Value: string(r)}
		} else if r == '*' {
			return Token{Type: TokenMul, Value: string(r)}
		} else if r == '/' {
			return Token{Type: TokenDiv, Value: string(r)}
		} else if r == '^' {
			return Token{Type: TokenPow, Value: string(r)}
		} else if r == '(' {
			return Token{Type: TokenLParen, Value: string(r)}
		} else if r == ')' {
			return Token{Type: TokenRParen, Value: string(r)}
		} else if r == '[' {
			return Token{Type: TokenLBracket, Value: string(r)}
		} else if r == ']' {
			return Token{Type: TokenRBracket, Value: string(r)}
		} else if r == '{' {
			return Token{Type: TokenLBrace, Value: string(r)}
		} else if r == '}' {
			return Token{Type: TokenRBrace, Value: string(r)}
		} else if r == ',' {
			return Token{Type: TokenComma, Value: string(r)}
		} else if r >= '0' && r <= '9' {
			value := string(r)
			decimalFound := false
			for {
				r = l.peek()
				if r >= '0' && r <= '9' {
					value += string(l.read())
				} else if r == '.' && !decimalFound {
					decimalFound = true
					value += string(l.read())
				} else {
					break
				}
			}
			return Token{Type: TokenNumber, Value: value}
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			value := string(r)
			for {
				r = l.peek()
				if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
					value += string(l.read())
				} else {
					break
				}
			}
			return Token{Type: TokenIdent, Value: value}
		} else if r == 0 {
			return Token{Type: TokenEOF, Value: ""}
		}

		panic("unexpected character: " + string(r))
	}
}
