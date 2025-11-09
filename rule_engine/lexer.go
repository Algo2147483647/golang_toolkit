package rule_engine

import (
	"fmt"
	"unicode"
)

// TokenType represents the type of a token
type TokenType int

// Token types
const (
	TokenEOF TokenType = iota
	TokenIdentifier
	TokenNumber
	TokenString
	TokenOperator
	TokenKeyword
	TokenLParen
	TokenRParen
	TokenLBrace
	TokenRBrace
	TokenLBracket
	TokenRBracket
	TokenComma
	TokenDot
	TokenSemicolon
	TokenError
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// String returns a string representation of the token
func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %d, Literal: %s, Line: %d, Column: %d}", t.Type, t.Literal, t.Line, t.Column)
}

// Lexer is a lexical analyzer for rule expressions
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// NewLexer creates a new lexer for the given input
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 1}
	l.readChar()
	return l
}

// readChar reads the next character from the input and updates position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL"
	} else {
		l.ch = rune(l.input[l.readPosition])
	}

	l.position = l.readPosition
	l.readPosition++

	// Update line and column tracking
	if l.ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
}

// peekChar returns the next character without advancing the lexer
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	var tok Token

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenOperator, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenOperator, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenOperator, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenOperator, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenError, Literal: fmt.Sprintf("illegal character '%c'", l.ch), Line: l.line, Column: l.column}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Literal: string(ch) + string(l.ch), Line: l.line, Column: l.column - 1}
		} else {
			tok = Token{Type: TokenError, Literal: fmt.Sprintf("illegal character '%c'", l.ch), Line: l.line, Column: l.column}
		}
	case '+', '-', '*', '/', '%':
		tok = Token{Type: TokenOperator, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '(':
		tok = Token{Type: TokenLParen, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ')':
		tok = Token{Type: TokenRParen, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '{':
		tok = Token{Type: TokenLBrace, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '}':
		tok = Token{Type: TokenRBrace, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '[':
		tok = Token{Type: TokenLBracket, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ']':
		tok = Token{Type: TokenRBracket, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ',':
		tok = Token{Type: TokenComma, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '.':
		tok = Token{Type: TokenDot, Literal: string(l.ch), Line: l.line, Column: l.column}
	case ';':
		tok = Token{Type: TokenSemicolon, Literal: string(l.ch), Line: l.line, Column: l.column}
	case '"', '\'':
		tok = l.readString(l.ch)
	case 0:
		tok = Token{Type: TokenEOF, Literal: "", Line: l.line, Column: l.column}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.lookupIdentifier(tok.Literal)
			tok.Line = l.line
			tok.Column = l.column - len(tok.Literal)
			return tok
		} else if isDigit(l.ch) || (l.ch == '-' && isDigit(l.peekChar())) {
			tok.Literal, tok.Type = l.readNumber()
			tok.Line = l.line
			tok.Column = l.column - len(tok.Literal)
			return tok
		} else {
			tok = Token{Type: TokenError, Literal: fmt.Sprintf("illegal character '%c'", l.ch), Line: l.line, Column: l.column}
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) && l.ch != 0 {
		l.readChar()
	}
}

// readIdentifier reads an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' || l.ch == '$' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number (integer or float)
func (l *Lexer) readNumber() (string, TokenType) {
	position := l.position

	// Handle negative numbers
	if l.ch == '-' {
		l.readChar()
	}

	// Read digits
	for isDigit(l.ch) {
		l.readChar()
	}

	// Check for decimal point
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar() // consume the '.'
		for isDigit(l.ch) {
			l.readChar()
		}
		return l.input[position:l.position], TokenNumber
	}

	return l.input[position:l.position], TokenNumber
}

// readString reads a string literal
func (l *Lexer) readString(quote rune) Token {
	position := l.position + 1 // skip opening quote
	for {
		l.readChar()
		if l.ch == rune(quote) || l.ch == 0 {
			break
		}
	}

	// If we reached EOF without closing quote
	if l.ch == 0 {
		return Token{Type: TokenError, Literal: "unterminated string literal", Line: l.line, Column: l.column}
	}

	// Return the string content (without quotes)
	literal := l.input[position:l.position]
	return Token{Type: TokenString, Literal: literal, Line: l.line, Column: l.column - len(literal) - 1}
}

// lookupIdentifier checks if an identifier is a keyword
func (l *Lexer) lookupIdentifier(ident string) TokenType {
	keywords := map[string]TokenType{
		"true":  TokenKeyword,
		"false": TokenKeyword,
		"in":    TokenKeyword,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return TokenIdentifier
}
