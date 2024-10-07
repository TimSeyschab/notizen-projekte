package lexer

import (
	"interpreter/token"
	"testing"
)

type TokenExpection struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
                   let ten = 10;
                   let add = fn(x, y) {
                      x + y;
                   };
                   let result = add(five, ten);
                   `

	tests := []TokenExpection{
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
		{token.FUNCTION, "fn"},
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

	checkTokenizedResult(input, tests, t)
}

func TestNumberDefinitions(t *testing.T) {
	input := `let five = 5;
                   let tenPointEight = 10.8;
                   let pointTwo = .2;
                   `

	tests := []TokenExpection{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "tenPointEight"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.8"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "pointTwo"},
		{token.ASSIGN, "="},
		{token.FLOAT, ".2"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	checkTokenizedResult(input, tests, t)
}

func TestIllegalNumber(t *testing.T) {
	input := `let five = 5.4.4;
                   `

	tests := []TokenExpection{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.ILLEGAL, "5.4.4"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	checkTokenizedResult(input, tests, t)
}

func checkTokenizedResult(input string, tests []TokenExpection, t *testing.T) {
	t.Helper()

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
