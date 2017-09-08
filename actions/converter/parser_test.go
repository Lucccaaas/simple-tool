package converter

import (
	"testing"
	"fmt"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestNumber(t *testing.T) {
	var jsonObj *map[string]interface{}
	jsonObj = ReadJson("./static/1.json")
	if jsonObj != nil {
		t.Log("JSON parser success", checkMark)
	}
	{
		t.Log("convert json to class", checkMark)
		classDesc := JsonToJavaClass("Demo", *jsonObj)
		t.Log("convert json success", checkMark)
		if len(classDesc.fields) <= 0 {
			t.Fatal("expect 1 or more field", ballotX)
		} else {
			if classDesc.fields[0].filedName != "aNumberKey" {
				t.Fatal("integer key is error", classDesc.fields[0].filedName, ballotX)
			}
			if classDesc.fields[0].filedType != JAVA_Integer {
				t.Fatal("integer type is error", classDesc.fields[0].filedType, ballotX)
			}

			if classDesc.fields[1].filedName != "bFloatKey" {
				t.Fatal("float key is error", ballotX)
			}
			if classDesc.fields[1].filedType != JAVA_Float {
				t.Fatal("float type is error", ballotX)
			}

			if classDesc.fields[2].filedName != "cStringKey" {
				t.Fatal("string key is error", classDesc.fields[2].filedName, ballotX)
			}
			if classDesc.fields[2].filedType != JAVA_String {
				t.Fatal("string type is error", classDesc.fields[2].filedType, ballotX)
			}

			if classDesc.fields[3].filedName != "dBoolKey" {
				t.Fatal("boolkey is error", classDesc.fields[3].filedName, ballotX)
			}
			if classDesc.fields[3].filedType != JAVA_Bool {
				t.Fatal("booltype is error", classDesc.fields[3].filedType, ballotX)
			}

			if len(classDesc.innerClass) == 0 {
				t.Fatal("inner class is error", classDesc.innerClass, ballotX)
			}

			if classDesc.fields[4].filedType != "EObjKey" {
				t.Fatal("nested obj is error", classDesc.fields[4].filedType, ballotX)
			}

			if classDesc.fields[5].filedType != "List<Integer>" {
				t.Fatal("array obj type is error", classDesc.fields[5].filedType, ballotX)
			}

			if classDesc.fields[6].filedType != "List<String>" {
				t.Fatal("array obj type is error", classDesc.fields[6].filedType, ballotX)
			}

			if classDesc.fields[7].filedType != "List<HArrayObj>" {
				t.Fatal("array obj type is error", classDesc.fields[7].filedType, ballotX)
			}

			t.Log("checke filed success", checkMark)
		}
		t.Log("generator JavaClass")
		{

			content := Convert(ClassToJavaFile(&classDesc))
			fmt.Println(content)
			WriteToFile(content, "./static/"+classDesc.className+".java")
		}
	}

}

func TestEqual(t *testing.T) {
	a := javaImport{"lombok.AllArgsConstructor"}
	b := javaImport{"lombok.Data"}
	fmt.Println(a)
	fmt.Println(b)
}
