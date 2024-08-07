package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type PointR struct {
	X float64
	Y float64
}

func (p *PointR) GetX() float64 { return p.X }
func (p *PointR) GetY() float64 { return p.Y }

func rayStart() {
	fmt.Println("创建示例数据：  ", time.Now().Format(TIMEFORMAT))
	polygons := createPolygon()
	point := PointR{X: 123, Y: 85}
	var sums [][]PointR
	
	fmt.Println("创建示例数据完成，开始筛选：  ", time.Now().Format(TIMEFORMAT))
	for _, polygon := range polygons {
		if inPolygon := IfPointsInPolygon(point, polygon); inPolygon {
			sums = append(sums, polygon)
		}
	}
	fmt.Println("筛选完成：  ", time.Now().Format(TIMEFORMAT))
	fmt.Println(sums)
}

func IfPointsInPolygon(point PointR, area []PointR) bool {
	// 目标点的x, y坐标
	x := point.X
	y := point.Y
	
	// 多边形的点数
	count := len(area)
	
	// 点是否在多边形中
	var inInside bool
	
	// 浮点类型计算与0的容差
	precision := 2e-10
	
	// 依次计算每条边，根据每边两端点和目标点的状态栏判断
	for i, j := 0, count-1; i < count; j, i = i, i+1 {
		// 记录每条边上的两个点坐标
		x1 := area[i].X
		y1 := area[i].Y
		x2 := area[j].X
		y2 := area[j].Y
		
		// 判断点与多边形顶点是否重合
		if (x1 == x && y1 == y) || (x2 == x && y2 == y) {
			return true
		}
		
		// 判断点是否在水平直线上
		if (y == y1) && (y == y2) {
			return true
		}
		
		// 判断线段两端点是否在射线两侧
		if (y >= y1 && y < y2) || (y < y1 && y >= y2) {
			// 斜率
			k := (x2 - x1) / (y2 - y1)
			
			// 相交点的 x 坐标
			_x := x1 + k*(y-y1)
			
			// 点在多边形的边上
			if _x == x {
				return true
			}
			
			// 浮点类型计算容差
			if math.Abs(_x-x) < precision {
				return true
			}
			
			// 射线平行于x轴，穿过多边形的边
			if _x > x {
				inInside = !inInside
			}
		}
	}
	
	return inInside
}

func createPolygon() [][]PointR {
	var areas [][]PointR
	for i := 0; i < 10000; i++ {
		x := rand.Float64()
		y := rand.Float64()
		height := rand.Float64()
		weight := rand.Float64()
		area := []PointR{{x, y}, {x + weight, y}, {x + height, y + height}, {x, y + weight}}
		areas = append(areas, area)
	}
	return areas
}
