package converter

import (
	"reflect"
	"strings"
	"sort"
)

func hasList(description *javaClassDescription) bool {
	for _, value := range description.fields {
		if strings.Contains(value.filedType, "List<") {
			return true
		}
	}
	for _, value := range description.innerClass {
		if hasList(&value) {
			return true
		}
	}
	return false
}

func keys(maps map[string]interface{}) []string {
	var keys []string
	for key, _ := range maps {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func findParent(classDesc *javaClassDescription) *javaClassDescription {
	if classDesc.parentClass == nil {
		return classDesc
	}
	return findParent(classDesc.parentClass)
}

func listType(elmType string) string {
	return "List<" + elmType + ">"
}

//识别基本类型
func parseType(value interface{}) string {
	valueType := reflect.TypeOf(value)
	//基本类型 数字
	if strings.Contains(valueType.Name(), GO_Number) {
		var javaType string
		//判断是否是整数, fixme 0.00 也会识别为整数
		if float64(int(value.(float64)))-value.(float64) != 0 {
			javaType = JAVA_Float
		} else {
			javaType = JAVA_Integer
		}
		return javaType
	}
	//字符串
	if valueType.Name() == GO_String {
		return JAVA_String
	}
	//布尔值
	if valueType.Name() == GO_Bool {
		return JAVA_Bool
	}
	panic("传参非基本类型:" + valueType.String())
}
