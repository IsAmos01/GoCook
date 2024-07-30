package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// --- cpu 分析示例 start---
	// 创建cpu分析文件
	fc, err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Println("create cpu.prof err:", err.Error())
		return
	}
	defer fc.Close()
	
	// 开始分析cpu
	err = pprof.StartCPUProfile(fc)
	if err == nil {
		defer pprof.StopCPUProfile()
	}
	// --- cpu 分析示例 end---
	
	var count int
	for i := 0; i < 10000; i++ {
		count++
	}
	
	// --- 内存 分析示例 start---
	fm, err := os.Create("./memory.prof")
	if err != nil {
		fmt.Println("create memory.prof err:", err.Error())
		return
	}
	defer fm.Close()
	
	// 开始分析内存
	err = pprof.WriteHeapProfile(fm)
	if err != nil {
		fmt.Println("write heap prof err:", err.Error())
		return
	}
	// --- 内存 分析示例 end---
	
	for i := 0; i < 10000; i++ {
		count++
	}
	fmt.Println("do finish......count:", count)
}
