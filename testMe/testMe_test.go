package testme

import (
	"testing"
)

func TestF1(t *testing.T) {
	if f1(0) != 0 {
		t.Error("f1(0) != 0")
	}
	if f1(1) != 1 {
		t.Error("f1(1) != 1")
	}
	if f1(3) != 2 {
		t.Error("f1(3) != 2")
	}
}

func TestF2(t *testing.T) {
	if f2(0) != 0 {
		t.Error("f2(0) != 0")
	}
	if f2(1) != 1 {
		t.Error("f2(1) != 1")
	}
	if f2(3) != 2 {
		t.Error("f2(3) != 2")
	}
}

func TestS1(t *testing.T) {
	if s1("123456789") != 9 {
		t.Error("s1(\"123456789\") != 9")
	}
	if s1("") != 0 {
		t.Error("s1(\"\") != 0")
	}
}

func TestS2(t *testing.T) {
	if s2("123456789") != 9 {
		t.Error("s2(\"123456789\") != 9")
	}
	if s2("") != 0 {
		t.Error("s2(\"\") != 0")
	}
}
