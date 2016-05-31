// Provide convert functions (string => snake_case, UPPER_SNAKE_CASE, camelCase, PascalCase).
package casee

import (
	"strings"
	"unicode"
)

// Convert argument to snake_case style string.
// If argument is empty, return itself.
func ToSnakeCase(s string) string {
	if IsSnakeCase(s) {
		return s
	}
	if IsUpperSnakeCase(s) {
		return strings.ToLower(s)
	}

	return s
}

// If argument is snake_case style string, return true.
func IsSnakeCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByLowerAndDigit(field) {
				return false
			}
		}
		return true
	} else {
		return isMadeByLowerAndDigit(s)
	}
}

// Convert argument to camelCase style string
// If argument is empty, return itself
func ToCamelCase(s string) string {
	return s
}

// If argument is camelCase style string, return true.
// If first character is digit, always returns false
func IsCamelCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else {
		return isMadeByAlphanumeric(s) && isFirstRuneLower(s)
	}
}

// Convert argument to PascalCase style string
// If argument is empty, return itself
func ToPascalCase(s string) string {
	if IsPascalCase(s) {
		return s
	}

	return s
}

// If argument is PascalCase style string, return true.
// If first character is digit, always returns false
func IsPascalCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else {
		return isMadeByAlphanumeric(s) && isFirstRuneUpper(s)
	}
}

// Convert argument to UPPER_SNAKE_CASE style string
// If argument is empty, return itself
func ToUpperSnakeCase(s string) string {
	if IsUpperSnakeCase(s) {
		return s
	}
	if IsSnakeCase(s) {
		return strings.ToUpper(s)
	}

	return s
}

// If argument is UPPER_SNAKE_CASE style string, return true.
func IsUpperSnakeCase(s string) bool {
	if isFirstRuneDigit(s) {
		return false
	} else if strings.Contains(s, "_") {
		fields := strings.Split(s, "_")
		for _, field := range fields {
			if !isMadeByUpperAndDigit(field) {
				return false
			}
		}
		return true
	} else {
		return isMadeByUpperAndDigit(s)
	}
}

func isMadeByLowerAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByUpperAndDigit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isMadeByAlphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && !unicode.IsLower(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isFirstRuneUpper(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsUpper(getRuneAt(s, 0))
}

func isFirstRuneLower(s string) bool {
	if len(s) == 0 {
		return false
	}

	return unicode.IsLower(getRuneAt(s, 0))
}

func isFirstRuneDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.IsDigit(getRuneAt(s, 0))
}

func getRuneAt(s string, i int) rune {
	if len(s) == 0 {
		return 0
	}

	rs := []rune(s)
	return rs[0]
}
