package main

import (
	"fmt"
	"reflect"
)

func testFunc() {
	fmt.Println("This is a test function.")
}

func main() {
	// funcValue := reflect.ValueOf(&testFunc)
	// funcValue = funcValue.Elem()
	// funcName := funcValue.Type().Name()
	fmt.Println(reflect.TypeOf(testFunc).Name())
}
