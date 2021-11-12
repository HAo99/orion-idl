package lexer

import "regexp"

var (
	_NumberPattern   = regexp.MustCompile("^[0-9]$")
	_LetterPattern   = regexp.MustCompile("^[a-zA-Z]$")
	_LiteralPattern  = regexp.MustCompile("^[a-zA-Z0-9_]$")
	_OperatorPattern = regexp.MustCompile(`^[=\>\-]$`)
)

// r2str converts rune to a string.
func r2str(r rune) string {
	return string([]rune{r})
}

// isNumber returns true if r is a number (0~9).
func isNumber(r rune) bool {
	return _NumberPattern.MatchString(r2str(r))
}

// isLetter returns true if r is a letter (a~z and A~Z).
func isLetter(r rune) bool {
	return _LetterPattern.MatchString(r2str(r))
}

// isLiteral returns true if r is literal (0~9, a~z, A~Z and _).
func isLiteral(r rune) bool {
	return _LiteralPattern.MatchString(r2str(r))
}

// isOperator returns true if r is operator.
func isOperator(r rune) bool {
	// return r == '-' || r == '=' || r == '>'
	return _OperatorPattern.MatchString(r2str(r))
}

var (
	_Keyword = []string{
		"fn",
		"interface",
		"struct",
	}
	_FundamentalType = []string{
		"i8", "i16", "i32", "i64",
		"u8", "u16", "u32", "u64",
		"f32", "f64", "string",
	}
)

func isKeyword(s string) bool {
	for _, kw := range _Keyword {
		if s == kw {
			return true
		}
	}
	return false
}

func isFundamentalType(s string) bool {
	for _, ft := range _FundamentalType {
		if s == ft {
			return true
		}
	}
	return false
}
