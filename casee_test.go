package casee

import (
	"testing"
)

type boolExpectedCase struct {
	target   string
	expected bool
}

func TestIsSnakeCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo_bar", true},
		{"foo1_bar2", true},
		{"foo_bar_1", true},
		{"foo_bar_1", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},

		{"", false},
		{"1_foo_bar", false},
		{"1foobar", false},
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

func TestIsChainCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo-bar", true},
		{"foo1-bar2", true},
		{"foo-bar-1", true},
		{"foo-bar-1", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},

		{"", false},
		{"1-foo-bar", false},
		{"1foobar", false},
		{"FOO-BAR", false},
		{"fooBar", false},
		{"FooBar", false},
		{"FOOBAR", false},
		{"foo-@bar", false},
		{"foo_bar", false},
		{"テスト", false},
		{"テスト-テスト", false},
	}

	for _, tc := range testCases {
		if actual := IsChainCase(tc.target); actual != tc.expected {
			t.Errorf("IsChainCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestIsCamelCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"fooBar", true},
		{"fooBar1", true},
		{"foo1Bar", true},

		{"", false},
		{"1FooBar", false},
		{"1fooBar", false},
		{"foo@Bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"FooBar", false},
		{"Foo@Bar", false},
	}

	for _, tc := range testCases {
		if actual := IsCamelCase(tc.target); actual != tc.expected {
			t.Errorf("IsCamelCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

func TestIsPascalCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"FooBar", true},
		{"FooBar1", true},
		{"Foo1Bar", true},

		{"", false},
		{"1FooBar", false},
		{"1fooBar", false},
		{"Foo@Bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"fooBar", false},
		{"foo@Bar", false},
	}

	for _, tc := range testCases {
		if actual := IsPascalCase(tc.target); actual != tc.expected {
			t.Errorf("IsPascalCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}
