package main

import (
	"fmt"
	"reflect"
)

type cow struct {
	name  string
	age   int
	grade int
}

func (c *cow) getName(name string) {

}

func reflectStruct() {
	var c = cow{
		name:  "牛马",
		age:   10,
		grade: 15,
	}
	structOps(&c)
}

func structOps(a interface{}) {
	typeA := reflect.TypeOf(a)
	valueA := reflect.ValueOf(a)
	switch typeA.Kind() {
	case reflect.Struct:
		fmt.Println("struct")
	case reflect.Ptr:
		fmt.Println("ptr")
		fieldNum := typeA.NumField()
		for i := 0; i < fieldNum; i++ {
			field := typeA.Field(i)
			fmt.Println("field:", field.Name)
			valueA.Elem().Field(i).SetString("大牛马")
		}
	default:
		fmt.Println("not struct")
	}
}

// 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}
