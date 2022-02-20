package dialect
// 用于隔离不同数据库之间的差异，便于拓展
// 使用反射机制获取任意 struct 对象的名称和字段，映射为数据中的表

import "reflect"

var dialectsMap = map[string]Dialect{}


type Dialect interface {
	// 用于将 go 语言的类型转换为该数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	// 返回某个表是否存在的 SQL 语句，参数是表名
	TableExistSQL(tableName string) (string, []interface{})
}


// 注册 Dialect 实例
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}


// 获取 Dialect 实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
