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

type BinaryOperator struct {
	Operator
}

type UnaryOperator struct {
	Operator
	isPostfix   bool
	canOverflow bool
}

type CompoundAssignOperator struct {
	Operator
	computeResultType *TypeClang
	computeLHSType    *TypeClang
}

type Operator struct {
	NodeParam
	valueCategory string
	opcode        string
}

func (p *BinaryOperator) t2go() string {
	return ""
}

func (p *UnaryOperator) t2go() string {
	return ""
}

func (p *CompoundAssignOperator) t2go() string {
	return ""
}
