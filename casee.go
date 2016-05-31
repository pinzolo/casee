package casee

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	if IsSnakeCase(s) {
		return s
	}
	if IsUpperSnakeCase(s) {
		return strings.ToLower(s)
	}

	return s
}

func IsSnakeCase(s string) bool {
	if strings.Contains(s, "_") {
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

func ToCamelCase(s string) string {
	return s
}

func IsCamelCase(s string) bool {
	return true
}

func ToPascalCase(s string) string {
	if IsPascalCase(s) {
		return s
	}

	return s
}

func IsPascalCase(s string) bool {
	return true
}

func ToUpperSnakeCase(s string) string {
	if IsUpperSnakeCase(s) {
		return s
	}
	if IsSnakeCase(s) {
		return strings.ToUpper(s)
	}

	return s
}

func IsUpperSnakeCase(s string) bool {
	if strings.Contains(s, "_") {
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
