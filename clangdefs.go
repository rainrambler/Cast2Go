package main

import (
	"fmt"
)

type ClangNode interface {
	t2go() string
}

type Node interface {
	Desc() string
}

// https://clang.llvm.org/doxygen/classclang_1_1SourceLocation.html
type SourceLocation struct {
	offset       int
	line         int
	col          int
	tokLen       int
	file1        string
	includedFrom *IncludedFrom
}

func convertSourceLocation(content interface{}) *SourceLocation {
	var inst SourceLocation
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "offset":
			inst.offset = int(v.(float64))
		case "col":
			inst.col = int(v.(float64))
		case "tokLen":
			inst.tokLen = int(v.(float64))
		case "line":
			inst.line = int(v.(float64))
		case "file":
			inst.file1 = v.(string)
		case "includedFrom":
			inst.includedFrom = convertIncludedFrom(v)
		default:
			fmt.Printf("[DBG][SourceLocation]Unknown [%v]:%v\n", k, v)
		}
	}

	//fmt.Printf("[DBG][SourceLocation]%+v\n", inst)
	return &inst
}

func convertSourceRange(content interface{}) *SourceRange {
	// TODO implement
	return nil
}

func convertInnerNodes(content interface{}) []ClangNode {
	nodes := []ClangNode{}

	arr := content.([]interface{})
	for _, v := range arr {
		di := convertNode(v)
		if di != nil {
			nodes = append(nodes, di)

			//fmt.Printf("[DBG]Add Node [%v]: %+v\n", &di, di)
		}
	}
	return nodes
}

func convertNode(content interface{}) ClangNode {
	varmap := content.(map[string]interface{})
	if len(varmap) == 0 {
		return nil
	}
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
		return convertVarDecl(content)
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
	case "AsmLabelAttr":
		return convertAsmLabelAttr(content)
	case "BuiltinAttr":
		return convertBuiltinAttr(content)
	case "ConstAttr":
		return convertConstAttr(content)
	case "FormatAttr":
		return convertFormatAttr(content)
	case "NonNullAttr":
		return convertNonNullAttr(content)
	case "NoThrowAttr":
		return convertNoThrowAttr(content)
	case "PureAttr":
		return convertPureAttr(content)
	case "RestrictAttr":
		return convertRestrictAttr(content)
	case "WarnUnusedResultAttr":
		return convertWarnUnusedResultAttr(content)
	case "BinaryOperator":
		return convertBinaryOperator(content)
	case "UnaryOperator":
		return convertUnaryOperator(content)
	case "CompoundAssignOperator":
		return convertCompoundAssignOperator(content)
	case "ArraySubscriptExpr":
		return convertArraySubscriptExpr(content)
	case "ImplicitCastExpr":
		return convertImplicitCastExpr(content)
	case "CallExpr":
		return convertCallExpr(content)
	case "CStyleCastExpr":
		return convertCStyleCastExpr(content)
	case "DeclRefExpr":
		return convertDeclRefExpr(content)
	case "ParenExpr":
		return convertParenExpr(content)
	case "IntegerLiteral":
		return convertIntegerLiteral(content)
	case "CharacterLiteral":
		return convertCharacterLiteral(content)
	case "StringLiteral":
		return convertStringLiteral(content)
	case "BuiltinType":
		return convertBuiltinType(content)
	default:
		fmt.Printf("[DBG][Node]Unknown kind: %v\n", vkind)
		return nil
	}
}

// https://clang.llvm.org/doxygen/classclang_1_1Type.html
type TypeClang struct {
	qtype           *QualType
	sugerqtype      *QualType
	typeAliasDeclId string // point to another id
}

func createTypeClang() *TypeClang {
	var tc TypeClang
	tc.qtype = new(QualType)
	tc.sugerqtype = new(QualType)
	return &tc
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

func (p *TypeClang) t2go() string {
	return p.qtype.typestr // ?
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
		case "typeAliasDeclId":
			tc.typeAliasDeclId = v.(string)
		default:
			fmt.Printf("[DBG][TypeClang]Unknown [%v]:%v\n", k, v)
		}
	}
	//fmt.Printf("[DBG][TypeClang]%s\n", tc.getString())
	return &tc
}

type IncludedFrom struct {
	fromfile string
}

func convertIncludedFrom(content interface{}) *IncludedFrom {
	var inst IncludedFrom
	varmap := content.(map[string]interface{})
	for k, v := range varmap {
		switch k {
		case "file":
			inst.fromfile = v.(string)
		default:
			fmt.Printf("[DBG][IncludedFrom]Unknown [%v]:%v\n", k, v)
		}
	}
	return &inst
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

type TypeParam struct {
	kind  string
	id    string
	type1 *TypeClang
}

type BuiltinType struct {
	TypeParam
}

type RecordType struct{}
type PointerType struct{}
type ConstantArrayType struct{}
type TypedefType struct{}
type ElaboratedType struct{}

type ParenType struct{}
type FunctionProtoType struct{}

func (p BuiltinType) t2go() string {
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

type NodeParam struct {
	id     string
	kind   string
	range1 *SourceRange // "range" is a keyword
	type1  *TypeClang   // ?
	inner  []ClangNode
}
