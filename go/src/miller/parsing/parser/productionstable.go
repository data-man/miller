// Code generated by gocc; DO NOT EDIT.

package parser

import "miller/dsl"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Root	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Root : StatementBlock	<< dsl.NewAST(X[0]) >>`,
		Id:         "Root",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewAST(X[0])
		},
	},
	ProdTabEntry{
		String: `StatementBlock : Statement	<< dsl.NewASTNodeUnary(nil, X[0], dsl.NodeTypeStatementBlock) >>`,
		Id:         "StatementBlock",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(nil, X[0], dsl.NodeTypeStatementBlock)
		},
	},
	ProdTabEntry{
		String: `StatementBlock : StatementBlock ";"	<< dsl.Nestable(X[0]) >>`,
		Id:         "StatementBlock",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.Nestable(X[0])
		},
	},
	ProdTabEntry{
		String: `StatementBlock : StatementBlock ";" Statement	<< dsl.AppendChild(X[0], X[2]) >>`,
		Id:         "StatementBlock",
		NTType:     2,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.AppendChild(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Statement : StatementInBody	<<  >>`,
		Id:         "Statement",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StatementInBody : SrecAssignment	<<  >>`,
		Id:         "StatementInBody",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SrecAssignment : SrecDirectAssignment	<<  >>`,
		Id:         "SrecAssignment",
		NTType:     5,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "||=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "^^=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "&&=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "|=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "^=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "&=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "<<=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName ">>=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "+=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName ".=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "-=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "*=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "/=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "//=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "%!=(MISSING)" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `SrecDirectAssignment : FieldName "**=" RHS	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment) >>`,
		Id:         "SrecDirectAssignment",
		NTType:     6,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeSrecDirectAssignment)
		},
	},
	ProdTabEntry{
		String: `RHS : TernaryTerm	<<  >>`,
		Id:         "RHS",
		NTType:     7,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TernaryTerm : LogicalOrTerm "?" TernaryTerm ":" TernaryTerm	<< dsl.NewASTNodeTernary(X[1], X[0], X[2], X[4], dsl.NodeTypeOperator) >>`,
		Id:         "TernaryTerm",
		NTType:     8,
		Index:      26,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeTernary(X[1], X[0], X[2], X[4], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `TernaryTerm : LogicalOrTerm	<<  >>`,
		Id:         "TernaryTerm",
		NTType:     8,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LogicalOrTerm : LogicalOrTerm "||" LogicalXORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "LogicalOrTerm",
		NTType:     9,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `LogicalOrTerm : LogicalXORTerm	<<  >>`,
		Id:         "LogicalOrTerm",
		NTType:     9,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LogicalXORTerm : LogicalXORTerm "^^" LogicalAndTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "LogicalXORTerm",
		NTType:     10,
		Index:      30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `LogicalXORTerm : LogicalAndTerm	<<  >>`,
		Id:         "LogicalXORTerm",
		NTType:     10,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LogicalAndTerm : LogicalAndTerm "&&" EqneTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "LogicalAndTerm",
		NTType:     11,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `LogicalAndTerm : EqneTerm	<<  >>`,
		Id:         "LogicalAndTerm",
		NTType:     11,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EqneTerm : EqneTerm "=~" CmpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "EqneTerm",
		NTType:     12,
		Index:      34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `EqneTerm : EqneTerm "!=~" CmpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "EqneTerm",
		NTType:     12,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `EqneTerm : EqneTerm "==" CmpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "EqneTerm",
		NTType:     12,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `EqneTerm : EqneTerm "!=" CmpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "EqneTerm",
		NTType:     12,
		Index:      37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `EqneTerm : CmpTerm	<<  >>`,
		Id:         "EqneTerm",
		NTType:     12,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CmpTerm : CmpTerm ">" BitwiseORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "CmpTerm",
		NTType:     13,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `CmpTerm : CmpTerm ">=" BitwiseORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "CmpTerm",
		NTType:     13,
		Index:      40,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `CmpTerm : CmpTerm "<" BitwiseORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "CmpTerm",
		NTType:     13,
		Index:      41,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `CmpTerm : CmpTerm "<=" BitwiseORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "CmpTerm",
		NTType:     13,
		Index:      42,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `CmpTerm : BitwiseORTerm	<<  >>`,
		Id:         "CmpTerm",
		NTType:     13,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitwiseORTerm : BitwiseORTerm "|" BitwiseXORTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "BitwiseORTerm",
		NTType:     14,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `BitwiseORTerm : BitwiseXORTerm	<<  >>`,
		Id:         "BitwiseORTerm",
		NTType:     14,
		Index:      45,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitwiseXORTerm : BitwiseXORTerm "^" BitwiseANDTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "BitwiseXORTerm",
		NTType:     15,
		Index:      46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `BitwiseXORTerm : BitwiseANDTerm	<<  >>`,
		Id:         "BitwiseXORTerm",
		NTType:     15,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitwiseANDTerm : BitwiseANDTerm "&" BitwiseShiftTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "BitwiseANDTerm",
		NTType:     16,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `BitwiseANDTerm : BitwiseShiftTerm	<<  >>`,
		Id:         "BitwiseANDTerm",
		NTType:     16,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BitwiseShiftTerm : BitwiseShiftTerm "<<" AddsubdotTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "BitwiseShiftTerm",
		NTType:     17,
		Index:      50,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `BitwiseShiftTerm : BitwiseShiftTerm ">>" AddsubdotTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "BitwiseShiftTerm",
		NTType:     17,
		Index:      51,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `BitwiseShiftTerm : AddsubdotTerm	<<  >>`,
		Id:         "BitwiseShiftTerm",
		NTType:     17,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : AddsubdotTerm "+" MuldivTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : AddsubdotTerm "-" MuldivTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : AddsubdotTerm ".+" MuldivTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      55,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : AddsubdotTerm ".-" MuldivTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      56,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : AddsubdotTerm "." MuldivTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      57,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `AddsubdotTerm : MuldivTerm	<<  >>`,
		Id:         "AddsubdotTerm",
		NTType:     18,
		Index:      58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm "*" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      59,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm "/" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      60,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm "//" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      61,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm "%!"(MISSING) UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      62,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm ".*" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      63,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm "./" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      64,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : MuldivTerm ".//" UnaryBitwiseOpTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      65,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `MuldivTerm : UnaryBitwiseOpTerm	<<  >>`,
		Id:         "MuldivTerm",
		NTType:     19,
		Index:      66,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : "+" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      67,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : "-" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      68,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : ".+" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      69,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : ".-" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      70,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : "!" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      71,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : "~" PowTerm	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator) >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      72,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `UnaryBitwiseOpTerm : PowTerm	<<  >>`,
		Id:         "UnaryBitwiseOpTerm",
		NTType:     20,
		Index:      73,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PowTerm : AtomOrFunction "**" PowTerm	<< dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator) >>`,
		Id:         "PowTerm",
		NTType:     21,
		Index:      74,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeBinary(X[1], X[0], X[2], dsl.NodeTypeOperator)
		},
	},
	ProdTabEntry{
		String: `PowTerm : AtomOrFunction	<<  >>`,
		Id:         "PowTerm",
		NTType:     21,
		Index:      75,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : "(" RHS ")"	<< dsl.Nestable(X[1]) >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      76,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.Nestable(X[1])
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : FieldName	<<  >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      77,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldName : DirectFieldName	<<  >>`,
		Id:         "FieldName",
		NTType:     23,
		Index:      78,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FieldName : IndirectFieldName	<<  >>`,
		Id:         "FieldName",
		NTType:     23,
		Index:      79,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DirectFieldName : md_token_field_name	<< dsl.NewASTNodeStripDollarPlease(X[0], dsl.NodeTypeDirectFieldName) >>`,
		Id:         "DirectFieldName",
		NTType:     24,
		Index:      80,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeStripDollarPlease(X[0], dsl.NodeTypeDirectFieldName)
		},
	},
	ProdTabEntry{
		String: `IndirectFieldName : "$[" RHS "]"	<< dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeIndirectFieldName) >>`,
		Id:         "IndirectFieldName",
		NTType:     25,
		Index:      81,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeUnary(X[0], X[1], dsl.NodeTypeIndirectFieldName)
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : md_token_string_literal	<< dsl.NewASTNodeStripDoubleQuotePairPlease(
    X[0],
    dsl.NodeTypeStringLiteral,
  ) >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      82,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNodeStripDoubleQuotePairPlease(
    X[0],
    dsl.NodeTypeStringLiteral,
  )
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : md_token_int_literal	<< dsl.NewASTNode(X[0], dsl.NodeTypeIntLiteral) >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      83,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeIntLiteral)
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : md_token_float_literal	<< dsl.NewASTNode(X[0], dsl.NodeTypeFloatLiteral) >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      84,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeFloatLiteral)
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : md_token_boolean_literal	<< dsl.NewASTNode(X[0], dsl.NodeTypeBoolLiteral) >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      85,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeBoolLiteral)
		},
	},
	ProdTabEntry{
		String: `AtomOrFunction : ContextVariable	<<  >>`,
		Id:         "AtomOrFunction",
		NTType:     22,
		Index:      86,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_IPS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      87,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_IFS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      88,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_IRS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      89,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_OPS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      90,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_OFS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      91,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_ORS	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      92,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_NF	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      93,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_NR	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      94,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_FNR	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      95,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_FILENAME	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
	ProdTabEntry{
		String: `ContextVariable : md_token_FILENUM	<< dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable) >>`,
		Id:         "ContextVariable",
		NTType:     26,
		Index:      97,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return dsl.NewASTNode(X[0], dsl.NodeTypeContextVariable)
		},
	},
}
