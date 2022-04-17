package main

import (
	"testing"
)

func TestParmVarDecl1(t *testing.T) {
	var pvd ParmVarDecl
	pvd.name = "f1"
	pvd.type1 = createTypeClang()
	pvd.type1.qtype.typestr = "int"

	res := pvd.t2go()
	expected := "f1 int"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
