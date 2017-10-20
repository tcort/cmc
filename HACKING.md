# cmc internals

## `cmc.go`

This is the main driver program. It ingests the source file and runs the code through
all stages of the compiler.

## scanner

As is the first stage of the compiler pipeline, this package tokenizes the source text.
The token structure is defined in `scanner/token.go`
