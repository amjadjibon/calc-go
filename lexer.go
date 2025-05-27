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
		switch r {
		case ' ', '\t', '\n':
			continue
		case '+':
			return Token{Type: TokenPlus, Value: string(r)}
		case '-':
			return Token{Type: TokenMinus, Value: string(r)}
		case '*':
			return Token{Type: TokenMul, Value: string(r)}
		case '/':
			return Token{Type: TokenDiv, Value: string(r)}
		case '^':
			return Token{Type: TokenPow, Value: string(r)}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			value := string(r)
			for {
				r = l.peek()
				if r >= '0' && r <= '9' {
					value += string(l.read())
				} else {
					break
				}
			}
			return Token{Type: TokenNumber, Value: value}
		case 0:
			return Token{Type: TokenEof, Value: ""}
		default:
			panic("unexpected character: " + string(r))
		}
	}
}
