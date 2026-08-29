package lex

type PrimitiveToken struct {
	Text string
	Type TokenType
}

var (
	primitives = []PrimitiveToken{
		{"\n", TokenTypeNewLine},

		{"(", TokenTypeOpenParen},
		{")", TokenTypeCloseParen},
		{",", TokenTypeSeparator},

		{"=", TokenTypeAssignment},

		{"+", TokenTypeAdd},
		{"-", TokenTypeSubtract},
		{"*", TokenTypeMultiply},
		{"/", TokenTypeDivide},
		{"%", TokenTypeModulo},
		{"^", TokenTypeExponent},

		{"+=", TokenTypeAddAssign},
		{"-=", TokenTypeSubtractAssign},
		{"*=", TokenTypeMultiplyAssign},
		{"/=", TokenTypeDivideAssign},
		{"%=", TokenTypeModuloAssign},
		{"^=", TokenTypeExponentAssign},

		{"++", TokenTypeIncrement},
		{"--", TokenTypeDecrement},

		{"&", TokenTypeAND},
		{"|", TokenTypeOR},
		{"!&", TokenTypeNAND},
		{"!|", TokenTypeNOR},
		{"!", TokenTypeNOT},

		{"==", TokenTypeEqual},
		{"!=", TokenTypeNotEqual},
		{"<", TokenTypeLessThan},
		{">", TokenTypeGreaterThan},
		{"<=", TokenTypeLessThanOrEqual},
		{">=", TokenTypeGreaterThanOrEqual},

		{"func", TokenTypeFunction},
		{"var", TokenTypeVariable},
	}
)

type TokenMatcher func([]byte) *Token

var (
	matchers = []TokenMatcher{
		func(cand []byte) *Token {
			return MatchPrimitives(cand)
		},

		func(cand []byte) *Token {
			return MatchLiteralNumber(cand)
		},

		func(cand []byte) *Token {
			return MatchIdentifier(cand)
		},
	}
)

func MatchChars(candidate []byte, reference []byte) bool {
	if len(candidate) > len(reference) {
		return false
	}

	for i, char := range candidate {
		if char != reference[i] {
			return false
		}
	}

	return true
}

func IdentAllowed(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '_' || c == '-'
}

func NumberLiteralAllowed(c byte) bool {
	return (c >= '0' && c <= '9') || c == '.'
}

func MatchPrimitives(cand []byte) *Token {
	for _, primitive := range primitives {
		match := MatchChars(cand, []byte(primitive.Text))

		if match {
			return &Token{
				Type: primitive.Type,
			}
		}
	}

	return nil
}

func MatchLiteralNumber(cand []byte) *Token {
	for i, char := range cand {
		if !NumberLiteralAllowed(char) {
			return nil
		}

		if i == 0 && char == '.' {
			return nil
		}
	}

	return &Token{
		Type: TokenTypeLiteralNumber,
		Value: string(cand),
	}
}

func MatchIdentifier(cand []byte) *Token {
	for i, char := range cand {
		if !IdentAllowed(char) {
			return nil
		}

		if i == 0 && (char >= '0' && char <= '9') {
			return nil
		}
	}

	return &Token{
		Type: TokenTypeIdentifier,
		Value: string(cand),
	}
}

func MatchToken(cand []byte) *Token {
	for _, matcher := range matchers {
		t := matcher(cand)

		if t != nil{
			return t
		}
	}

	return nil
}
