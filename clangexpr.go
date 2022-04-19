package main

import (
	"fmt"
	"log"
)

type ExprParam struct {
	NodeParam
	valueCategory string
}

type ParenExpr struct {
	ExprParam
}

type CStyleCastExpr struct {
	ExprParam
	castKind string
}

type ImplicitCastExpr struct {
	ExprParam
	castKind             string
	isPartOfExplicitCast bool
}
type DeclRefExpr struct {
	ExprParam
	referencedDecl  ClangNode
	nonOdrUseReason string
}

// Represents a function call (C99 6.5.2.2, C++ [expr.call]).
// https://clang.llvm.org/doxygen/classclang_1_1CallExpr.html
type CallExpr struct {
	ExprParam
}

type ArraySubscriptExpr struct {
	ExprParam
}

// eg: &(qoi_desc){.width = w, .height = h}
type CompoundLiteralExpr struct {
	ExprParam
}

type InitListExpr struct {
	ExprParam
}

type MemberExpr struct {
	ExprParam
	name                 string
	isArrow              bool
	referencedMemberDecl string
}

type UnaryExprOrTypeTraitExpr struct {
	ExprParam
	name string
}

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
		case "isPartOfExplicitCast":
			inst.isPartOfExplicitCast = v.(bool)
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
		case "nonOdrUseReason":
			inst.nonOdrUseReason = v.(string)
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
		case "castKind":
			inst.castKind = v.(string)
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

func convertCompoundLiteralExpr(content interface{}) *CompoundLiteralExpr {
	var inst CompoundLiteralExpr
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
			fmt.Printf("[DBG][CompoundLiteralExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertInitListExpr(content interface{}) *InitListExpr {
	var inst InitListExpr
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
			fmt.Printf("[DBG][InitListExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertMemberExpr(content interface{}) *MemberExpr {
	var inst MemberExpr
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
		case "name":
			inst.name = v.(string)
		case "isArrow":
			inst.isArrow = v.(bool)
		case "referencedMemberDecl":
			inst.referencedMemberDecl = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][MemberExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertUnaryExprOrTypeTraitExpr(content interface{}) *UnaryExprOrTypeTraitExpr {
	var inst UnaryExprOrTypeTraitExpr
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
		case "name":
			inst.name = v.(string)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][UnaryExprOrTypeTraitExpr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
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
		s += nd.t2go() + CommaStr
	}

	s = RemoveLastSubStr(s, CommaStr)
	s += ")"
	return s
}

func (p *ParenExpr) t2go() string {
	switch len(p.inner) {
	case 0:
		return ""
	case 1:
		{
			// Only one child?
			return p.inner[0].t2go()
		}
	default:
		{
			fmt.Printf("[DBG][ParenExpr]Format error: %+v!\n", p.inner)
			return ""
		}
	}
}

func (p *CStyleCastExpr) t2go() string {
	switch len(p.inner) {
	case 0:
		return ""
	case 1:
		{
			// Only one child?
			return p.inner[0].t2go()
		}
	default:
		{
			fmt.Printf("[DBG][CStyleCastExpr]Format error: %+v!\n", p.inner)
			return ""
		}
	}
}

func (p *ImplicitCastExpr) t2go() string {
	switch len(p.inner) {
	case 0:
		return ""
	case 1:
		{
			// Only one child?
			return p.inner[0].t2go()
		}
	default:
		{
			fmt.Printf("[DBG][ImplicitCastExpr]Format error: %+v!\n", p.inner)
			return ""
		}
	}
}

func (p *DeclRefExpr) t2go() string {
	if p.referencedDecl == nil {
		log.Printf("[DBG][DeclRefExpr]Empty ref decl in %+v!\n", p)
		return ""
	}
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
	num := len(p.inner)
	if num != 2 {
		log.Printf("[DBG][ArraySubscriptExpr]Format error: %+v\n", p)
		return ""
	}

	// the first is func name
	s := p.inner[0].t2go() + "[" + p.inner[1].t2go() + "]"
	return s
}

func (p *CompoundLiteralExpr) t2go() string {
	num := len(p.inner)
	if num != 1 {
		log.Printf("[DBG][CompoundLiteralExpr]Format error: %+v\n", p)
		return ""
	}

	// ?
	s := "&" + p.type1.t2go() + LeftBraceStr + p.inner[0].t2go() + RightBraceStr
	return s
}

func (p *InitListExpr) t2go() string {
	num := len(p.inner)
	if num == 0 {
		log.Printf("[DBG][InitListExpr]Format error: %+v\n", p)
		return ""
	}

	s := ""
	for _, nd := range p.inner {
		s += nd.t2go() + CommaStr
	}

	s = RemoveLastSubStr(s, CommaStr)
	return s
}

func (p *MemberExpr) t2go() string {
	num := len(p.inner)
	if num != 1 {
		log.Printf("[DBG][MemberExpr]Format error: %+v\n", p)
		return ""
	}

	// the first is struct name
	s := p.inner[0].t2go() + "." + p.name
	return s
}

func (p *UnaryExprOrTypeTraitExpr) t2go() string {
	num := len(p.inner)
	if num != 1 {
		log.Printf("[DBG][UnaryExprOrTypeTraitExpr]Format error: %+v\n", p)
		return ""
	}

	// the first is struct name
	s := p.name + "(" + p.inner[0].t2go() + ")"
	return s
}
