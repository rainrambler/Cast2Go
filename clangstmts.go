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
		case *BinaryOperator, *UnaryOperator, *CompoundAssignOperator:
			s += nd.t2go() + SemicolonStr
		default:
			s = RemoveLastSubStr(s, SemicolonStr)
			s += LeftBraceStr + EnterStr
			s += nd.t2go()
		}
	}

	s += RightBraceStr
	return s
}

func (p *IfStmt) t2go() string {
	s := "if "

	inited := false
	size := len(p.inner)
	for i := 0; i < size-1; i++ {
		nd := p.inner[i]
		switch nd.(type) {
		case *BinaryOperator, *UnaryOperator:
			s += nd.t2go()
		default:
			if !inited {
				s += LeftBraceStr + EnterStr
				inited = true
			}
			s += nd.t2go()
		}
	}

	if p.hasElse {
		if inited {
			//if !strings.HasSuffix(s, RightBraceStr) {
			s += EnterStr + RightBraceStr
			//}
			s += " else " + LeftBraceStr + EnterStr
			s += p.inner[size-1].t2go() // The last should be "else" stmt?
		} else {
			s += LeftBraceStr + EnterStr
			inited = true
		}
	} else {
		if !inited {
			s += LeftBraceStr + EnterStr
		}
		s += p.inner[size-1].t2go()
	}

	s += EnterStr + RightBraceStr
	return s
}

func (p *DeclStmt) t2go() string {
	s := ""
	for _, nd := range p.inner {
		s += nd.t2go() + EnterStr
	}
	return RemoveLastSubStr(s, EnterStr)
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
	return RemoveLastSubStr(s, EnterStr)
}
