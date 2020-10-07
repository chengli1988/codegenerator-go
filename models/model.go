package models

import (
	"strings"
	"unicode"

	"codegenerator-go/utils"
)

// TableModel 表实体
type TableModel struct {
	PackageName string
	TableName   string
	IdColumn    ColumnModel
	Columns     []ColumnModel
}

// ColumnModel 列结构体
type ColumnModel struct {
	ColumnName             string `json:"column_name"`
	DataType               string `json:"data_type"`
	CharacterMaximumLength string `json:"character_maximum_length"`
	IsNullable             string `json:"is_nullable"`
	ColumnComment          string `json:"column_comment"`
	ColumnDefault          string `json:"column_default"`
}

// GetSupperTableName 获取表大写名称(去掉下划线)
// sys_user ===> SYSUSER
func (table TableModel) GetSupperTableName() string {
	return strings.ReplaceAll(strings.ToUpper(table.TableName), "_", "")
}

// GetTableName 获取表名称(去掉下划线)
// sys_user ===> SysUser
func (table TableModel) GetTableName() string {
	return utils.ToCamelCase(table.TableName)
}

// GetStructName 获取结构体名称(去掉前缀、下划线)
// sys_user ===> User
func (table TableModel) GetStructName() string {

	// 去掉前缀
	index := strings.Index(table.TableName, "_")
	if index != -1 {
		index++
	}

	tableName := utils.ToCamelCase(table.TableName[index:])

	return tableName
}

// GetTableNameNoPrefix 获取数据库表名(去掉前缀)
//  sys_user ===> user
func (table TableModel) GetTableNameNoPrefix() string {
	// 去掉前缀
	index := strings.Index(table.TableName, "_")
	if index != -1 {
		index++
	}
	return table.TableName[index:]
}

// GetStructVariableName 获取结构体变量名称(去掉前缀、下划线)
// sys_user ===> user
func (table TableModel) GetStructVariableName() string {

	// 去掉前缀
	index := strings.Index(table.TableName, "_")
	if index != -1 {
		index++
	}

	variableName := utils.ToCamelCase(table.TableName[index:])

	// 首字母变小写
	variableNameRune := []rune(variableName)
	variableNameRune[0] = unicode.ToLower(variableNameRune[0])

	return string(variableNameRune)
}

// GetSupperColumnName 获取列大写名称(去掉下划线)
// user_id ===> USERID
func (column ColumnModel) GetSupperColumnName() string {
	return strings.ReplaceAll(strings.ToUpper(column.ColumnName), "_", "")
}

// GetStructFieldName 获取结构体字段名称(去掉下划线，首字母大写)
// user_id ===> UserId
func (column ColumnModel) GetStructFieldName() string {
	return utils.ToCamelCase(column.ColumnName)
}

// GetTagJsonName 获取结构体json格式名称(去掉下划线，首字母小写)
// user_id ===> userId
func (column ColumnModel) GetTagJsonName() string {
	columnName := utils.ToCamelCase(column.ColumnName)

	columnNameRunes := []rune(columnName)
	columnNameRunes[0] = unicode.ToLower(columnNameRunes[0])

	return string(columnNameRunes)
}

// GetVarType 获取列对应的Go数据类型
func (column ColumnModel) GetVarType() string {
	varType := "string"

	switch column.DataType {
	case "varchar", "datetime":
		varType = "string"
	case "number":
		varType = "float64"
	case "int":
		varType = "int"
	}

	return varType
}
