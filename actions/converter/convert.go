package converter

import (
	"regexp"
	"strings"
)

var indent string = "    "

var defaultClassName string = "Generated"

//基本类型
var typeMappings map[string]string = make(map[string]string)

func init() {
	//数字类型
	typeMappings["integer"] = "Integer"
	typeMappings["3"] = "Float"
	//字符串类型
	typeMappings["string"] = "String"
	//布尔类型
	typeMappings["bool"] = "Boolean"
}

func replace(template string, className string, body string) string {
	classMatcher, _ := regexp.Compile("\\$className")
	bodyMatcher, _ := regexp.Compile("\\$body")

	template = classMatcher.ReplaceAllLiteralString(template, className)
	template = bodyMatcher.ReplaceAllLiteralString(template, body)
	return template
}

func classToString(classDesc *javaClassDescription) string {
	if classDesc == nil {
		panic("null refer")
	}

	//genField:: []field -> string
	var fieldStringList []string
	for _, field := range classDesc.fields {
		var annotationString = mkAnnotations(field.fieldAnnotation)
		if annotationString != "" {
			fieldStringList = append(fieldStringList, annotationString)
		}
		fieldString := addIndent("private "+field.filedType+" "+field.filedName+";", indent)
		fieldStringList = append(fieldStringList, fieldString)
	}

	body := strings.Join(fieldStringList, "\r\n")

	//genClass:: []innerClass -> string
	for _, innerClass := range classDesc.innerClass {
		body = body + "\r\n" + addIndent(classToString(&innerClass), indent)
	}

	var classTemplate string
	if classDesc.isPublic {
		classTemplate = publicClassTemplate
	} else {
		classTemplate = innerClassTemplate
	}
	return replace(classTemplate, classDesc.className, body)
}

func FileDescriptionToString(description *javaFileDescription) string {
	packageString := mkPackage(description.packageName)
	importString := mkImports(description.imports)
	commentString := description.fileComment
	classContent := classToString(description.classDesc)
	return strings.Join([]string{packageString, importString, commentString, classContent}, "\n\n")
}

func mkImports(imports []javaImport) string {
	var importArray []string
	for _, value := range imports {
		importArray = append(importArray, "import "+value.importedClass+";")
	}
	return strings.Join(importArray, "\n")
}

func mkPackage(name string) string {
	return "package " + name + ";"
}

func mkAnnotations(annotation fieldAnnotation) string {
	var annotationStringArray []string
	if len(annotation.annotations) > 0 {
		for _, annotation := range annotation.annotations {
			annotationStringArray = append(annotationStringArray, mkAnnotation(annotation))
		}
		return strings.Join(annotationStringArray, "\n")
	}
	return ""
}

func mkAnnotation(item annotation) string {
	return addIndent("@"+item.name+"(\""+item.value+"\")", indent)
}

func isUpperCaseFirst(key string) bool {
	runes := []rune(key)
	if (runes[0] >= 'a') && (runes[0] <= 'z') {
		return false
	}
	return true
}

func addIndent(multiLine string, indent string) string {
	splitString := strings.Split(multiLine, "\n")
	var indentedStrings []string

	for _, splitItem := range splitString {
		indentString := indent + splitItem
		indentedStrings = append(indentedStrings, indentString)
	}
	return strings.Join(indentedStrings, "\n")
}
