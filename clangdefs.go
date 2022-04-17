package main

import (
	"fmt"
)

type ClangNode interface {
	//Pos() int
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

func convertInnerNodes(content interface{}) []ClangNode {
	nodes := []ClangNode{}

	arr := content.([]interface{})
	for _, v := range arr {
		di := convertNode(v)
		if di != nil {
			nodes = append(nodes, di)
		}
	}
	return nodes
}

func convertNode(content interface{}) ClangNode {
	varmap := content.(map[string]interface{})
	vkind, exists := varmap["kind"]
	if !exists {
		fmt.Printf("[DBG][Node]No kind. Cannot convert %+v\n", content)
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
	case "CompoundStmt":
		return convertCompoundStmt(content)
	case "DeclStmt":
		return convertDeclStmt(content)
	case "ForStmt":
		return convertForStmt(content)
	case "IfStmt":
		return convertIfStmt(content)
	case "ReturnStmt":
		return convertReturnStmt(content)
	case "NonNullAttr":
		return convertNonNullAttr(content)
	case "NoThrowAttr":
		return convertNoThrowAttr(content)
	case "BinaryOperator":
		return convertBinaryOperator(content)
	case "UnaryOperator":
		return convertUnaryOperator(content)
	case "CompoundAssignOperator":
		return convertCompoundAssignOperator(content)
	case "ImplicitCastExpr":
		return convertImplicitCastExpr(content)
	case "CallExpr":
		return convertCallExpr(content)
	default:
		fmt.Printf("[DBG][Node]Unknown kind: %v\n", vkind)
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
		case "init":
			tud.init1 = v.(string)
		case "isImplicit":
			tud.isImplicit = v.(bool)
		case "isUsed":
			tud.isUsed = v.(bool)
		case "mangledName":
			tud.mangledName = v.(string)
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
		case "mangledName":
			tud.mangledName = v.(string)
		case "isUsed":
			tud.isUsed = v.(bool)
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
			tud.inner = convertInnerNodes(v)
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

func convertCompoundStmt(content interface{}) *CompoundStmt {
	var inst CompoundStmt
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][CompoundStmt]%+v\n", inst)
	return &inst
}

func convertDeclStmt(content interface{}) *DeclStmt {
	var inst DeclStmt
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][DeclStmt]%+v\n", inst)
	return &inst
}

func convertForStmt(content interface{}) *ForStmt {
	var inst ForStmt
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][ForStmt]%+v\n", inst)
	return &inst
}

func convertIfStmt(content interface{}) *IfStmt {
	var inst IfStmt
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][IfStmt]%+v\n", inst)
	return &inst
}

func convertReturnStmt(content interface{}) *ReturnStmt {
	var inst ReturnStmt
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundStmt]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][ReturnStmt]%+v\n", inst)
	return &inst
}

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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][BinaryOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][BinaryOperator]%+v\n", inst)
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][UnaryOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][UnaryOperator]%+v\n", inst)
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
			// ignore
			inst.inner = convertInnerNodes(v)
		default:
			fmt.Printf("[DBG][CompoundAssignOperator]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][CompoundAssignOperator]%+v\n", inst)
	return &inst
}

func convertNoThrowAttr(content interface{}) *NoThrowAttr {
	return nil
}

func convertNonNullAttr(content interface{}) *NonNullAttr {
	return nil
}

func convertImplicitCastExpr(content interface{}) *ImplicitCastExpr {
	return nil
}

func convertCallExpr(content interface{}) *CallExpr {
	return nil
}

// https://clang.llvm.org/doxygen/classclang_1_1Type.html
type TypeClang struct {
	qtype      *QualType
	sugerqtype *QualType
}

func (p *TypeClang) dump() {
	fmt.Printf("%s\n", p.getString())
}

func (p *TypeClang) getString() string {
	if p.sugerqtype.isEmpty() {
		return fmt.Sprintf("Type: %s",
			p.qtype.typestr)
	} else {
		return fmt.Sprintf("Type: %s, desugaredQualType: %s",
			p.qtype.typestr, p.sugerqtype.typestr)
	}
}

type QualType struct {
	typestr string
}

func (p *QualType) getAsString() string {
	return p.typestr
}

func (p *QualType) isEmpty() bool {
	return len(p.typestr) == 0
}

func convertTypeClang(content interface{}) *TypeClang {
	var tc TypeClang
	tc.qtype = new(QualType)
	tc.sugerqtype = new(QualType)

	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "qualType":
			tc.qtype.typestr = v.(string)
		case "desugaredQualType":
			tc.sugerqtype.typestr = v.(string)
		default:
			fmt.Printf("[DBG][TypeClang]Unknown [%v]:%v\n", k, v)
		}
	}
	fmt.Printf("[DBG][TypeClang]%s\n", tc.getString())
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
	type1       *TypeClang
	isImplicit  bool
	name        string
	mangledName string
	init1       string
	isUsed      bool
}

type FunctionDecl struct {
	Decl
	mangledName  string
	storageClass string
	name         string
	type1        *TypeClang // ?
	isUsed       bool
	inner        []ClangNode // Actual type?
}

type BuiltinType struct{}
type RecordType struct{}
type RecordDecl struct {
	Decl
	mangledName string
	isUsed      bool
}
type PointerType struct{}
type ConstantArrayType struct{}
type TypedefType struct{}
type FieldDecl struct{ Decl }
type ElaboratedType struct{}
type VarDecl struct{ Decl }

type ParmVarDecl struct {
	Decl
	mangledName string
	name        string
	type1       *TypeClang // ?
	isUsed      bool
}
type NoThrowAttr struct{}
type RestrictAttr struct{}
type BuiltinAttr struct{}
type FormatAttr struct{}
type AsmLabelAttr struct{}
type ConstAttr struct{}
type PureAttr struct{}
type NonNullAttr struct{}
type ModeAttr struct{}
type CompoundStmt struct {
	kind   string
	id     string
	range1 *SourceRange // "range" is a keyword
	inner  []ClangNode
}
type ReturnStmt struct {
	kind   string
	id     string
	range1 *SourceRange // "range" is a keyword
	inner  []ClangNode
}
type ParenExpr struct{}
type CStyleCastExpr struct{}

type BinaryOperator struct {
	NodeParam
}
type ImplicitCastExpr struct{}
type DeclRefExpr struct{}
type IntegerLiteral struct{}
type IndirectFieldDecl struct{ Decl }
type WarnUnusedResultAttr struct{}
type ParenType struct{}
type FunctionProtoType struct{}

type DeclStmt struct {
	NodeParam
}

type ForStmt struct {
	NodeParam
}

type UnaryOperator struct {
	NodeParam
}

type CompoundAssignOperator struct {
	NodeParam
}
type IfStmt struct {
	NodeParam
}
type CallExpr struct{}
type StringLiteral struct{}
type ArraySubscriptExpr struct{}
type CharacterLiteral struct{}

type NodeParam struct {
	kind   string
	id     string
	range1 *SourceRange // "range" is a keyword
	type1  *TypeClang   // ?
	inner  []ClangNode
}
