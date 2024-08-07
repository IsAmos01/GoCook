package main

import (
	"fmt"
	"github.com/dhconnelly/rtreego"
	"math"
	"math/rand"
	"time"
)

var TIMEFORMAT = "2006-01-02 15:04:05.99999"

// Region 定义区域的结构体
type Region struct {
	ID           int
	PolygonBound *rtreego.Rect
}

func (r *Region) Bounds() rtreego.Rect {
	return *r.PolygonBound
}

func dhconnellyRtree() {
	// 示例：假设有100万个区域
	regions := make([]Region, 0)
	fmt.Println("加载数据：  ", time.Now().Format(TIMEFORMAT))
	// 假设从数据库或文件加载这些区域数据，这里简化为手动创建示例数据
	for i := 1; i <= 100000; i++ {
		// 随机生成区域的位置和大小
		x := rand.Float64() * 1000
		y := rand.Float64() * 1000
		width := rand.Float64() * 50
		height := rand.Float64() * 50
		newRectBound, err := boundingBox([][]float64{{x, y}, {x + width, y + height}, {x - width, y - height}})
		if err != nil {
			return
		}
		region := Region{
			ID:           i,
			PolygonBound: &newRectBound,
		}
		regions = append(regions, region)
	}
	// 添加更多区域数据，实际应用中从数据库或文件加载
	fmt.Println("加载完成，开始构建树：  ", time.Now().Format(TIMEFORMAT))
	// 构建 R-tree 空间索引
	tree := rtreego.NewTree(2, 25, 50) // 参数可以根据数据量进行调整
	
	// 将所有区域添加到 R-tree 中
	for _, region := range regions {
		tree.Insert(&region)
	}
	
	fmt.Println("构建树完成，开始查找：  ", time.Now().Format(TIMEFORMAT))
	// 需要查询的点
	point := rtreego.Point{100, 100}
	r, _ := rtreego.NewRect(point, []float64{10, 10})
	
	// 查找包含给定点的区域
	result := tree.SearchIntersect(r)
	
	// 输出结果
	if result != nil {
		for _, k := range result {
			fmt.Println(k.Bounds())
		}
		fmt.Println(result)
	} else {
		fmt.Println("Point is not within any region.")
	}
	
	fmt.Println("查找完成：  ", time.Now().Format(TIMEFORMAT))
}

func boundingBox(polygon [][]float64) (rtreego.Rect, error) {
	if len(polygon) == 0 {
		return rtreego.NewRect([]float64{0, 0}, []float64{0, 0})
	}
	
	minX, minY := polygon[0][0], polygon[0][1]
	maxX, maxY := polygon[0][0], polygon[0][1]
	
	// 遍历顶点，更新边界值
	for _, vertex := range polygon {
		minX = math.Min(minX, vertex[0])
		minY = math.Min(minY, vertex[1])
		maxX = math.Max(maxX, vertex[0])
		maxY = math.Max(maxY, vertex[1])
	}
	
	return rtreego.NewRectFromPoints([]float64{minX, minY}, []float64{maxX, maxY})
}
