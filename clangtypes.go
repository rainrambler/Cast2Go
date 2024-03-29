package main

import (
	"fmt"
)

type TypeParam struct {
	kind  string
	id    string
	type1 *TypeClang
}

type BuiltinType struct {
	NodeParam
}

type RecordType struct {
	NodeParam
	decl ClangNode
}

type PointerType struct {
	NodeParam
}

type ParenType struct {
	NodeParam
}

type ConstantArrayType struct {
	NodeParam
	size int
}

type TypedefType struct {
	NodeParam
	decl ClangNode
}

type ElaboratedType struct {
	NodeParam
	ownedTagDecl ClangNode
}

type FunctionProtoType struct {
	NodeParam
	cc string
}

type QualType struct {
	NodeParam
	typestr    string
	qualifiers string
}

func (p *QualType) getAsString() string {
	return p.typestr
}

func (p *QualType) isEmpty() bool {
	return len(p.typestr) == 0
}

func (p BuiltinType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p ConstantArrayType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p ElaboratedType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p FunctionProtoType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p ParenType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p PointerType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p QualType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p RecordType) t2go() string {
	panic("NoImpl")
	return ""
}

func (p TypedefType) t2go() string {
	panic("NoImpl")
	return ""
}

func convertBuiltinType(content interface{}) *BuiltinType {
	var inst BuiltinType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		default:
			fmt.Printf("[DBG][BuiltinType]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][BuiltinType]%+v\n", inst)
	return &inst
}

func convertConstantArrayType(content interface{}) *ConstantArrayType {
	var inst ConstantArrayType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "size":
			inst.size = int(v.(float64))
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ConstantArrayType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertElaboratedType(content interface{}) *ElaboratedType {
	var inst ElaboratedType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		case "ownedTagDecl":
			inst.ownedTagDecl = convertNode(v)
		default:
			fmt.Printf("[DBG][ElaboratedType]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][ElaboratedType]%+v\n", inst)
	return &inst
}

func convertFunctionProtoType(content interface{}) *FunctionProtoType {
	var inst FunctionProtoType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "cc":
			inst.cc = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][FunctionProtoType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertParenType(content interface{}) *ParenType {
	var inst ParenType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][ParenType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertPointerType(content interface{}) *PointerType {
	var inst PointerType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][PointerType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertQualType(content interface{}) *QualType {
	var inst QualType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "qualifiers":
			inst.qualifiers = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][QualType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertRecordType(content interface{}) *RecordType {
	var inst RecordType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "decl":
			inst.decl = convertNode(v)
		default:
			fmt.Printf("[DBG][RecordType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertTypedefType(content interface{}) *TypedefType {
	var inst TypedefType
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "type":
			inst.type1 = convertTypeClang(v)
		case "decl":
			inst.decl = convertNode(v)
		case "inner":
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][TypedefType]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}
