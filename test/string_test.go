package main

import (
	"testing"
	"fmt"
)

func TestString(t *testing.T) {
	str := "Abc"
	fmt.Println(string(str[0] + 32))
}
