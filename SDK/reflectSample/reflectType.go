package main

import (
	"fmt"
	"reflect"
)

func reflectType() {
	var x = 3.5
	typeOps(x)
}

func typeOps(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	default:
		fmt.Println("all not")
	}
	
	// 函数类型入参第i个参数的类型
	// fmt.Println("t.in: ", t.In(1))
	fmt.Println("t.name: ", t.Name())
	fmt.Println("t.Align: ", t.Align())
	// 返回结构体的第i个字段
	// fmt.Println("t.Field: ", t.Field(1))
	
	// 返回结构体的字段
	fmt.Println("t.FieldByIndex: ", t.FieldByIndex([]int{1, 2}))
}
