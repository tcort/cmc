# cmc

`cmc` is a compiler for the C Minus programming language, a small subset of the `C` programming language
which is described in "Compiler Construction: Principles and Practice" by Kenneth C. Louden.

## Motivation

In my spare time, I have an interest in compilers and interpreters for toy languages. I've developed a
[TinyBASIC](https://github.com/tcort/tcbasic) interpreter, a [Micro Manual LISP](https://github.com/tcort/edgar) interpreter,
and an assembler and virtual machine for the [Little Man Computer](https://github.com/tcort/lmc). However, I've
never developed a full compiler, so this is the project to do it.

I was first introduced to the C Minus programming language in a university compilers course. It was
a one semester undergraduate course, so we weren't able to develop a full compiler, only a scanner, parser,
and semantic checker. Since then, I've always wanted to do a full C Minus compiler.

I really want to gain more experience with the [go](https://golang.org/) programming language, so it was
chosen as the implementation language for no particular reason other than me wanting to experiment.

## Current Status

The project is still in its infancy. I've got a working scanner, but there is still a lot of work left.

## Requirements

Compiling this software just requires a [go](https://golang.org/) toolchain. There are no external dependencies.

## Important Files

* `LICENSE` - the license which applies to all files within this repository
* `LANG.md` - language specification for the C Minus programming language
* `HACKING.md` - an English language map of the code base
