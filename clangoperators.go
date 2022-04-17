package main

import (
	"fmt"
)

func convertBinaryOperator(content interface{}) *BinaryOperator {
	var inst BinaryOperator
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
		case "type":
			inst.type1 = convertTypeClang(v)
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "opcode":
			inst.opcode = v.(string)
		default:
			fmt.Printf("[DBG][BinaryOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertUnaryOperator(content interface{}) *UnaryOperator {
	var inst UnaryOperator
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
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "opcode":
			inst.opcode = v.(string)
		case "isPostfix":
			inst.isPostfix = v.(bool)
		case "canOverflow":
			inst.canOverflow = v.(bool)
		case "type":
			inst.type1 = convertTypeClang(v)
		default:
			fmt.Printf("[DBG][UnaryOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertCompoundAssignOperator(content interface{}) *CompoundAssignOperator {
	var inst CompoundAssignOperator
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
		case "valueCategory":
			inst.valueCategory = v.(string)
		case "opcode":
			inst.opcode = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "computeLHSType":
			inst.computeLHSType = convertTypeClang(v)
		case "computeResultType":
			inst.computeResultType = convertTypeClang(v)
		default:
			fmt.Printf("[DBG][CompoundAssignOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][CompoundAssignOperator]%+v\n", inst)
	return &inst
}

// A builtin binary operation expression
// https://clang.llvm.org/doxygen/classclang_1_1BinaryOperator.html
type BinaryOperator struct {
	Operator
}

func (p *BinaryOperator) t2go() string {
	if len(p.inner) != 2 {
		return p.opcode
	}

	left := p.inner[0]
	right := p.inner[1]
	return left.t2go() + " " + p.opcode + " " + right.t2go()
}

type UnaryOperator struct {
	Operator
	isPostfix   bool
	canOverflow bool
}

// For compound assignments
// https://clang.llvm.org/doxygen/classclang_1_1CompoundAssignOperator.html
type CompoundAssignOperator struct {
	Operator
	computeResultType *TypeClang
	computeLHSType    *TypeClang
}

func (p *CompoundAssignOperator) t2go() string {
	if len(p.inner) != 2 {
		return p.opcode
	}

	left := p.inner[0]
	right := p.inner[1]

	// TODO valueCategory
	return left.t2go() + " " + p.opcode + " " + right.t2go()
}

type Operator struct {
	NodeParam
	valueCategory string
	opcode        string
}

func (p *UnaryOperator) t2go() string {
	if len(p.inner) != 1 {
		return p.opcode
	}

	if p.isPostfix {
		return p.inner[0].t2go() + p.opcode
	} else {
		return p.opcode + p.inner[0].t2go()
	}
}
