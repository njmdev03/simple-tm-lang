package lex


type TextLocation struct {
	Line   int
	Column int
}

type TokenType int

const (
	TokenTypeUnknown TokenType = iota
	TokenTypeNewLine

	TokenTypeComment

	TokenTypeFunction
	TokenTypeVariable

	TokenTypeIdentifier

	TokenTypeOpenParen
	TokenTypeCloseParen

	TokenTypeSeparator

	TokenTypeOpenScope
	TokenTypeCloseScope

	TokenTypeLiteralNumber

	TokenTypeAssignment

	TokenTypeAdd
	TokenTypeSubtract
	TokenTypeMultiply
	TokenTypeDivide
	TokenTypeModulo
	TokenTypeExponent

	TokenTypeAddAssign
	TokenTypeSubtractAssign
	TokenTypeMultiplyAssign
	TokenTypeDivideAssign
	TokenTypeModuloAssign
	TokenTypeExponentAssign

	TokenTypeIncrement
	TokenTypeDecrement

	TokenTypeAND
	TokenTypeNAND
	TokenTypeOR
	TokenTypeNOR
	TokenTypeNOT

	TokenTypeEqual
	TokenTypeNotEqual
	TokenTypeLessThan
	TokenTypeGreaterThan
	TokenTypeLessThanOrEqual
	TokenTypeGreaterThanOrEqual
)

type Token struct {
	Type   TokenType
	Value  string
	Start  TextLocation
	Length int
}

var (
	stripTokens = []TokenType{
		TokenTypeNewLine,
		TokenTypeComment,
	}
)

func shouldStrip(t TokenType) bool {
	for _, s := range stripTokens {
		if t == s {
			return true
		}
	}

	return false
}

func Lex(input []byte, strip bool) ([]*Token, error) {

	tokens := make([]*Token, 0)

	position := TextLocation{Line: 0, Column: 0}

	// var lastToken *Token
	var lastString []byte

	for i, char := range input {

		currentString := append(lastString, char)

		if i < len(input) - 1 {
			nextString := append(currentString, input[i+1])

			if MatchToken(nextString, false) != nil {
				lastString = currentString

				continue
			}
		}

		currentToken := MatchToken(currentString, true)

		if currentToken != nil {
			currentToken.Start = position
			currentToken.Length = len(currentString)

			position.Column += len(currentString)

			if currentToken.Type == TokenTypeNewLine {
				position.Column = 0
				position.Line += 1
			}

			if strip {
				if !shouldStrip(currentToken.Type) {
					tokens = append(tokens, currentToken)
				}
			} else {
				tokens = append(tokens, currentToken)
			}
		} else {
			position.Column += 1
		}

		lastString = []byte{}
	}

	return tokens, nil
}