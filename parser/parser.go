package parser

import (
	"github.com/theonlynanu/go-interpreter/ast"
	"github.com/theonlynanu/go-interpreter/lexer"
	"github.com/theonlynanu/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Reads two tokens so currToken and peekToken are set to start
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Placeholder for writing test
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
