package casee

import (
	"testing"
)

type boolExpectedCase struct {
	target   string
	expected bool
}

type stringExpectedCase struct {
	target   string
	expected string
}

func TestIsSnakeCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo_bar", true},
		{"foo1_bar2", true},
		{"foo_bar_1", true},
		{"foo_bar_1", true},
		{"111_foo_bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"111foobar", true},

		{"", false},
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

func TestToSnakeCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foo_bar"},
		{"foo-bar", "foo_bar"},
		{"foo-bar_baz", "foo_bar_baz"},
		{"foo--bar__baz", "foo_bar_baz"},
		{"fooBar", "foo_bar"},
		{"FooBar", "foo_bar"},
		{"foo bar", "foo_bar"},
		{"   foo   bar   ", "foo_bar"},
		{"fooBar111", "foo_bar_111"},
		{"111FooBar", "111_foo_bar"},
		{"foo-111-Bar", "foo_111_bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToSnakeCase(tc.target); actual != tc.expected {
			t.Errorf("ToSnakeCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsChainCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foo-bar", true},
		{"foo1-bar2", true},
		{"foo-bar-1", true},
		{"foo-bar-1", true},
		{"111-foo-bar", true},
		{"foobar", true},
		{"foobar1", true},
		{"foo1bar", true},
		{"1foobar", true},

		{"", false},
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

func TestToChainCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foo-bar"},
		{"foo-bar", "foo-bar"},
		{"foo-bar_baz", "foo-bar-baz"},
		{"foo--bar__baz", "foo-bar-baz"},
		{"fooBar", "foo-bar"},
		{"FooBar", "foo-bar"},
		{"foo bar", "foo-bar"},
		{"   foo   bar   ", "foo-bar"},
		{"fooBar111", "foo-bar-111"},
		{"111FooBar", "111-foo-bar"},
		{"foo-111-Bar", "foo-111-bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToChainCase(tc.target); actual != tc.expected {
			t.Errorf("ToChainCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
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

func TestToCamelCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "fooBar"},
		{"foo-bar", "fooBar"},
		{"foo-bar_baz", "fooBarBaz"},
		{"foo--bar__baz", "fooBarBaz"},
		{"fooBar", "fooBar"},
		{"FooBar", "fooBar"},
		{"foo bar", "fooBar"},
		{"   foo   bar   ", "fooBar"},
		{"fooBar111", "fooBar111"},
		{"111FooBar", "111FooBar"},
		{"foo-111-Bar", "foo111Bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToCamelCase(tc.target); actual != tc.expected {
			t.Errorf("ToCamelCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
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

func TestToPascalCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "FooBar"},
		{"foo-bar", "FooBar"},
		{"foo-bar_baz", "FooBarBaz"},
		{"foo--bar__baz", "FooBarBaz"},
		{"fooBar", "FooBar"},
		{"FooBar", "FooBar"},
		{"foo bar", "FooBar"},
		{"   foo   bar   ", "FooBar"},
		{"fooBar111", "FooBar111"},
		{"111FooBar", "111FooBar"},
		{"foo-111-Bar", "Foo111Bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToPascalCase(tc.target); actual != tc.expected {
			t.Errorf("ToPascalCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestToFlatCase(t *testing.T) {
	testCases := []stringExpectedCase{
		{"foo_bar", "foobar"},
		{"foo-bar", "foobar"},
		{"foo-bar_baz", "foobarbaz"},
		{"foo--bar__baz", "foobarbaz"},
		{"fooBar", "foobar"},
		{"FooBar", "foobar"},
		{"foo bar", "foobar"},
		{"   foo   bar   ", "foobar"},
		{"fooBar111", "foobar111"},
		{"111FooBar", "111foobar"},
		{"foo-111-Bar", "foo111bar"},
		{"", ""},
	}

	for _, tc := range testCases {
		if actual := ToFlatCase(tc.target); actual != tc.expected {
			t.Errorf("ToFlatCase(%s) returns %s, but expected %s", tc.target, actual, tc.expected)
		}
	}
}

func TestIsFlatCase(t *testing.T) {
	testCases := []boolExpectedCase{
		{"foobar", true},
		{"foo1bar", true},

		{"", false},
		{"1foobar", false},
		{"foo@bar", false},
		{"foo_bar", false},
		{"FOO_BAR", false},
		{"FooBar", false},
		{"fooBar", false},
		{"foo_bar", false},
	}

	for _, tc := range testCases {
		if actual := IsFlatCase(tc.target); actual != tc.expected {
			t.Errorf("IsFlatCase(%s) returns %t, but expected %t", tc.target, actual, tc.expected)
		}
	}
}

