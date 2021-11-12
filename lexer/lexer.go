package lexer

import (
	"encoding/json"
	"strings"

	piter "github.com/HAo99/orion-idl/common/peekiterator"
)

type tokenSource struct {
	tokens []Token
	cur    int
}

func (s *tokenSource) Read() (tkn interface{}) {
	tkn = s.tokens[s.cur]
	s.cur++
	return
}

func (s *tokenSource) Empty() bool {
	return s.cur >= len(s.tokens)
}

func (s *tokenSource) String() string {
	bdr := strings.Builder{}
	for _, tkn := range s.tokens {
		bdr.WriteByte('<')
		bdr.WriteString(tkn.Type().String())
		bdr.WriteByte('>')

		bdr.WriteByte(':')
		bdr.WriteByte(' ')
		bdr.WriteString(tkn.Value())
		bdr.WriteByte('\n')
	}
	return bdr.String()
}

func (s *tokenSource) JSON() string {
	type jsonToken struct {
		Type  TokenType `json:"type"`
		Value string    `json:"value"`
	}
	jts := make([]jsonToken, len(s.tokens))
	for i, tkn := range s.tokens {
		jts[i] = jsonToken{
			Type:  tkn.Type(),
			Value: tkn.Value(),
		}
	}
	b, err := json.Marshal(jts)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func WrapTokenSource(s piter.Source) (*tokenSource, bool) {
	ts, ok := s.(*tokenSource)
	return ts, ok
}

// Analyze converts rune source to token source.
func Analyze(src piter.Source) piter.Source {
	var iter = piter.New(src)
	var ts = &tokenSource{
		tokens: make([]Token, 0),
		cur:    0,
	}

	for iter.HasNext() {
		var (
			r  rune
			ok bool
		)
		if r, ok = iter.Next().(rune); !ok {
			panic("lexer.Analyze: invalid source")
		}

		if r == ' ' || r == '\n' {
			continue
		}

		if r == '/' {
			if iter.Next() == '/' {
				for iter.Next() != '\n' {
				}
				iter.PutBack()
				continue
			} else {
				iter.PutBack()
			}
		}

		if r == '{' || r == '}' || r == '(' || r == ')' {
			tkn := makeBracketToken(r)
			ts.tokens = append(ts.tokens, tkn)
		}

		if isLetter(r) {
			iter.PutBack()
			tkn := makeIdentiferOrKeywordToken(iter)
			ts.tokens = append(ts.tokens, tkn)
		}

		if isNumber(r) {
			iter.PutBack()
			tkn := makeNumberToken(iter)
			ts.tokens = append(ts.tokens, tkn)
		}

		if isOperator(r) {
			iter.PutBack()
			tkn := makeOperatorToken(iter)
			ts.tokens = append(ts.tokens, tkn)
		}
	}
	return ts
}
