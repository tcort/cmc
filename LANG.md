# C Minus Language Reference

This is a general language reference for the C Minus programming language.

## Lexical Conventions

The following outlines the various lexical conventions of C Minus source files.
C Minus source files are treated as ASCII text.

### Keywords

All keywords are reserved words and must be written exactly as they appear below.
For example, `returned` surrounded by whitespace would be recognized as an identifier.
Keywords are case sensitive. `Return` surrounded by whitespace would also be recognized
as an identifier.

#### `if`

Conditional execution construct.

Example:

```lang=C
int x = input();
if (x < 10) {
    output(1);
}
```

#### `else`

Used with `if` for conditional execution.

Example:

```lang=C
int x = input();
if (x == 10) {
    output(10);
} else {
    output(0);
}
```

#### `while`

Looping construct.

Example:

```lang=C
int x = 0;
while (x < 10) {
    output(x);
    x = x + 1;
}
```

#### `return`

Returns control from a function call.

### Types

There are two type specifiers: `int` and `void`. No other types are provided.
As with keywords, the types must be written exactly as they appear below.

#### `int`

Denotes an integer. The width of the variable is machine dependent.

#### `void`

Denotes the lack of a value. For example, a `void` function does not return a value.

### Relational Operators

Relational operators. These evaluate the relationship between the expressions on either side of them.
If the relation is true, they evaluate to `1`, otherwise `0`.

#### `<=`

The left hand side is less than or equal to the right hand side.

#### `<`

The left hand side is less than to the right hand side.

#### `>`

The left hand side is greater than the right hand side.

#### `>=`

The left hand side is greater than or equal to the right hand side.

#### `==`

The left hand side is equal to the right hand side.

#### `!=`

The left hand side is not equal to the right hand side.

### Arithmetic Operators

There are 4 arithmetic operations available. In complex expressions, multiplication and division take precedence over addition and subtraction.
There is no protection against integer overflow.

#### `+`

Integer addition.

#### `-`

Integer subtraction.

#### `*`

Integer multiplication.

#### `/`

Integer division.

### Identifiers

Identifiers consist of a case sensitive string of one or more mixed case letters. For example, `foo`, `Foo`, `fOO` and `f` are
all acceptable and distinct identifiers. There is no maximum length of an identifier but particular implementions may impose one.

### Literals

Literal values may appear in the source code.

#### Integers

Integer literals consist of a string of one or more digits. For example, `0`, `24`, `42` are all acceptable literals.

Integer literals must be positive numbers. For example, `-42` is not acceptable. As a work around, the expression
`0 - 42` may be used to represent `-42`.

Integers other than `0` may not begin with `0`. For example, `042` is not allowed.

### Comments

Comments begin with `/*` and end with `*/`. Comments may not be nested (e.g. `/* foo /* bar */ baz */` is not acceptable).
Comments may appear where whitespace may appear. They cannot be placed within tokens. For example, `if/*x*/(x ===1)` is
OK but `i/*x*/f (x==1)` is not acceptable. Comments may contain new lines.

### Whitespace

Whitespace consists of spaces, new lines (`\n`), and tabs (`\t`).

Other whitespace characters are not acceptable (e.g. `\r`, `\f`, etc)

## Semantic Conventions

There are several rules which apply to all programs.

### Return Statement

All functions must return and it must be used in accordance with the function's return type.

Either the last statement must be a `return` statement or if the last statement is an `if`/`else` statement, all branches
must `return`.

### Minimal C Minus Program

A valid source file must have at least one declaration.

### Declarations

Identifiers must be declared before use.

### Identifier usage

The use of an identifier is checked against the declaration (e.g. a variable cannot be called nor can a function be assigned).

### Duplicate names

Identifiers can only be declared once per scope. For example, two functions cannot be defined to have the same name.

### Array Subscripts

Arrays must be subscripted except when passing an array as an argument to a function.

### Array Size

Zero length arrays are not allowed.

### Variable types

Variable declarations can only be of type `int`. A `void` variable is not acceptable.

### Function Parameters

Function calls must match the function signature. The number of arguments must match the number of parameters and the types must match.

The parameters list may either be `void` or a comma separated list of variables. The types of the variables cannot be `void`.

### `void main(void)`

All programs must defined a `main` function, and it must be the last declaration in the source file.

## Runtime Behaviour

Very little about runtime behaviour is ensconced in the language definition.
There is one exception.

### Negative Array Indexes

When a negative value is used as an array index, the program halts.

## Standard Library

Two I/O functions are provided for reading and writing numbers as numeric strings.

### `int input(void)`

Accepts an integer string entered on standard input and returns it.
The input is terminated by a new line. For example, if standard
input is a text terminal, the user would type `4`, `2`, and press `[enter]`;
this function would return the integer value `42`. There is no facility for error
detection nor truncation detection.

### `void output(int x)`

Outputs an integer value as a string followed by a new line to standard output.
For example, if `output(42)` were called, `42\n` would be written to standard output.

## Quick Reference

### Tokens

* Identifiers
  * `[A-Za-z]+`
* Numbers
  * `[0-9]+`
* Keywords
  * `while`
  * `if`
  * `else`
  * `int`
  * `void`
  * `return`
* Special Characters
  * `+`
  * `-`
  * `*`
  * `/`
  * `<`
  * `<=`
  * `>`
  * `>=`
  * `==`
  * `!=`
  * `=`
  * `;`
  * `,`
  * `(`
  * `)`
  * `[`
  * `]`
  * `{`
  * `}`

### Grammar

* `Program`
  * `DeclarationList`
* `DeclarationList`
  * `DeclarationList`
  * `Declaration`
* `Declaration`
  * `VarDeclaration`
  * `FunDeclaration`
* `VarDeclaration`
  * `TypeSepcifier` //Identifier// **;**
  * `TypeSepcifier` //Identifier// **[** //Number// **]** **;**
* `TypeSpecifier`
  * **int**
  * **void**
* `FunDeclaration`
  * `TypeSpecifier` //Identifier// **(** `Params` **)** `CompoundStatement`
* `Params`
  * `ParamList`
  * **void**
* `ParamList`
  * `ParamList` **,** `Param`
* `Param`
  * `TypeSpecifier` //Identifier//
  * `TypeSpecifier` //Identifier// **[** **]**
* `CompoundStatement`
  * **{** `LocalDeclarations` `StatementList` **}**
* `LocalDeclarations`
  * `LocalDeclarations` `VarDeclaration`
  * //empty//
* `StatementList`
  * `StatementList` `Statement`
  * //empty//
* `Statement`
  * `ExpressionStatement`
  * `CompoundStatement`
  * `SelectionStatement`
  * `IterationStatement`
  * `ReturnStatement`
* `ExpressionStatement`
  * `Expression` **;**
  * **;**
* `SelectionStatement`
  * **if** **(** `Expression` **)** `Statement`
  * **if** **(** `Expression` **)** `Statement` **else** `Statement`
* `IterationStatement`
  * **while** **(** `Expression` **)** `Statement`
* `ReturnStatement`
  * **return** `Expression` **;**
  * **return** **;**
* `Expression`
  * `Var` **=** `Expression`
  * `SimpleExpression`
* `SimpleExpression`
  * `AdditiveExpression` `Relop` `AdditiveExpression`
  * `AdditiveExpression`
* `Relop`
  * **<=**
  * **<**
  * **>**
  * **>=**
  * **==**
  * **!=**
* `AdditiveExpression`
  * `AdditiveExpression` `Addop` `Term`
  * `Term`
* `Addop`
  * **+**
  * **-**
* `Term`
  * `Term` `Mulop` `Factor`
  * `Factor`
* `Mulop`
  * *****
  * **/**
* `Factor`
  * **(** `Expression` **)**
  * `Var`
  * `Call`
  * //Number//
* `Call`
  * //Identifier// **(** `Args` **)**
* `Args`
  * `ArgsList`
  * `Empty`
* `ArgsList`
  * `ArgsList` **,** `Expression`
  * `Expression`
