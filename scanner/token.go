package scanner

type TokenType string

const (
	// Keywords
	TokenElse   = TokenType("TokenElse")
	TokenIf     = TokenType("TokenIf")
	TokenInt    = TokenType("TokenInt")
	TokenReturn = TokenType("TokenReturn")
	TokenVoid   = TokenType("TokenVoid")
	TokenWhile  = TokenType("TokenWhile")

	// Special Symbols
	TokenPlus                 = TokenType("TokenPlus")                 // +
	TokenMinus                = TokenType("TokenMinus")                // -
	TokenTimes                = TokenType("TokenTimes")                // *
	TokenDivide               = TokenType("TokenDivide")               // /
	TokenLessThan             = TokenType("TokenLessThan")             // <
	TokenLessThanOrEqualTo    = TokenType("TokenLessThanOrEqualTo")    // <=
	TokenGreaterThan          = TokenType("TokenGreaterThan")          // >
	TokenGreaterThanOrEqualTo = TokenType("TokenGreaterThanOrEqualTo") // >=
	TokenEquals               = TokenType("TokenEquals")               // ==
	TokenNotEquals            = TokenType("TokenNotEquals")            // !=
	TokenAssign               = TokenType("TokenAssign")               // =
	TokenEndOfStatement       = TokenType("TokenEndOfStatement")       // ;
	TokenListSeparator        = TokenType("TokenListSeparator")        // ,
	TokenOpenList             = TokenType("TokenOpenList")             // (
	TokenCloseList            = TokenType("TokenCloseList")            // )
	TokenOpenIndex            = TokenType("TokenOpenIndex")            // [
	TokenCloseIndex           = TokenType("TokenCloseIndex")           // ]
	TokenOpenBlock            = TokenType("TokenOpenBlock")            // {
	TokenCloseBlock           = TokenType("TokenCloseBlock")           // }

	// Other tokens (ID and NUM)
	TokenIdentifier = TokenType("TokenIdentifier")
	TokenNumber     = TokenType("TokenNumber")
)

type Token struct {
	Type       TokenType
	Text       string
	LineNumber int
}
