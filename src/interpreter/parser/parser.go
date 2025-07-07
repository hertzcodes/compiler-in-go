package parser

import (
	"errors"
	"fmt"

	ast "github.com/hertzcodes/compiler-in-go/src/interpreter/parser/ast"
	"github.com/hertzcodes/compiler-in-go/src/interpreter/parser/ast/nodes"
	"github.com/hertzcodes/compiler-in-go/src/interpreter/token"

	"github.com/hertzcodes/compiler-in-go/src/interpreter/lexer"
)

type Parser struct {
	lexer *lexer.Lexer

	current token.Token
	peek    token.Token
	errors  []error
}

func New(l *lexer.Lexer) *Parser {

	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected %s, got %s instead", t, p.peek.Type)
	p.errors = append(p.errors, errors.New(msg))
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	pg := &ast.Program{}
	pg.Statements = []nodes.Statement{}

	for p.current.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			pg.Statements = append(pg.Statements, stmt)
		}
		p.nextToken()

	}
	return pg
}

func (p *Parser) parseStatement() nodes.Statement {
	switch p.current.Type {
	case token.LET:
		return p.parseLetStatement()

	default:
		return nil
	}
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peek.Type == t {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// IMPLMENTATION OF PARSERS

func (p *Parser) parseLetStatement() *nodes.LetStatement {
	stmt := &nodes.LetStatement{Token: p.current}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &nodes.Identifier{Token: p.current, Value: p.current.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for p.current.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}
