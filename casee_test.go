package casee

import (
	"testing"
)

type boolCase struct {
	target   string
	expected bool
}

func TestIsSnakeCase(t *testing.T) {
	testCases := []boolCase{
		{"foo_bar", true},
		{"foo1_bar2", true},
		{"foo_bar_1", true},
		{"foo_bar_1", true},
		{"1_foo_bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"1foobar", true},

		{"FOO_BAR", false},
		{"fooBar", false},
		{"FooBar", false},
		{"FOOBAR", false},
		{"foo_@bar", false},
		{"foo-bar", false},
		{"テスト", false},
		{"テスト_テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsSnakeCase(tc.target); actual != tc.expected {
			t.Errorf("IsSnakeCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestIsUpperSnakeCase(t *testing.T) {
	testCases := []boolCase{
		{"FOO_BAR", true},
		{"FOO_BAR_0", true},
		{"FOO_0_BAR", true},
		{"0_FOO_BAR", true},
		{"FOOBAR", true},
		{"FOOBAR1", true},
		{"FOO1BAR", true},
		{"1FOOBAR", true},

		{"foo_bar", false},
		{"fooBar", false},
		{"FooBar", false},
		{"foobar", false},
		{"FOO_@BAR", false},
		{"FOO-BAR", false},
		{"テスト", false},
		{"テスト_テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsUpperSnakeCase(tc.target); actual != tc.expected {
			t.Errorf("IsUpperSnakeCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}
