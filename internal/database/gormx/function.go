package gormx

import (
	"pool/internal/database/gormx/constants"
	"strings"
)

type Function struct {
	funcStr string
}

func (f *Function) As(asName any) string {
	return f.funcStr + " " + constants.As + " " + getColumnName(asName)
}

func (f *Function) Eq(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Eq, value)
}

func (f *Function) Ne(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Ne, value)
}

func (f *Function) Gt(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Gt, value)
}

func (f *Function) Ge(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Ge, value)
}

func (f *Function) Lt(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Lt, value)
}

func (f *Function) Le(value int64) (string, int64) {
	return buildFuncStr(f.funcStr, constants.Le, value)
}

func (f *Function) In(values ...any) (string, []any) {
	// 构建占位符
	placeholder := buildPlaceholder(values)
	return f.funcStr + " " + constants.In + placeholder.String(), values
}

func (f *Function) NotIn(values ...any) (string, []any) {
	// 构建占位符
	placeholder := buildPlaceholder(values)
	return f.funcStr + " " + constants.Not + " " + constants.In + placeholder.String(), values
}

func (f *Function) Between(start int64, end int64) (string, int64, int64) {
	return f.funcStr + " " + constants.Between + " ? " + constants.And + " ?", start, end
}

func (f *Function) NotBetween(start int64, end int64) (string, int64, int64) {
	return f.funcStr + " " + constants.Not + " " + constants.Between + " ? " + constants.And + " ?", start, end
}

func Sum(columnName any) *Function {
	return &Function{funcStr: addBracket(constants.SUM, getColumnName(columnName))}
}

func Avg(columnName any) *Function {
	return &Function{funcStr: addBracket(constants.AVG, getColumnName(columnName))}
}

func Max(columnName any) *Function {
	return &Function{funcStr: addBracket(constants.MAX, getColumnName(columnName))}
}

func Min(columnName any) *Function {
	return &Function{funcStr: addBracket(constants.MIN, getColumnName(columnName))}
}

func Count(columnName any) *Function {
	return &Function{funcStr: addBracket(constants.COUNT, getColumnName(columnName))}
}

func As(columnName any, asName any) string {
	return getColumnName(columnName) + " " + constants.As + " " + getColumnName(asName)
}

func addBracket(function string, columnNameStr string) string {
	return function + constants.LeftBracket + columnNameStr + constants.RightBracket
}

func buildFuncStr(funcStr string, typeStr string, value int64) (string, int64) {
	return funcStr + " " + typeStr + " ?", value
}

func buildPlaceholder(values []any) strings.Builder {
	var placeholder strings.Builder
	placeholder.WriteString(constants.LeftBracket)
	for i := 0; i < len(values); i++ {
		if i == len(values)-1 {
			placeholder.WriteString("?")
			placeholder.WriteString(constants.RightBracket)
			break
		}
		placeholder.WriteString("?")
		placeholder.WriteString(",")
	}
	return placeholder
}
