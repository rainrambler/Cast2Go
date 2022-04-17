package main

import (
	"fmt"
)

type AttrParam struct {
	id        string
	range1    *SourceRange // "range" is a keyword
	kind      string
	implicit  bool // ?
	inherited bool // ?
}

type NoThrowAttr struct {
	AttrParam
}
type RestrictAttr struct {
	AttrParam
}
type BuiltinAttr struct {
	AttrParam
}
type FormatAttr struct {
	AttrParam
}
type AsmLabelAttr struct {
	AttrParam
}
type ConstAttr struct {
	AttrParam
}
type PureAttr struct {
	AttrParam
}
type NonNullAttr struct {
	AttrParam
}
type ModeAttr struct{}

type WarnUnusedResultAttr struct {
	AttrParam
}

func (p *NoThrowAttr) t2go() string {
	return ""
}

func (p *RestrictAttr) t2go() string {
	return ""
}

func (p *BuiltinAttr) t2go() string {
	return ""
}

func (p *FormatAttr) t2go() string {
	return ""
}

func (p *AsmLabelAttr) t2go() string {
	return ""
}

func (p *ConstAttr) t2go() string {
	return ""
}

func (p *PureAttr) t2go() string {
	return ""
}

func (p *NonNullAttr) t2go() string {
	return ""
}

func (p *ModeAttr) t2go() string {
	return ""
}

func (p *WarnUnusedResultAttr) t2go() string {
	return ""
}

func convertNoThrowAttr(content interface{}) *NoThrowAttr {
	var inst NoThrowAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][NoThrowAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][NoThrowAttr]%+v\n", inst)
	return &inst
}

func convertNonNullAttr(content interface{}) *NonNullAttr {
	var inst NonNullAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][NonNullAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][NonNullAttr]%+v\n", inst)
	return &inst
}

func convertConstAttr(content interface{}) *ConstAttr {
	var inst ConstAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][ConstAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertPureAttr(content interface{}) *PureAttr {
	var inst PureAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][PureAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertRestrictAttr(content interface{}) *RestrictAttr {
	var inst RestrictAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][RestrictAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertFormatAttr(content interface{}) *FormatAttr {
	var inst FormatAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][FormatAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertAsmLabelAttr(content interface{}) *AsmLabelAttr {
	var inst AsmLabelAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][AsmLabelAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertBuiltinAttr(content interface{}) *BuiltinAttr {
	var inst BuiltinAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][BuiltinAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}

func convertWarnUnusedResultAttr(content interface{}) *WarnUnusedResultAttr {
	var inst WarnUnusedResultAttr
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			inst.id = v.(string)
		case "kind":
			inst.kind = v.(string)
		case "implicit":
			inst.implicit = v.(bool)
		case "inherited":
			inst.inherited = v.(bool)
		case "range":
			inst.range1 = convertSourceRange(v)
		default:
			fmt.Printf("[DBG][WarnUnusedResultAttr]Unknown [%v]:%v\n", k, v)
		}
	}

	return &inst
}
