package main

import (
	"fmt"
)

type ClangDecl interface {
	getKind() string
}

type Decl struct {
	sourceLoc *SourceLocation
	sourceRg  *SourceRange
	kind      string
	id        string
	loc       *SourceLocation
	range1    *SourceRange // "range" is a keyword
}

func (p *Decl) getSourceRange() *SourceRange {
	return p.sourceRg
}

// TODO real kind type?
func (p *Decl) getKind() string {
	return p.kind
}

// https://clang.llvm.org/doxygen/classclang_1_1DeclContext.html
type DeclContext struct {
}

type TypedefDecl struct {
	Decl
	type1        *TypeClang
	name         string
	mangledName  string
	init1        string
	isUsed       bool
	isReferenced bool
	isImplicit   bool
}

type FunctionDecl struct {
	Decl
	mangledName  string
	storageClass string
	name         string
	type1        *TypeClang // ?
	isUsed       bool
	isImplicit   bool
	previousDecl string
	inner        []ClangNode // Actual type?
}

func (p *FunctionDecl) dump() {
	fmt.Printf("func %s, storage: %s, isUsed: %v, type: %v\n",
		p.name, p.storageClass, p.isUsed, p.type1.getString())
	if len(p.inner) > 0 {
		for _, child := range p.inner {
			fmt.Printf("	%+v\n", child)
		}
		fmt.Println("-----------------------")
	}
}

type RecordDecl struct {
	Decl
	mangledName string
	isUsed      bool
	tagUsed     string
}
type FieldDecl struct{ Decl }
type VarDecl struct{ Decl }

type ParmVarDecl struct {
	Decl
	mangledName string
	name        string
	type1       *TypeClang // ?
	isUsed      bool
}

type IndirectFieldDecl struct{ Decl }

type TranslationUnitDecl struct {
	Decl

	inner []ClangNode
}

func convertTranslationUnitDecl(content interface{}) *TranslationUnitDecl {
	var tud TranslationUnitDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "inner":
			tud.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][TranslationUnitDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	fmt.Printf("[DBG]%+v\n", tud)
	for _, v := range tud.inner {
		fmt.Printf("[DBG]%+v\n", v)
	}
	return &tud
}

func convertTypedefDecl(content interface{}) *TypedefDecl {
	var tud TypedefDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "type":
			tud.type1 = convertTypeClang(v)
		case "name":
			tud.name = v.(string)
		case "init":
			tud.init1 = v.(string)
		case "isImplicit":
			tud.isImplicit = v.(bool)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "isReferenced":
			tud.isReferenced = v.(bool)
		case "mangledName":
			tud.mangledName = v.(string)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][TypedefDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	fmt.Printf("[DBG][TypedefDecl]%+v\n", tud)
	return &tud
}

func convertRecordDecl(content interface{}) *RecordDecl {
	var tud RecordDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "mangledName":
			tud.mangledName = v.(string)
		case "tagUsed":
			tud.tagUsed = v.(string)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][RecordDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG]%+v\n", tud)
	return &tud
}

func convertFieldDecl(content interface{}) *FieldDecl {
	var tud FieldDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][FieldDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG]%+v\n", tud)
	return &tud
}

func convertVarDecl(content interface{}) *VarDecl {
	var tud VarDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][VarDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG]%+v\n", tud)
	return &tud
}

func convertFunctionDecl(content interface{}) *FunctionDecl {
	var tud FunctionDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "type":
			tud.type1 = convertTypeClang(v)
		case "name":
			tud.name = v.(string)
		case "mangledName":
			tud.mangledName = v.(string)
		case "storageClass":
			tud.storageClass = v.(string)
		case "previousDecl":
			tud.previousDecl = v.(string)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "isImplicit":
			tud.isImplicit = v.(bool)
		case "inner":
			tud.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][FunctionDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	tud.dump()
	return &tud
}

func convertParmVarDecl(content interface{}) *ParmVarDecl {
	var tud ParmVarDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "type":
			tud.type1 = convertTypeClang(v)
		case "name":
			tud.name = v.(string)
		case "mangledName":
			tud.mangledName = v.(string)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][ParmVarDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG]%+v\n", tud)
	return &tud
}

func convertIndirectFieldDecl(content interface{}) *IndirectFieldDecl {
	var tud IndirectFieldDecl
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "id":
			tud.id = v.(string)
		case "kind":
			tud.kind = v.(string)
		case "loc":
			tud.loc = convertSourceLocation(v)
		case "range":
			tud.range1 = convertSourceRange(v)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG][IndirectFieldDecl]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG]%+v\n", tud)
	return &tud
}

func (p *TranslationUnitDecl) getKind() string {
	return p.kind
}

func (p *TypedefDecl) getKind() string {
	return p.kind
}

func (p *RecordDecl) getKind() string {
	return p.kind
}

func (p *FieldDecl) getKind() string {
	return p.kind
}

func (p *VarDecl) getKind() string {
	return p.kind
}

func (p *FunctionDecl) getKind() string {
	return p.kind
}

func (p *ParmVarDecl) getKind() string {
	return p.kind
}

func (p *IndirectFieldDecl) getKind() string {
	return p.kind
}