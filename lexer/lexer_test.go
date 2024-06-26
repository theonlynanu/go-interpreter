package lexer

import (
	"testing"

	"github.com/theonlynanu/go-interpreter/token"
)

type tokenTestType struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {

	t.Run("Basic symbol recognition", func(t *testing.T) {
		input := "=+(){},;-!*/<>"

		tests := []tokenTestType{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.MINUS, "-"},
			{token.BANG, "!"},
			{token.ASTERISK, "*"},
			{token.SLASH, "/"},
			{token.LT, "<"},
			{token.GT, ">"},
			{token.EOF, ""},
		}

		assertTokens(t, input, tests)
	})

	t.Run("Initialize variables and functions", func(t *testing.T) {
		input := `let five = 5;
		let ten = 10;
		
		let add = function (x, y) {
			x + y;
		};
		
		let result = add(five, ten);`

		tests := []tokenTestType{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "function"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},

			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		assertTokens(t, input, tests)
	})

	t.Run("Basic conditionals and keywords", func(t *testing.T) {
		input := `let ten = 10;
		let five = 5;
		
		let greaterThan = function(x, y) {
			if (x > y) {
				return true;
			} else {
				return false;
			}
		}`
		tests := []tokenTestType{
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "greaterThan"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "function"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},

			{token.LBRACE, "{"},
			{token.IF, "if"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.GT, ">"},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.TRUE, "true"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.ELSE, "else"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.FALSE, "false"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.RBRACE, "}"},
			{token.EOF, ""},
		}

		assertTokens(t, input, tests)
	})

	t.Run("Test two-character tokens", func(t *testing.T) {
		input := `10 == 10;
		5 != 10;`

		tests := []tokenTestType{
			{token.INT, "10"},
			{token.EQ, "=="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},

			{token.INT, "5"},
			{token.NOT_EQ, "!="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
		}

		assertTokens(t, input, tests)
	})
}

func assertTokens(t testing.TB, input string, tests []tokenTestType) {
	t.Helper()

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Tests[%d] - tokentype incorrect. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Tests[%d] - literal incorrect. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
