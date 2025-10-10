package rule_engine

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Parser is a simple parser for rule expressions
type Parser struct {
	expr string
	pos  int
}

// NewParser creates a new parser for the given expression
func NewParser(expr string) *Parser {
	return &Parser{
		expr: strings.TrimSpace(expr),
		pos:  0,
	}
}

// Parse parses the expression and returns the root ExprNode
func (p *Parser) Parse() (*ExprNode, error) {
	node, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// Check if we consumed the entire expression
	if p.pos < len(p.expr) {
		return nil, fmt.Errorf("unexpected token at end of expression: %s", p.expr[p.pos:])
	}

	return node, nil
}

// parseExpression parses logical expressions (||)
func (p *Parser) parseExpression() (*ExprNode, error) {
	left, err := p.parseAndExpression()
	if err != nil {
		return nil, err
	}

	for p.peek() == '|' {
		p.consume('|')
		p.consume('|')
		right, err := p.parseAndExpression()
		if err != nil {
			return nil, err
		}
		left = NewBinaryNode(OpOr, left, right)
	}

	return left, nil
}

// parseAndExpression parses logical AND expressions (&&)
func (p *Parser) parseAndExpression() (*ExprNode, error) {
	left, err := p.parseComparison()
	if err != nil {
		return nil, err
	}

	for p.peek() == '&' {
		p.consume('&')
		p.consume('&')
		right, err := p.parseComparison()
		if err != nil {
			return nil, err
		}
		left = NewBinaryNode(OpAnd, left, right)
	}

	return left, nil
}

// parseComparison parses comparison expressions (=, !=, <, <=, >, >=, in)
func (p *Parser) parseComparison() (*ExprNode, error) {
	left, err := p.parseAdditive()
	if err != nil {
		return nil, err
	}

	// Check for comparison operators
	op := p.parseOperator()
	if op != "" {
		right, err := p.parseAdditive()
		if err != nil {
			return nil, err
		}

		// Validate operator
		if !IsComparisonOperator(op) {
			return nil, fmt.Errorf("invalid comparison operator: %s", op)
		}

		return NewBinaryNode(op, left, right), nil
	}

	return left, nil
}

// parseAdditive parses additive expressions (+, -)
func (p *Parser) parseAdditive() (*ExprNode, error) {
	left, err := p.parseMultiplicative()
	if err != nil {
		return nil, err
	}

	for {
		op := p.peek()
		if op == '+' {
			p.consume('+')
			right, err := p.parseMultiplicative()
			if err != nil {
				return nil, err
			}
			left = NewBinaryNode(OpAdd, left, right)
		} else if op == '-' {
			p.consume('-')
			right, err := p.parseMultiplicative()
			if err != nil {
				return nil, err
			}
			left = NewBinaryNode(OpSubtract, left, right)
		} else {
			break
		}
	}

	return left, nil
}

// parseMultiplicative parses multiplicative expressions (*, /, %)
func (p *Parser) parseMultiplicative() (*ExprNode, error) {
	left, err := p.parseUnary()
	if err != nil {
		return nil, err
	}

	for {
		op := p.peek()
		if op == '*' {
			p.consume('*')
			right, err := p.parseUnary()
			if err != nil {
				return nil, err
			}
			left = NewBinaryNode(OpMultiply, left, right)
		} else if op == '/' {
			p.consume('/')
			right, err := p.parseUnary()
			if err != nil {
				return nil, err
			}
			left = NewBinaryNode(OpDivide, left, right)
		} else if op == '%' {
			p.consume('%')
			right, err := p.parseUnary()
			if err != nil {
				return nil, err
			}
			left = NewBinaryNode(OpMod, left, right)
		} else {
			break
		}
	}

	return left, nil
}

// parseUnary parses unary expressions (!)
func (p *Parser) parseUnary() (*ExprNode, error) {
	if p.peek() == '!' {
		p.consume('!')
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		return NewUnaryNode(OpNot, right), nil
	}

	return p.parsePrimary()
}

// parsePrimary parses primary expressions (literals, variables, parentheses)
func (p *Parser) parsePrimary() (*ExprNode, error) {
	p.skipWhitespace()

	if p.pos >= len(p.expr) {
		return nil, fmt.Errorf("unexpected end of expression")
	}

	// Handle parentheses
	if p.peek() == '(' {
		p.consume('(')
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		p.consume(')')
		return node, nil
	}

	// Handle string literals (enclosed in quotes)
	if p.peek() == '"' || p.peek() == '\'' {
		return p.parseStringLiteral()
	}

	// Handle numbers
	if isDigit(p.peek()) || p.peek() == '-' {
		return p.parseNumber()
	}

	// Handle variables and booleans
	return p.parseIdentifier()
}

// parseStringLiteral parses a string literal
func (p *Parser) parseStringLiteral() (*ExprNode, error) {
	quote := p.peek()
	p.consume(rune(quote))

	var sb strings.Builder
	for p.pos < len(p.expr) && p.peek() != rune(quote) {
		if p.peek() == '\\' {
			p.consume('\\')
			if p.pos < len(p.expr) {
				sb.WriteRune(p.peek())
				p.consume(p.peek())
			}
		} else {
			sb.WriteRune(p.peek())
			p.consume(p.peek())
		}
	}

	if p.pos >= len(p.expr) {
		return nil, fmt.Errorf("unterminated string literal")
	}

	p.consume(rune(quote))
	return NewValueNode(sb.String()), nil
}

// parseNumber parses a numeric literal
func (p *Parser) parseNumber() (*ExprNode, error) {
	start := p.pos

	// Handle negative numbers
	if p.peek() == '-' {
		p.consume('-')
	}

	// Parse digits
	for p.pos < len(p.expr) && (isDigit(p.peek()) || p.peek() == '.') {
		p.consume(p.peek())
	}

	numStr := p.expr[start:p.pos]

	// Try to parse as integer first
	if i, err := strconv.ParseInt(numStr, 10, 64); err == nil {
		return NewValueNode(i), nil
	}

	// Try to parse as float
	if f, err := strconv.ParseFloat(numStr, 64); err == nil {
		return NewValueNode(f), nil
	}

	return nil, fmt.Errorf("invalid number: %s", numStr)
}

// parseIdentifier parses identifiers (variables, booleans)
func (p *Parser) parseIdentifier() (*ExprNode, error) {
	p.skipWhitespace()
	start := p.pos

	// Parse identifier
	for p.pos < len(p.expr) && (isLetter(p.peek()) || isDigit(p.peek()) || p.peek() == '_' || p.peek() == '$') {
		p.consume(p.peek())
	}

	id := p.expr[start:p.pos]

	// Check for boolean literals
	if id == "true" {
		return NewValueNode(true), nil
	}
	if id == "false" {
		return NewValueNode(false), nil
	}

	// Treat as variable
	return NewValueNode("$" + id), nil
}

// parseOperator parses an operator
func (p *Parser) parseOperator() string {
	p.skipWhitespace()

	if p.pos >= len(p.expr) {
		return ""
	}

	// Multi-character operators
	if p.pos <= len(p.expr)-2 {
		op := p.expr[p.pos : p.pos+2]
		switch op {
		case "==":
			p.consume('=')
			p.consume('=')
			return OpEqual
		case "!=":
			p.consume('!')
			p.consume('=')
			return OpNotEqual
		case "<=":
			p.consume('<')
			p.consume('=')
			return OpLessEqual
		case ">=":
			p.consume('>')
			p.consume('=')
			return OpGreaterEqual
		}
	}

	// Single-character operators
	switch p.peek() {
	case '<':
		p.consume('<')
		return OpLessThan
	case '>':
		p.consume('>')
		return OpGreaterThan
	case '=':
		p.consume('=')
		return OpEqual
	}

	return ""
}

// Helper methods

func (p *Parser) peek() rune {
	if p.pos >= len(p.expr) {
		return 0
	}
	return rune(p.expr[p.pos])
}

func (p *Parser) consume(expected rune) {
	if p.pos >= len(p.expr) || rune(p.expr[p.pos]) != expected {
		panic(fmt.Sprintf("expected '%c' but found '%c' at position %d", expected, p.peek(), p.pos))
	}
	p.pos++
}

func (p *Parser) skipWhitespace() {
	for p.pos < len(p.expr) && unicode.IsSpace(rune(p.expr[p.pos])) {
		p.pos++
	}
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// ParseExpression is a convenience function to parse an expression string
func ParseExpression(expr string) (*ExprNode, error) {
	parser := NewParser(expr)
	return parser.Parse()
}
