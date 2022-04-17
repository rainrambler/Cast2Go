package main

import (
	"fmt"
)

func convertImplicitCastExpr(content interface{}) *ImplicitCastExpr {
	var inst ImplicitCastExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "castKind":
			inst.castKind = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ImplicitCastExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertParenExpr(content interface{}) *ParenExpr {
	var inst ParenExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ParenExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertDeclRefExpr(content interface{}) *DeclRefExpr {
	var inst DeclRefExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "referencedDecl":
			inst.referencedDecl = convertNode(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][DeclRefExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertCallExpr(content interface{}) *CallExpr {
	var inst CallExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CallExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertCStyleCastExpr(content interface{}) *CStyleCastExpr {
	var inst CStyleCastExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CStyleCastExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertArraySubscriptExpr(content interface{}) *ArraySubscriptExpr {
	var inst ArraySubscriptExpr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "range":
			inst.range1 = convertSourceRange(v)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ArraySubscriptExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

type ParenExpr struct {
	ExprParam
}
type CStyleCastExpr struct {
	ExprParam
}

type ImplicitCastExpr struct {
	ExprParam
	castKind string
}
type DeclRefExpr struct {
	ExprParam
	referencedDecl ClangNode
}

// Represents a function call (C99 6.5.2.2, C++ [expr.call]).
// https://clang.llvm.org/doxygen/classclang_1_1CallExpr.html
type CallExpr struct {
	ExprParam
}

func (p *CallExpr) t2goOld() string {

	if len(p.inner) == 0 {
		return ""
	}

	s := ""

	for _, nd := range p.inner {
		s += nd.t2go()
	}

	//s += EnterStr
	return s
}

func (p *CallExpr) t2go() string {
	num := len(p.inner)
	if num == 0 {
		return ""
	}

	s := ""

	// the first is func name
	s += p.inner[0].t2go() + "("

	// the others are parameters
	for i := 1; i < num; i++ {
		nd := p.inner[i]
		s += nd.t2go() + ","
	}

	s = RemoveLastSubStr(s, ",")
	s += ")"
	return s
}

type ArraySubscriptExpr struct {
	ExprParam
}

type ExprParam struct {
	NodeParam
	valueCategory string
}

func (p *ParenExpr) t2go() string {
	if len(p.inner) != 1 {
		fmt.Printf("[DBG][CStyleCastExpr]Format error: %+v!\n", p.inner)
		return ""
	}

	// Only one child?
	return p.inner[0].t2go()
}

func (p *CStyleCastExpr) t2go() string {
	if len(p.inner) != 1 {
		fmt.Printf("[DBG][CStyleCastExpr]Format error: %+v!\n", p.inner)
		return ""
	}

	// Only one child?
	return p.inner[0].t2go()
}

func (p *ImplicitCastExpr) t2go() string {
	if len(p.inner) != 1 {
		fmt.Printf("[DBG][ImplicitCastExpr]Format error: %+v!\n", p.inner)
		return ""
	}

	// Only one child?
	return p.inner[0].t2go()
}

func (p *DeclRefExpr) t2go() string {
	switch p.referencedDecl.(type) {
	case *ParmVarDecl:
		decl := p.referencedDecl.(*ParmVarDecl)
		return decl.name
	case *VarDecl:
		decl := p.referencedDecl.(*VarDecl)
		return decl.name
	case *FunctionDecl:
		decl := p.referencedDecl.(*FunctionDecl)
		return decl.name
	default:
		return p.referencedDecl.t2go()
	}
}

func (p *ArraySubscriptExpr) t2go() string {
	return ""
}

func get_sum(num int) {
	var i int
	var sum int
	for i = 0; i < num; i++ {
		sum += i
	}

}
