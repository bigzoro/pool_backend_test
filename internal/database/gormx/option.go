package gormx

import "gorm.io/gorm"

type Option struct {
	DB          *gorm.DB
	Selects     []any
	Omits       []any
	IgnoreTotal bool
}

type OptionFunc func(*Option)

// DB 使用传入的Db对象
func DB(db *gorm.DB) OptionFunc {
	return func(o *Option) {
		o.DB = db
	}
}

// Session 创建会话
func Session(session *gorm.Session) OptionFunc {
	return func(o *Option) {
		o.DB = globalDB.Session(session)
	}
}

// Select 指定需要查询的字段
func Select(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Selects = append(o.Selects, columns...)
	}
}

// Omit 指定需要忽略的字段
func Omit(columns ...any) OptionFunc {
	return func(o *Option) {
		o.Omits = append(o.Omits, columns...)
	}
}

// IgnoreTotal 分页查询忽略总数 issue: https://github.com/acmestack/gorm-plus/issues/37
func IgnoreTotal() OptionFunc {
	return func(o *Option) {
		o.IgnoreTotal = true
	}
}
