package lexer

import (
	"github.com/eroslang/gotoken"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{token.ILLEGAL, ""}
	switch l.ch {
	case '=':
		tok.Type = token.ASSIGN
		tok.Literal = string(l.ch)
	case '+':
		tok.Type = token.PLUS
		tok.Literal = string(l.ch)
	case ',':
		tok.Type = token.COMMA
		tok.Literal = string(l.ch)
	case ';':
		tok.Type = token.SEMICOLON
		tok.Literal = string(l.ch)
	case '(':
		tok.Type = token.LPAREN
		tok.Literal = string(l.ch)
	case ')':
		tok.Type = token.RPAREN
		tok.Literal = string(l.ch)
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = string(l.ch)
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = string(l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
