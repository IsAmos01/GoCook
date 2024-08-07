package main

import (
	"fmt"
	"reflect"
)

func reflectValue() {
	var x = 3.5
	valueOps(&x)
}

func valueOps(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
		v.SetFloat(10.2)
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	default:
		fmt.Println("all not")
	}
}
