package scanner

type TokenType string

const (
	// Keywords
	T_ELSE   = TokenType("T_ELSE")
	T_IF     = TokenType("T_IF")
	T_INT    = TokenType("T_INT")
	T_RETURN = TokenType("T_RETURN")
	T_VOID   = TokenType("T_VOID")
	T_WHILE  = TokenType("T_WHILE")

	// Special Symbols
	T_PLUS                     = TokenType("T_PLUS")                     // +
	T_MINUS                    = TokenType("T_MINUS")                    // -
	T_TIMES                    = TokenType("T_TIMES")                    // *
	T_DIVIDE                   = TokenType("T_DIVIDE")                   // /
	T_LESS_THAN                = TokenType("T_LESS_THAN")                // <
	T_LESS_THAN_OR_EQUAL_TO    = TokenType("T_LESS_THAN_OR_EQUAL_TO")    // <=
	T_GREATER_THAN             = TokenType("T_GREATER_THAN")             // >
	T_GREATER_THAN_OR_EQUAL_TO = TokenType("T_GREATER_THAN_OR_EQUAL_TO") // >=
	T_EQUALS                   = TokenType("T_EQUALS")                   // ==
	T_NOT_EQUALS               = TokenType("T_NOT_EQUALS")               // !=
	T_ASSIGN                   = TokenType("T_ASSIGN")                   // =
	T_END_OF_STATEMENT         = TokenType("T_END_OF_STATEMENT")         // ;
	T_LIST_SEPARATOR           = TokenType("T_LIST_SEPARATOR")           // ,
	T_OPEN_LIST                = TokenType("T_OPEN_LIST")                // (
	T_CLOSE_LIST               = TokenType("T_CLOSE_LIST")               // )
	T_OPEN_INDEX               = TokenType("T_OPEN_INDEX")               // [
	T_CLOSE_INDEX              = TokenType("T_CLOSE_INDEX")              // ]
	T_OPEN_BLOCK               = TokenType("T_OPEN_BLOCK")               // {
	T_CLOSE_BLOCK              = TokenType("T_CLOSE_BLOCK")              // }

	// Other tokens (ID and NUM)
	T_IDENTIFIER = TokenType("T_IDENTIFIER")
	T_NUMBER     = TokenType("T_NUMBER")
)

type Token struct {
	Type       TokenType
	Text       string
	LineNumber int
}
