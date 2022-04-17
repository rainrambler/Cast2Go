package main

import (
	"testing"
)

func TestRemoveLastSubStr1(t *testing.T) {
	s := "abcd;"

	res := RemoveLastSubStr(s, ";")
	expected := "abcd"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestRemoveLastSubStr2(t *testing.T) {
	s := "abcd"

	res := RemoveLastSubStr(s, ";")
	expected := "abcd"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestRemoveLastSubStr3(t *testing.T) {
	s := "abcd"

	res := RemoveLastSubStr(s, "cd")
	expected := "ab"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestRemoveLastSubStrEmpty(t *testing.T) {
	s := ""

	res := RemoveLastSubStr(s, "")
	expected := ""

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
