package main

import (
	"fmt"
)

func convertCompoundStmt(content interface{}) *CompoundStmt {
	var inst CompoundStmt
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][CompoundStmt]%+v\n", inst)
	return &inst
}

func convertDeclStmt(content interface{}) *DeclStmt {
	var inst DeclStmt
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][DeclStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertForStmt(content interface{}) *ForStmt {
	var inst ForStmt
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ForStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertIfStmt(content interface{}) *IfStmt {
	var inst IfStmt
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "hasElse":
			inst.hasElse = v.(bool)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][IfStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertReturnStmt(content interface{}) *ReturnStmt {
	var inst ReturnStmt
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ReturnStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

// https://clang.llvm.org/doxygen/classclang_1_1Stmt.html
type Stmt struct {
}

type CompoundStmt struct {
	NodeParam
}

type ReturnStmt struct {
	NodeParam
}

type DeclStmt struct {
	NodeParam
}

type ForStmt struct {
	NodeParam
}

type IfStmt struct {
	NodeParam
	hasElse bool
}

func (p *ForStmt) t2go() string {
	s := "for "

	for _, nd := range p.inner {
		switch nd.(type) {
		case *BinaryOperator, *UnaryOperator:
			s += nd.t2go() + ";"
		default:
			s = RemoveLastSubStr(s, ";")
			s += "{\n"
			s += nd.t2go()
		}
	}

	s += "}"
	return s
}

func (p *IfStmt) t2go() string {
	s := "if "

	inited := false
	for _, nd := range p.inner {
		switch nd.(type) {
		case *BinaryOperator, *UnaryOperator:
			s += nd.t2go()
		default:
			if !inited {
				s += "{\n"
				inited = true
			}
			s += nd.t2go()
		}
	}

	s += "\n}"
	return s
}

func (p *DeclStmt) t2go() string {
	s := ""
	for _, nd := range p.inner {
		s += nd.t2go()
	}
	return s
}

func (p *ReturnStmt) t2go() string {
	if len(p.inner) == 0 {
		return "return"
	}

	return "return " + p.inner[0].t2go() // multi inner item?
}

func (p *CompoundStmt) t2go() string {
	s := ""
	for _, nd := range p.inner {
		s += nd.t2go() + EnterStr
	}
	return s
}
