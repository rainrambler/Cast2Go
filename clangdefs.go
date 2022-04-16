package main

import (
	"fmt"
)

type Decl struct {
	sourceLoc *SourceLocation
	sourceRg  *SourceRange
	kind      string
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
	id  string
	loc map[interface{}]interface{}
}

func (p *TranslationUnitDecl) convert(content interface{}) {
	varmap := content.(map[string]interface{})
	vid, exists := varmap["id"]
	if !exists {
		fmt.Printf("[DBG]No ID in {%+v}!\n", content)
		return
	}
	p.id = vid.(string)
}

// https://clang.llvm.org/doxygen/classclang_1_1Type.html
type ClangType struct {
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

type TypedefDecl struct{}
type BuiltinType struct{}
type RecordType struct{}
type RecordDecl struct{}
type PointerType struct{}
type ConstantArrayType struct{}
type TypedefType struct{}
type FieldDecl struct{}
type ElaboratedType struct{}
type VarDecl struct{}
type FunctionDecl struct{}
type ParmVarDecl struct{}
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
type IndirectFieldDecl struct{}
type QualType struct{}
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
