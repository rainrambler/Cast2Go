package main

import (
	"fmt"
	"log"
)

type ClangDecl interface {
	getKind() string
}

type Decl struct {
	sourceRg *SourceRange
	kind     string
	id       string
	loc      *SourceLocation
	range1   *SourceRange // "range" is a keyword
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
	inner        []ClangNode
}

func (p *TypedefDecl) t2go() string {
	if p.isImplicit {
		return ""
	}

	// TODO
	return ""
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

func (p *FunctionDecl) t2go() string {
	if !p.hasValidFileLocation() && (p.name != "main") {
		log.Printf("[DBG]System Func: %v\n", p.name)
		return ""
	}

	if !p.hasStmt() {
		// forword declaration
		return ""
	}

	s := ""
	fstart := fmt.Sprintf("func %s (", p.name)
	s += fstart

	rettype := p.type1.getReturnType()

	if len(p.inner) == 0 {
		s += ") " + rettype + "{}"
		return s
	}

	pos := 0
	for {
		if pos >= len(p.inner) {
			break
		}

		pvd := p.inner[pos] // the first should be ParmVarDecl
		if isParmVarDecl(pvd) {
			decl := pvd.(*ParmVarDecl)
			s += decl.t2go() + ","
			pos++
		} else {
			//otherType := reflect.TypeOf(pvd)
			//fmt.Printf("Unknown Node: %+v\n", otherType)
			break
		}
	}

	s = RemoveLastSubStr(s, ",")
	s += ") " + rettype + "{" + EnterStr
	//s += ") {\n"
	for {
		if pos >= len(p.inner) {
			break
		}

		stmt := p.inner[pos]
		s += stmt.t2go() + EnterStr
		pos++
	}

	s += "}" + EnterStr
	return s
}

func (p *FunctionDecl) hasValidFileLocation() bool {
	if p.loc == nil {
		return false
	}

	if p.storageClass == "extern" {
		return false
	}

	if isSystemPath(p.loc.file1) {
		return false
	}

	if (p.loc.includedFrom != nil) && p.loc.includedFrom.isFromSystem() {
		return false
	}

	return true
}

func (p *FunctionDecl) hasStmt() bool {
	for _, nd := range p.inner {
		switch nd.(type) {
		case *CompoundStmt, *DeclStmt, *ForStmt, *IfStmt, *ReturnStmt:
			return true
		default:
		}
	}
	return false
}

func isStmt(v ClangNode) bool {
	switch v.(type) {
	case *ParmVarDecl:
		//fmt.Printf("[DBG]PVD: %+v\n", v)
		return true
	default:
		return false
	}
}

func isParmVarDecl(v ClangNode) bool {
	switch v.(type) {
	case *ParmVarDecl:
		//fmt.Printf("[DBG]PVD: %+v\n", v)
		return true
	default:
		return false
	}
}

type RecordDecl struct {
	Decl
	mangledName string
	isUsed      bool
	tagUsed     string
}
type FieldDecl struct{ Decl }

type VarDecl struct {
	Decl
	isUsed      bool
	name        string
	mangledName string
	init1       string
	type1       *TypeClang
	inner       []ClangNode
}

type ParmVarDecl struct {
	Decl
	mangledName string
	name        string
	type1       *TypeClang
	isUsed      bool
}

func (p ParmVarDecl) t2go() string {
	return p.name + " " + p.type1.t2go()
}

type IndirectFieldDecl struct{ Decl }

type TranslationUnitDecl struct {
	Decl

	inner []ClangNode
}

func (p *TranslationUnitDecl) t2go() string {
	s := ""
	for _, nd := range p.inner {
		s += nd.t2go()
	}
	return s
}

func (p *RecordDecl) t2go() string {
	if p.isSystemDecl() {
		return ""
	}
	if p.mangledName == "" {
		return ""
	}
	return p.mangledName + "!!!"
}

func (p *RecordDecl) isSystemDecl() bool {
	if p.loc == nil {
		return true
	}

	return len(p.loc.file1) == 0
}

func (p *FieldDecl) t2go() string {
	panic("NoImpl")
	return ""
}

func (p *VarDecl) t2go() string {
	//return "var " + p.name + " " + p.type1.t2go() + EnterStr
	return "var " + p.name + " " + p.type1.t2go()
}

func (p *IndirectFieldDecl) t2go() string {
	panic("NoImpl")
	return ""
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

	//fmt.Printf("[DBG]%+v\n", tud)
	/*
		for _, v := range tud.inner {
			fmt.Printf("[DBG]%+v\n", v)
		}
	*/
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
			tud.inner = convertInnerNodes(v)
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
		case "mangledName":
			tud.mangledName = v.(string)
		case "name":
			tud.name = v.(string)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "type":
			tud.type1 = convertTypeClang(v)
		case "init":
			tud.init1 = v.(string)
		case "inner":
			tud.inner = convertInnerNodes(v)
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
