package lex

// A simple token which exactly matches a fixed string character for character
type SimpleToken struct {
	Text string
	Type TokenType
}

var (
	// primitives = []PrimitiveToken{
	// 	{"\n", TokenTypeNewLine},

	// 	{"(", TokenTypeOpenParen},
	// 	{")", TokenTypeCloseParen},
	// 	{",", TokenTypeSeparator},

	// 	{"{", TokenTypeOpenParen},
	// 	{"}", TokenTypeCloseParen},

	// 	{"=", TokenTypeAssignment},

	// 	{"+", TokenTypeAdd},
	// 	{"-", TokenTypeSubtract},
	// 	{"*", TokenTypeMultiply},
	// 	{"/", TokenTypeDivide},
	// 	{"%", TokenTypeModulo},
	// 	{"^", TokenTypeExponent},

	// 	{"+=", TokenTypeAddAssign},
	// 	{"-=", TokenTypeSubtractAssign},
	// 	{"*=", TokenTypeMultiplyAssign},
	// 	{"/=", TokenTypeDivideAssign},
	// 	{"%=", TokenTypeModuloAssign},
	// 	{"^=", TokenTypeExponentAssign},

	// 	{"++", TokenTypeIncrement},
	// 	{"--", TokenTypeDecrement},

	// 	{"&", TokenTypeAND},
	// 	{"|", TokenTypeOR},
	// 	{"!&", TokenTypeNAND},
	// 	{"!|", TokenTypeNOR},
	// 	{"!", TokenTypeNOT},

	// 	{"==", TokenTypeEqual},
	// 	{"!=", TokenTypeNotEqual},
	// 	{"<", TokenTypeLessThan},
	// 	{">", TokenTypeGreaterThan},
	// 	{"<=", TokenTypeLessThanOrEqual},
	// 	{">=", TokenTypeGreaterThanOrEqual},

	// 	{"func", TokenTypeFunction},
	// 	{"var", TokenTypeVariable},
	// }

	operators = []SimpleToken{
		{"(", TokenTypeOpenParen},
		{")", TokenTypeCloseParen},
		{",", TokenTypeSeparator},

		{"{", TokenTypeOpenParen},
		{"}", TokenTypeCloseParen},

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
	}

	keywords = []SimpleToken{
		{"func", TokenTypeFunction},
		{"var", TokenTypeVariable},
	}

	specialCharacters = []SimpleToken{
		{"\n", TokenTypeNewLine},
	}
)

type TokenMatcher func([]byte, bool) *Token

var (
	matchers = []TokenMatcher{
		func(cand []byte, s bool) *Token {
			return MatchSimpleTokens(cand, specialCharacters, s)
		},

		MatchSingleLineComment,
		MatchMultiLineComment,

		func(cand []byte, s bool) *Token {
			return MatchSimpleTokens(cand, operators, s)
		},

		func(cand []byte, s bool) *Token {
			return MatchSimpleTokens(cand, keywords, s)
		},

		MatchLiteralNumber,
		MatchIdentifier,
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

func MatchSimpleTokens(cand []byte, simpleTokens []SimpleToken, strict bool) *Token {
	for _, primitive := range simpleTokens {
		prim := []byte(primitive.Text)

		match := MatchChars(cand, prim)

		if match && !strict {
			return &Token{
				Type: primitive.Type,
			}
		} else if match && len(cand) == len(prim) {
			return &Token{
				Type: primitive.Type,
			}
		}
	}

	return nil
}

func MatchLiteralNumber(cand []byte, strict bool) *Token {
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

func MatchIdentifier(cand []byte, strict bool) *Token {
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

func MatchSingleLineComment(cand []byte, strict bool) *Token {
	for i, char := range cand {
		if (i == 0 || i == 1) && char != '/' {
			return nil
		} else if char == '\n' {
			return nil
		}
	}

	if strict && len(cand) < 2 {
		return nil
	}

	return &Token{
		Type: TokenTypeComment,
	}
}

func MatchMultiLineComment(cand []byte, strict bool) *Token {
	for i, char := range cand {
		if i == 0 && char != '/' {
			return nil
		} else if i == 1 && char != '*' {
			return nil
		}

		if i > 1 && cand[i - 1] == '/' && cand[i - 2] == '*' {
			return nil
		}
	}

	if strict && (len(cand) < 2 || cand[0] != '/' || cand[1] != '*') {
		return nil
	}

	var val string

	if len(cand) >= 4 {
		
	}

	return &Token{
		Type: TokenTypeComment,
		Value: val,
	}
}

func MatchToken(cand []byte, strict bool) *Token {
	for _, matcher := range matchers {
		t := matcher(cand, strict)

		if t != nil{
			return t
		}
	}

	return nil
}
