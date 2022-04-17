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
type CallExpr struct {
	ExprParam
}
type ArraySubscriptExpr struct{}

type ExprParam struct {
	NodeParam
	valueCategory string
}
