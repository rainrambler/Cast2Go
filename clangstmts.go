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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][DeclStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][DeclStmt]%+v\n", inst)
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ForStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][ForStmt]%+v\n", inst)
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][IfStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][IfStmt]%+v\n", inst)
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ReturnStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][ReturnStmt]%+v\n", inst)
	return &inst
}

// https://clang.llvm.org/doxygen/classclang_1_1Stmt.html
type Stmt struct {
}

type CompoundStmt struct {
	kind   string
	id     string
	range1 *SourceRange // "range" is a keyword
	inner  []ClangNode
}

type ReturnStmt struct {
	kind   string
	id     string
	range1 *SourceRange // "range" is a keyword
	inner  []ClangNode
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
