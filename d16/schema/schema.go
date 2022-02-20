package shcema
// 实现对象与表的转换

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)


// 数据表每列的数据结构
type Field struct {
	Name string  // 列名称
	Type string  // 列类型
	Tag  string  // 列约束条件
}


// 数据表数据结构
type Schema struct {
	Model      interface{}  // 被映射的对象
	Name       string  // 表名称
	Fields     []*Field  // 列对象
	FieldNames []string  // 列名
	fieldMap   map[string]*Field  // 列名与列对象映射关系
}


func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}


func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}


type ITableName interface {
	TableName() string
}


// Parse 实现将任意对象解析为 Schema 实例
func Parse(dest interface{}, d dialect.dialect) *Schema {
	// reflect.Indirect 获取指针指向的实例
	// modelType 结构体
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	var tableName string
	// dest对象实现 ITableName 接口，自定义表名，若没有实现，则取dest对象名称作为表名
	t, ok := dest.(ITableName)
	if !ok {
		tableName = modelType.Name()
	} else {
		tableName = t.TableName()
	}

	schema := &Schema{
		Model:    dest,
		Name:     tableName,
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
