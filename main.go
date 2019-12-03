package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
	"time"

	"codegenerator-go/models"
	"codegenerator-go/utils"

	dbutil "github.com/chengli1988/go-dbutil-mysql"
)

func printArg(arg interface{}, verb rune) {
	switch f := arg.(type) {
	default:
		fmt.Println(reflect.ValueOf(f).Kind())
	}

}

func main() {

	dbName := "demo"
	tableName := "sys_user"
	packageName := "system"
	dirName := "target"

	beginTime := time.Now().UnixNano()
	fmt.Println("开始生成代码...")
	// 初始化数据连接池
	dbutil.InitPool("mysql", "root", "root", "127.0.0.1", "3306", dbName, "utf8mb4")
	// 查询数据
	tableModel := queryData(dbName, tableName, packageName)

	// 创建目录
	saveFilePath := filepath.Join(dirName, packageName)
	os.MkdirAll(saveFilePath, os.ModePerm)

	// 生成model文件
	err := generatorModelFile(tableModel, filepath.Join(saveFilePath, tableModel.GetStructVariableName()+"_model.go"))
	if err != nil {
		fmt.Println(err)
	}
	// 生成model文件
	err = generatorHandlerFile(tableModel, filepath.Join(saveFilePath, tableModel.GetStructVariableName()+"_handler.go"))
	if err != nil {
		fmt.Println(err)
	}
	// 生成vue文件
	err = generatorVueFile(tableModel, filepath.Join(saveFilePath, tableModel.GetStructVariableName()+".vue"))
	if err != nil {
		fmt.Println(err)
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("生成代码完成, 耗时 %d 毫秒. \n", (endTime-beginTime)/1e6)
}

// queryData 查询数据
func queryData(dbName string, tableName string, packageName string) models.TableModel {
	var (
		db         dbutil.BaseDB
		tableModel models.TableModel
		idColumn   models.ColumnModel
	)

	tableModel.PackageName = packageName
	tableModel.TableName = tableName

	// 查询主键信息
	columnResult, _ := db.SelectOneBySql("SELECT column_name, data_type, character_maximum_length,"+
		" is_nullable, column_comment, column_default FROM information_schema.columns"+
		" WHERE table_name=? AND table_schema=? and COLUMN_KEY='PRI'", tableName, dbName)
	utils.MapToStruct(columnResult, &idColumn)
	tableModel.IdColumn = idColumn

	// 查询数据库表列信息
	results, _ := db.SelectBySql("SELECT column_name, data_type, character_maximum_length, is_nullable,"+
		" column_comment, column_default FROM information_schema.columns WHERE table_name=? AND table_schema=? "+
		" ORDER BY ORDINAL_POSITION", tableName, dbName)
	columnModels := make([]models.ColumnModel, 0)
	for _, columnMap := range results {

		var columnModel models.ColumnModel

		utils.MapToStruct(columnMap, &columnModel)
		columnModels = append(columnModels, columnModel)
	}

	tableModel.Columns = columnModels

	return tableModel
}

func generatorModelFile(tableModel models.TableModel, filename string) error {
	var modelBuffer bytes.Buffer
	modelTemplate := template.Must(template.New("model.tmpl").ParseFiles("templates/model.tmpl"))
	err := modelTemplate.Execute(&modelBuffer, tableModel)

	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, modelBuffer.Bytes(), os.ModePerm)
}

func generatorHandlerFile(tableModel models.TableModel, filename string) error {
	var handlerBuffer bytes.Buffer
	handlerTemplate := template.Must(template.New("handler.tmpl").ParseFiles("templates/handler.tmpl"))
	err := handlerTemplate.Execute(&handlerBuffer, tableModel)

	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, handlerBuffer.Bytes(), os.ModePerm)
}

func generatorVueFile(tableModel models.TableModel, filename string) error {
	var vueBuffer bytes.Buffer
	vueTemplate := template.Must(template.New("vue.tmpl").ParseFiles("templates/vue.tmpl"))
	err := vueTemplate.Execute(&vueBuffer, tableModel)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, vueBuffer.Bytes(), os.ModePerm)
}
