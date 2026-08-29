package lex

import (
	"fmt"
	"testing"
)

func TestMatchChars(t *testing.T) {
	tests := []struct{
		Candidate string
		Reference string
		Expected bool
	} {
		{ "", "", true },
		{ "A", "", false },
		{ "ABC", "ABC", true },
		{ "AB", "ABC", true },
		{ "A", "ABC", true },
		{ "ABD", "ABC", false },
		{ "ABC", "AB", false },
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Match %v:%v", test.Candidate, test.Reference)

		t.Run(testname, func(t *testing.T) {
			result := MatchChars([]byte(test.Candidate), []byte(test.Reference))

			if result != test.Expected {
				if result {
					t.Errorf("Matched non-matching strings \"%v\" and \"%v\"", test.Candidate, test.Reference)
				} else {
					t.Errorf("Did not match matching strings \"%v\" and \"%v\"", test.Candidate, test.Reference)
				}
			}
		})
	}
}

func TestLex(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []*Token
	}{
		// Empty input
		{
			Input:    "",
			Expected: []*Token{},
		},

		// Newline
		{
			Input: "\n",
			Expected: []*Token{
				{Type: TokenTypeNewLine, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
			},
		},

		// Parentheses and separator
		{
			Input: "( ) ,",
			Expected: []*Token{
				{Type: TokenTypeOpenParen, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeCloseParen, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeSeparator, Value: "", Start: TextLocation{Line: 0, Column: 4}, Length: 1},
			},
		},

		// Assignment
		{
			Input: "=",
			Expected: []*Token{
				{Type: TokenTypeAssignment, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
			},
		},

		// Arithmetic operators
		{
			Input: "+ - * / % ^",
			Expected: []*Token{
				{Type: TokenTypeAdd, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeSubtract, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeMultiply, Value: "", Start: TextLocation{Line: 0, Column: 4}, Length: 1},
				{Type: TokenTypeDivide, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 1},
				{Type: TokenTypeModulo, Value: "", Start: TextLocation{Line: 0, Column: 8}, Length: 1},
				{Type: TokenTypeExponent, Value: "", Start: TextLocation{Line: 0, Column: 10}, Length: 1},
			},
		},

		// Compound assignment operators
		{
			Input: "+= -= *= /= %= ^=",
			Expected: []*Token{
				{Type: TokenTypeAddAssign, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 2},
				{Type: TokenTypeSubtractAssign, Value: "", Start: TextLocation{Line: 0, Column: 3}, Length: 2},
				{Type: TokenTypeMultiplyAssign, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 2},
				{Type: TokenTypeDivideAssign, Value: "", Start: TextLocation{Line: 0, Column: 9}, Length: 2},
				{Type: TokenTypeModuloAssign, Value: "", Start: TextLocation{Line: 0, Column: 12}, Length: 2},
				{Type: TokenTypeExponentAssign, Value: "", Start: TextLocation{Line: 0, Column: 15}, Length: 2},
			},
		},

		// Increment / decrement
		{
			Input: "++ --",
			Expected: []*Token{
				{Type: TokenTypeIncrement, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 2},
				{Type: TokenTypeDecrement, Value: "", Start: TextLocation{Line: 0, Column: 3}, Length: 2},
			},
		},

		// Logical operators
		{
			Input: "& | !& !| !",
			Expected: []*Token{
				{Type: TokenTypeAND, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeOR, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeNAND, Value: "", Start: TextLocation{Line: 0, Column: 4}, Length: 2},
				{Type: TokenTypeNOR, Value: "", Start: TextLocation{Line: 0, Column: 7}, Length: 2},
				{Type: TokenTypeNOT, Value: "", Start: TextLocation{Line: 0, Column: 10}, Length: 1},
			},
		},

		// Comparison operators
		{
			Input: "== != < > <= >=",
			Expected: []*Token{
				{Type: TokenTypeEqual, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 2},
				{Type: TokenTypeNotEqual, Value: "", Start: TextLocation{Line: 0, Column: 3}, Length: 2},
				{Type: TokenTypeLessThan, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 1},
				{Type: TokenTypeGreaterThan, Value: "", Start: TextLocation{Line: 0, Column: 8}, Length: 1},
				{Type: TokenTypeLessThanOrEqual, Value: "", Start: TextLocation{Line: 0, Column: 10}, Length: 2},
				{Type: TokenTypeGreaterThanOrEqual, Value: "", Start: TextLocation{Line: 0, Column: 13}, Length: 2},
			},
		},

		// Keywords
		{
			Input: "func var",
			Expected: []*Token{
				{Type: TokenTypeFunction, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 4},
				{Type: TokenTypeVariable, Value: "", Start: TextLocation{Line: 0, Column: 5}, Length: 3},
			},
		},

		// Identifiers
		{
			Input: "foo bar myVariable _test",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "foo", Start: TextLocation{Line: 0, Column: 0}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "bar", Start: TextLocation{Line: 0, Column: 4}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "myVariable", Start: TextLocation{Line: 0, Column: 8}, Length: 10},
				{Type: TokenTypeIdentifier, Value: "_test", Start: TextLocation{Line: 0, Column: 19}, Length: 5},
			},
		},

		// Numbers
		{
			Input: "0 1 42 12345",
			Expected: []*Token{
				{Type: TokenTypeLiteralNumber, Value: "0", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "1", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "42", Start: TextLocation{Line: 0, Column: 4}, Length: 2},
				{Type: TokenTypeLiteralNumber, Value: "12345", Start: TextLocation{Line: 0, Column: 7}, Length: 5},
			},
		},

		// Function declaration
		{
			Input: "func foo()",
			Expected: []*Token{
				{Type: TokenTypeFunction, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 4},
				{Type: TokenTypeIdentifier, Value: "foo", Start: TextLocation{Line: 0, Column: 5}, Length: 3},
				{Type: TokenTypeOpenParen, Value: "", Start: TextLocation{Line: 0, Column: 8}, Length: 1},
				{Type: TokenTypeCloseParen, Value: "", Start: TextLocation{Line: 0, Column: 9}, Length: 1},
			},
		},

		// Variable assignment
		{
			Input: "var x = 42",
			Expected: []*Token{
				{Type: TokenTypeVariable, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 4}, Length: 1},
				{Type: TokenTypeAssignment, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "42", Start: TextLocation{Line: 0, Column: 8}, Length: 2},
			},
		},

		// Arithmetic expression
		{
			Input: "x + 10 * y",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeAdd, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "10", Start: TextLocation{Line: 0, Column: 4}, Length: 2},
				{Type: TokenTypeMultiply, Value: "", Start: TextLocation{Line: 0, Column: 7}, Length: 1},
				{Type: TokenTypeIdentifier, Value: "y", Start: TextLocation{Line: 0, Column: 9}, Length: 1},
			},
		},

		// Compound assignment expression
		{
			Input: "x += 5",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeAddAssign, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 2},
				{Type: TokenTypeLiteralNumber, Value: "5", Start: TextLocation{Line: 0, Column: 5}, Length: 1},
			},
		},

		// Comparisons
		{
			Input: "x >= 10",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeGreaterThanOrEqual, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 2},
				{Type: TokenTypeLiteralNumber, Value: "10", Start: TextLocation{Line: 0, Column: 5}, Length: 2},
			},
		},

		// Logical expression
		{
			Input: "x & y | !z",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeAND, Value: "", Start: TextLocation{Line: 0, Column: 2}, Length: 1},
				{Type: TokenTypeIdentifier, Value: "y", Start: TextLocation{Line: 0, Column: 4}, Length: 1},
				{Type: TokenTypeOR, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 1},
				{Type: TokenTypeNOT, Value: "", Start: TextLocation{Line: 0, Column: 8}, Length: 1},
				{Type: TokenTypeIdentifier, Value: "z", Start: TextLocation{Line: 0, Column: 9}, Length: 1},
			},
		},

		// Multiple lines
		{
			Input: "var x = 10\nvar y = 20",
			Expected: []*Token{
				{Type: TokenTypeVariable, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 0, Column: 4}, Length: 1},
				{Type: TokenTypeAssignment, Value: "", Start: TextLocation{Line: 0, Column: 6}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "10", Start: TextLocation{Line: 0, Column: 8}, Length: 2},
				{Type: TokenTypeNewLine, Value: "", Start: TextLocation{Line: 0, Column: 10}, Length: 1},
				{Type: TokenTypeVariable, Value: "", Start: TextLocation{Line: 1, Column: 0}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "y", Start: TextLocation{Line: 1, Column: 4}, Length: 1},
				{Type: TokenTypeAssignment, Value: "", Start: TextLocation{Line: 1, Column: 6}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "20", Start: TextLocation{Line: 1, Column: 8}, Length: 2},
			},
		},

		// Identifiers that match part of a keyword
		{
			Input: "v va f fu fun",
			Expected: []*Token{
				{Type: TokenTypeIdentifier, Value: "v", Start: TextLocation{Line: 0, Column: 0}, Length: 1},
				{Type: TokenTypeIdentifier, Value: "va", Start: TextLocation{Line: 0, Column: 2}, Length: 2},
				{Type: TokenTypeIdentifier, Value: "f", Start: TextLocation{Line: 0, Column: 5}, Length: 1},
				{Type: TokenTypeIdentifier, Value: "fu", Start: TextLocation{Line: 0, Column: 7}, Length: 2},
				{Type: TokenTypeIdentifier, Value: "fun", Start: TextLocation{Line: 0, Column: 10}, Length: 3},
			},
		},

		// Comment
		{
			Input: "	// This is a comment",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" This is a comment"*/, Start: TextLocation{Line: 0, Column: 1}, Length: 20},
			},
		},

		{
			Input: "//",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "", Start: TextLocation{Line: 0, Column: 0}, Length: 2},
			},
		},

		{
			Input: "//This is a comment",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*"This is a comment"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 19},
			},
		},

		{
			Input: "// var x = y * 4",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" var x = y * 4"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 16},
			},
		},

		{
			Input: "// var x = y * 4\nvar x = 42",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" var x = y * 4"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 16},
				{Type: TokenTypeNewLine, Value: "", Start: TextLocation{Line: 0, Column: 16}, Length: 1},
				{Type: TokenTypeVariable, Value: "", Start: TextLocation{Line: 1, Column: 0}, Length: 3},
				{Type: TokenTypeIdentifier, Value: "x", Start: TextLocation{Line: 1, Column: 4}, Length: 1},
				{Type: TokenTypeAssignment, Value: "", Start: TextLocation{Line: 1, Column: 6}, Length: 1},
				{Type: TokenTypeLiteralNumber, Value: "42", Start: TextLocation{Line: 1, Column: 8}, Length: 2},
			},
		},

		// Multi line comments
		{
			Input: "/* var x = y * 4\nvar x = 42",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" var x = y * 4\nvar x = 42"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 28},
			},
		},

		{
			Input: "/* var x = y * 4\nvar x = 42 */",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" var x = y * 4\nvar x = 42"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 31},
			},
		},

		{
			Input: "/* var x = y * 4\nvar x = 42 */",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*" var x = y * 4\nvar x = 42"*/, Start: TextLocation{Line: 0, Column: 0}, Length: 31},
			},
		},

		{
			Input: "/**/",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*""*/, Start: TextLocation{Line: 0, Column: 0}, Length: 4},
			},
		},

		{
			Input: "/*/",
			Expected: []*Token{
				{Type: TokenTypeComment, Value: "" /*""*/, Start: TextLocation{Line: 0, Column: 0}, Length: 3},
			},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Lex %v", i)

		t.Run(testname, func(t *testing.T) {
			result, err := Lex([]byte(test.Input), false)

			if err != nil {
				t.Errorf("Unexpected problem in lexing: %v", err)
			}

			if len(result) < len(test.Expected) {
				t.Errorf("Fewer output tokens than expected. Got %v, expected %v", len(result), len(test.Expected))
				return
			}

			for i, token := range test.Expected {
				if *result[i] != *token {
					t.Errorf("Found wrong token at %v. Found %v, expected %v", i, *result[i], *token)
				}
			}
		})
	}
}