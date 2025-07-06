package lexer

import (
	"github.com/hertzcodes/compiler-in-go/src/interpreter/token"
	"github.com/hertzcodes/compiler-in-go/src/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(in string) *Lexer {
	lex := &Lexer{input: in}
	lex.readChar()
	return lex
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) { // TODO: support all unicode range
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer) peekChar() byte { // same as readChar. doesn't modify l.position and l.ch
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position] // NOTE: This is only parsing 1 byte ASCIIs
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

// NextToken method creates tokens

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' { // TOKEN: generalize this & handle (==)+ (more than 2 equal signs strings)
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			t = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		t = token.New(token.PLUS, l.ch)
	case '-':
		t = token.New(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.NOT_EQUAL, Literal: string(ch) + string(l.ch)}
		} else {
			t = token.New(token.BANG, l.ch)
		}
	case '/':
		t = token.New(token.SLASH, l.ch)
	case '*':
		t = token.New(token.ASTERISK, l.ch)
	case '<':
		t = token.New(token.LT, l.ch)
	case '>':
		t = token.New(token.GT, l.ch)
	case ';':
		t = token.New(token.SEMICOLON, l.ch)
	case '(':
		t = token.New(token.LPAREN, l.ch)
	case ')':
		t = token.New(token.RPAREN, l.ch)
	case ',':
		t = token.New(token.COMMA, l.ch)
	case '{':
		t = token.New(token.LBRACE, l.ch)
	case '}':
		t = token.New(token.RBRACE, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF

	default:
		if utils.IsLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if utils.IsDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}
