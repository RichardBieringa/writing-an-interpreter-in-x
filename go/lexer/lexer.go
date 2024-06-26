// Pagage lexer contains an implementation of the lexer for the monkey
// programing language.
//
// The Lexer, sometimes called tokenizer or scanner, consumes the sources code
// of a program in textual form, and produces meaningful tokens. These tokens
// can range from language specific keywords such as `function` or `return` to
// identifiers used in variables and common symbols that carry a meaning.
//
// During the tokenization/lexing process we go over the input text, determine
// the type of token, and create a data structure to hold the type of token and
// the actual string literal content of it. This is useful for example in the
// case of an identifier or a number, where we want to know which number it,
// represented or wich identifier was used to store a value.
//
// It converts text like `for(int i=0;i <3; ++i)` to a tokens that make it
// easier to work with, deal with white space, etc.
// Example:
// for = FOR
// ( = OPEN_PAREN
// int = SYMBOL
// i = IDENTIFIER
// 0 = NUMBER
// ; = SEMI_COLON
// etc...
package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input string

	position     int  // current index position in the source code
	readPosition int  // position + 1
	currentChar  byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (l *Lexer) readChar() {
	l.currentChar = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	l.eatWhiteSpace()

	t := token.Token{
		Literal: string(l.currentChar),
	}

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			t.Type = token.EQ
			t.Literal = string(l.currentChar) + string(l.peekChar())
			l.readChar()
		} else {
			t.Type = token.ASSIGN
		}
	case '+':
		t.Type = token.PLUS
	case '-':
		t.Type = token.MINUS
	case '/':
		t.Type = token.SLASH
	case '*':
		t.Type = token.ASTERISK
	case '!':
		if l.peekChar() == '=' {
			t.Type = token.NOT_EQ
			t.Literal = string(l.currentChar) + string(l.peekChar())
			l.readChar()
		} else {
			t.Type = token.BANG
		}
	case ',':
		t.Type = token.COMMA
	case ';':
		t.Type = token.SEMICOLON
	case '(':
		t.Type = token.LPAREN
	case ')':
		t.Type = token.RPAREN
	case '{':
		t.Type = token.LBRACE
	case '}':
		t.Type = token.RBRACE
	case '<':
		t.Type = token.LT
	case '>':
		t.Type = token.GT
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = getIdentifier(t.Literal)
			return t
		} else if isNumber(l.currentChar) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t.Type = token.ILLEGAL
		}
	}

	l.readChar()
	return t
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func getIdentifier(identifier string) token.TokenType {
	keywords := map[string]token.TokenType{
		"let":    token.LET,
		"fn":     token.FUNCTION,
		"true":   token.TRUE,
		"false":  token.FALSE,
		"return": token.RETURN,
		"if":     token.IF,
		"else":   token.ELSE,
	}

	tokenType, ok := keywords[identifier]
	if ok {
		return tokenType
	}

	return token.IDENT
}

func (l *Lexer) readIdentifier() string {
	left := l.position

	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[left:l.position]
}

func (l *Lexer) readNumber() string {
	left := l.position

	for isNumber(l.currentChar) {
		l.readChar()
	}

	return l.input[left:l.position]
}

func (l *Lexer) eatWhiteSpace() {
	for isWhitespace(l.currentChar) {
		l.readChar()
	}
}
