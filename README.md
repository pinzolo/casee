# casee

[![Build Status](https://travis-ci.org/pinzolo/casee.png)](http://travis-ci.org/pinzolo/casee)

Golang liibrary for case convertion of string.

## Usage

1. `go get -d github.com/pinzolo/casee`.
2. Add `github.com/pinzolo/casee` to `import` section in your go file.

## Functions

### Convert functions

* `ToSnakeCase`  
  Convert to snake_case style string.
* `ToChainCase`  
  Convert to chain-case style string.
* `ToCamelCase`  
  Convert to camelCase style string.
* `ToPascalCase`  
  Convert to PascalCase style string.

### Check functions

* `IsSnakeCase`  
  Check argument string is snake_case.
* `IsChainCase`  
  Check argument string is chain-case.
* `IsCamelCase`  
  Check argument string is camelCase.
* `IsPascalCase`  
  Check argument string is PascalCase.

If first character is digit, `IsCamelCase` and `IsPascalCase` always returns false.  
Because cannot judge camelCase or PascalCase.

