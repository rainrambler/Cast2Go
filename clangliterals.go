package main

import (
	"fmt"
)

type LiteralParam struct {
	NodeParam
	valueCategory string
}

type IntegerLiteral struct {
	LiteralParam
	value1 string
}

type CharacterLiteral struct {
	LiteralParam
	value1 float64 // TODO switch value type
}

type StringLiteral struct {
	LiteralParam
	value1 string
}

func (p *IntegerLiteral) t2go() string {
	return p.value1
}

func (p *CharacterLiteral) t2go() string {
	return ""
}

func (p *StringLiteral) t2go() string {
	return p.value1
}

func convertIntegerLiteral(content interface{}) *IntegerLiteral {
	var inst IntegerLiteral
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
		case "value":
			inst.value1 = v.(string)
		default:
			fmt.Printf("[DBG][IntegerLiteral]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][IntegerLiteral]%+v\n", inst)
	return &inst
}

func convertCharacterLiteral(content interface{}) *CharacterLiteral {
	var inst CharacterLiteral
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
		case "value":
			inst.value1 = v.(float64) // ?
		default:
			fmt.Printf("[DBG][CharacterLiteral]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][CharacterLiteral]%+v\n", inst)
	return &inst
}

func convertStringLiteral(content interface{}) *StringLiteral {
	var inst StringLiteral
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
		case "value":
			inst.value1 = v.(string) // ?
		default:
			fmt.Printf("[DBG][StringLiteral]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][StringLiteral]%+v\n", inst)
	return &inst
}
