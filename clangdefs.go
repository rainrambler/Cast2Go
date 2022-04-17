package main

import (
	"fmt"
)

type ClangNode interface {
	Pos() int
}

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

type TranslationUnitDecl struct {
	Decl

	inner []ClangDecl
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
			tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
		}
	}

	fmt.Printf("[DBG]%+v\n", tud)
	for _, v := range tud.inner {
		fmt.Printf("[DBG]%+v\n", v)
	}
	return &tud
}

func convertSourceLocation(content interface{}) *SourceLocation {
	return nil
}

func convertSourceRange(content interface{}) *SourceRange {
	return nil
}

func convertInner(content interface{}) []ClangDecl {
	decls := []ClangDecl{}

	arr := content.([]interface{})
	for _, v := range arr {
		di := convertDecl(v)
		if di != nil {
			decls = append(decls, di)
		}
	}
	return decls
}

func convertDecl(content interface{}) ClangDecl {
	varmap := content.(map[string]interface{})
	vkind, exists := varmap["kind"]
	if !exists {
		fmt.Printf("[DBG]No kind. Cannot convert %+v\n", content)
		return nil
	}

	switch vkind {
	case "TranslationUnitDecl":
		fmt.Printf("[DBG]Multi TranslationUnitDecl. Cannot convert %+v\n", content)
		return nil
	case "TypedefDecl":
		return convertTypedefDecl(content)
	case "RecordDecl":
		return convertRecordDecl(content)
	case "FieldDecl":
		return convertFieldDecl(content)
	case "VarDecl":
		return convertTypedefDecl(content)
	case "FunctionDecl":
		return convertFunctionDecl(content)
	case "ParmVarDecl":
		return convertParmVarDecl(content)
	case "IndirectFieldDecl":
		return convertIndirectFieldDecl(content)

	default:
		fmt.Printf("[DBG]Unknown kind: %v\n", vkind)
		return nil
	}
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
		case "isImplicit":
			tud.isImplicit = v.(bool)
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG]Unknown [%v]:%v in TypedefDecl\n", k, v)
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
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
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
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
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
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
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
		case "isUsed":
			tud.isUsed = v.(bool)
		case "inner":
			tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG]Unknown [%v]:%v in FunctionDecl\n", k, v)
		}
	}

	//fmt.Printf("[DBG][FunctionDecl]%+v\n", tud)
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
		case "inner":
		// ignore
		//tud.inner = convertInner(v)
		default:
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
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
			fmt.Printf("[DBG]Unknown [%v]:%v\n", k, v)
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

// https://clang.llvm.org/doxygen/classclang_1_1Type.html
type TypeClang struct {
	qtype *QualType
}

func (p *TypeClang) dump() {
	fmt.Printf("Type: %s\n", p.qtype.typestr)
}

type QualType struct {
	typestr string
}

func (p *QualType) getAsString() string {
	return p.typestr
}

func convertTypeClang(content interface{}) *TypeClang {
	var tc TypeClang
	tc.qtype = new(QualType)
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "qualType":
			tc.qtype.typestr = v.(string)
		default:
			fmt.Printf("[DBG][TypeClang]Unknown [%v]:%v\n", k, v)
		}
	}
	fmt.Printf("[DBG][TypeClang]%s\n", tc.qtype.getAsString())
	return &tc
}

// https://clang.llvm.org/doxygen/classclang_1_1DeclContext.html
type DeclContext struct {
}

// https://clang.llvm.org/doxygen/classclang_1_1Stmt.html
type Stmt struct {
}

// https://clang.llvm.org/doxygen/classclang_1_1SourceLocation.html
type SourceLocation struct {
}

// https://clang.llvm.org/doxygen/classclang_1_1SourceRange.html
type SourceRange struct {
	beginLoc *SourceLocation
	endLoc   *SourceLocation
}

func (p *SourceRange) getBegin() *SourceLocation {
	return p.beginLoc
}

func (p *SourceRange) getEnd() *SourceLocation {
	return p.endLoc
}

func (p *SourceRange) setBegin(sl *SourceLocation) {
	p.beginLoc = sl
}

func (p *SourceRange) setEnd(sl *SourceLocation) {
	p.endLoc = sl
}

func (p *SourceRange) isValid() bool {
	return true
}

type TypedefDecl struct {
	Decl
	type1      *TypeClang
	isImplicit bool
	name       string
}

type FunctionDecl struct {
	Decl
	mangledName  string
	storageClass string
	name         string
	type1        *TypeClang // ?
	isUsed       bool
	inner        []ClangDecl // Actual type?
}

type BuiltinType struct{}
type RecordType struct{}
type RecordDecl struct{ Decl }
type PointerType struct{}
type ConstantArrayType struct{}
type TypedefType struct{}
type FieldDecl struct{ Decl }
type ElaboratedType struct{}
type VarDecl struct{ Decl }

type ParmVarDecl struct{ Decl }
type NoThrowAttr struct{}
type RestrictAttr struct{}
type BuiltinAttr struct{}
type FormatAttr struct{}
type AsmLabelAttr struct{}
type ConstAttr struct{}
type PureAttr struct{}
type NonNullAttr struct{}
type ModeAttr struct{}
type CompoundStmt struct{}
type ReturnStmt struct{}
type ParenExpr struct{}
type CStyleCastExpr struct{}
type BinaryOperator struct{}
type ImplicitCastExpr struct{}
type DeclRefExpr struct{}
type IntegerLiteral struct{}
type IndirectFieldDecl struct{ Decl }
type WarnUnusedResultAttr struct{}
type ParenType struct{}
type FunctionProtoType struct{}
type DeclStmt struct{}
type ForStmt struct{}
type UnaryOperator struct{}
type CompoundAssignOperator struct{}
type IfStmt struct{}
type CallExpr struct{}
type StringLiteral struct{}
type ArraySubscriptExpr struct{}
type CharacterLiteral struct{}
