package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"unicode"
)

// MapToStruct 将 Map 转换成 struct
func MapToStruct(mapResult map[string]interface{}, targetStruct interface{}) {

	jsonResult, err := json.Marshal(mapResult)

	if err != nil {
		log.Println(err)
	} else {
		json.Unmarshal(jsonResult, targetStruct)
	}

}

// ToCamelCase 将字符串转换成帕斯卡命名法格式，e.g. MyName；
// 1、根据下划线拆分成字符串数组；
// 2、将字符串数组的每个字符串首字母变为大写。
func ToCamelCase(str string) string {

	strArray := strings.Split(str, "_")
	var stringBuffer bytes.Buffer
	for _, param := range strArray {
		runes := []rune(param)
		runes[0] = unicode.ToUpper(runes[0])
		stringBuffer.WriteString(string(runes))
	}

	return stringBuffer.String()
}
