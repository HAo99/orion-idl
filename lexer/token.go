package lexer

import (
	piter "github.com/HAo99/orion-idl/common/peekiterator"
)

type Token interface {
	Type() TokenType
	Value() string
}

type TokenType int

const (
	TokenTypeKeyword = TokenType(iota)
	TokenTypeIdentifier
	TokenTypeOperator
	TokenTypeFundamentalType
	TokenTypeBracket
	TokenTypeInteger
	TokenTypeFloat
)

var (
	TokenTypeName = map[TokenType]string{
		TokenTypeKeyword:         "KEYWD",
		TokenTypeIdentifier:      "IDENT",
		TokenTypeOperator:        "OPERA",
		TokenTypeFundamentalType: "FTYPE",
		TokenTypeBracket:         "BRACK",
		TokenTypeInteger:         "INTEG",
		TokenTypeFloat:           "FLOAT",
	}
	TokenTypeValue = map[string]TokenType{
		"KEYWD": TokenTypeKeyword,
		"IDENT": TokenTypeIdentifier,
		"OPEAR": TokenTypeOperator,
		"FTYPE": TokenTypeFundamentalType,
		"BRACK": TokenTypeBracket,
		"INTEG": TokenTypeInteger,
		"FLOAT": TokenTypeFloat,
	}
)

func (x TokenType) IsKeyword() bool         { return x == TokenTypeKeyword }
func (x TokenType) IsIdentifier() bool      { return x == TokenTypeIdentifier }
func (x TokenType) IsOperator() bool        { return x == TokenTypeOperator }
func (x TokenType) IsFundamentalType() bool { return x == TokenTypeFundamentalType }
func (x TokenType) IsBracket() bool         { return x == TokenTypeBracket }
func (x TokenType) IsInteger() bool         { return x == TokenTypeInteger }
func (x TokenType) IsFloat() bool           { return x == TokenTypeFloat }

func (x TokenType) String() string {
	return TokenTypeName[x]
}

type token struct {
	typ TokenType
	val string
}

func (t token) Type() TokenType { return t.typ }
func (t token) Value() string   { return t.val }

func makeBracketToken(r rune) Token {
	return &token{
		typ: TokenTypeBracket,
		val: r2str(r),
	}
}

func makeNumberToken(iter piter.PeekIterator) Token {
	var (
		state = 0
		rs    = make([]rune, 0)
	)
	for iter.HasNext() {
		var (
			lookahead rune
			ok        bool
		)
		if lookahead, ok = iter.Peek().(rune); !ok {
			panic("lexer.makeNumberToken: invalid source")
		}
		switch state {
		case 0:
			switch {
			case lookahead == 0:
				state = 1
			case isNumber(lookahead):
				state = 2
			case lookahead == '+':
			case lookahead == '-':
				state = 3
			case lookahead == '.':
				state = 5
			}
		case 1:
			switch {
			case lookahead == '0':
				state = 1
			case isNumber(lookahead):
				state = 2
			case lookahead == '.':
				state = 4
			default:
				return &token{typ: TokenTypeInteger, val: string(rs)}
			}
		case 2:
			switch {
			case isNumber(lookahead):
				state = 2
			case lookahead == '.':
				state = 4
			default:
				return &token{typ: TokenTypeInteger, val: string(rs)}
			}
		case 3:
			switch {
			case isNumber(lookahead):
				state = 2
			case lookahead == '.':
				state = 5
			default:
				panic("lexer.makeNumberToken: invalid state")
			}
		case 4:
			switch {
			case lookahead == '.':
				panic("lexer.makeNumberToken: invalid state")
			case isNumber(lookahead):
				state = 4
			default:
				return &token{typ: TokenTypeFloat, val: string(rs)}
			}
		case 5:
			if isNumber(lookahead) {
				state = 4
			} else {
				panic("lexer.makeNumberToken: invalid state")
			}
		}
		iter.Next()
		rs = append(rs, lookahead)
	}
	panic("lexer.makeNumberToken: invalid state")
}

func makeIdentiferOrKeywordToken(iter piter.PeekIterator) Token {
	var rs = make([]rune, 0)
	for iter.HasNext() {
		var (
			lookahead rune
			ok        bool
		)
		if lookahead, ok = iter.Peek().(rune); !ok {
			panic("lexer.makeIdentiferOrKeywordToken: invalid source")
		}
		if isLiteral(lookahead) {
			rs = append(rs, lookahead)
		} else {
			break
		}
		iter.Next()
	}
	s := string(rs)
	if isKeyword(s) {
		return &token{typ: TokenTypeKeyword, val: s}
	}
	if isFundamentalType(s) {
		return &token{typ: TokenTypeFundamentalType, val: s}
	}
	return &token{typ: TokenTypeIdentifier, val: s}
}

func makeOperatorToken(iter piter.PeekIterator) Token {
	lookahead := iter.Peek()
	if lookahead == '=' {
		iter.Next()
		return &token{typ: TokenTypeOperator, val: "="}
	}
	if lookahead == '-' {
		iter.Next()
		lookahead = iter.Peek()
		if lookahead == '>' {
			iter.Next()
			return &token{typ: TokenTypeOperator, val: "->"}
		}
	}
	panic("lexer.makeOperatorToken: invalid state")
}
