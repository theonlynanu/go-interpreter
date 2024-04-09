package parser

import (
	"testing"

	"github.com/theonlynanu/go-interpreter/ast"
	"github.com/theonlynanu/go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	// There's a chance a mock here would make things easier, rather than calling
	// the lexer, but I like the idea of using readable source code instead.
	// Any bugs could be caused by the lexer too, so if there are
	// any unexplainable errors, check/test there
	input := `
	let x = 5;
	let y = 10;
	let foobar = 123456;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram returned nil.")
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let'. Got %q instead.", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. Got %T instead.", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %q, got %q instead.", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %q, got %q instead.", name, letStmt.TokenLiteral())
		return false
	}

	return true

}
