package lex


type TextLocation struct {
	Line   int
	Column int
}

type TokenType int

const (
	TokenTypeUnknown TokenType = iota
	TokenTypeNewLine

	TokenTypeFunction
	TokenTypeVariable

	TokenTypeIdentifier

	TokenTypeOpenParen
	TokenTypeCloseParen

	TokenTypeSeparator

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

func Lex(input []byte) ([]*Token, error) {

	tokens := make([]*Token, 0)

	position := TextLocation{Line: 1, Column: 0}

	var lastToken *Token
	var currentString []byte

	for _, char := range input {
		position.Column += 1

		currentString = append(currentString, char)

		currentToken := MatchToken(currentString)

		if currentToken == nil {
			if lastToken != nil {
				tokens = append(tokens, lastToken)
			}
			currentString = []byte{char}
			lastToken = MatchToken(currentString)
		} else {
			lastToken = currentToken
			lastToken.Length = len(currentString)

			if len(currentString) == 1 {
				lastToken.Start = position
			}
		}
	}

	return tokens, nil
}