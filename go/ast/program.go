package ast

import "bytes"

// The root of every parse tree of the programming language
// A program contains a series of statements, and is the root of the AST
//
// Program
// ├── Statement - `let x = 5;`
// ├── Statement - `let y = 8;`
// ├── Statement - `let z = x + y;`
// └── Statement - `return z;`
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, statement := range p.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}
