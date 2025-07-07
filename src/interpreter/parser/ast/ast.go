package ast

import "github.com/hertzcodes/compiler-in-go/src/interpreter/parser/ast/nodes"

type Program struct {
	Statements []nodes.Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}