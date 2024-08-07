package main

import (
	"fmt"
	"math"
)

// Coord 定义二维坐标
type Coord struct {
	X float64
	Y float64
}

// Polygon 定义多边形结构
type PolygonO struct {
	ID     int
	Points []Coord
}

// OctreeNode 定义八叉树节点
type OctreeNode struct {
	Bounds   Rect          // 节点边界
	Polygons []PolygonO    // 节点内的多边形
	Children []*OctreeNode // 子节点
}

// Rect 定义矩形边界
type Rect struct {
	MinX, MinY float64
	MaxX, MaxY float64
}

func testOctree() {
	// 假设有很多多边形和一个输入点
	polygons := []PolygonO{
		{ID: 1, Points: []Coord{{0, 0}, {5, 0}, {5, 5}, {0, 5}}},
		{ID: 2, Points: []Coord{{1, 1}, {4, 1}, {4, 4}, {1, 4}}},
		{ID: 3, Points: []Coord{{2, 2}, {3, 2}, {3, 3}, {2, 3}}},
		// 更多多边形...
	}
	
	// 创建八叉树根节点
	rootBounds := Rect{MinX: 0, MinY: 0, MaxX: 100, MaxY: 100}
	octree := NewOctreeNode(rootBounds)
	
	// 将所有多边形插入八叉树
	for _, polygon := range polygons {
		octree.Insert(polygon)
	}
	
	// 假设有一个输入点
	inputPoint := Coord{3, 3}
	
	// 查询包含输入点的多边形
	containingPolygon := octree.QueryPoint(inputPoint)
	
	if containingPolygon != nil {
		fmt.Printf("Point (%.2f, %.2f) is inside Polygon %d\n", inputPoint.X, inputPoint.Y, containingPolygon.ID)
	} else {
		fmt.Println("Point is not inside any Polygon")
	}
}

// 创建新的八叉树节点
func NewOctreeNode(bounds Rect) *OctreeNode {
	return &OctreeNode{
		Bounds:   bounds,
		Polygons: make([]PolygonO, 0),
		Children: make([]*OctreeNode, 0),
	}
}

// 向八叉树节点插入多边形
func (node *OctreeNode) Insert(polygon PolygonO) {
	if len(node.Children) > 0 {
		// 如果有子节点，将多边形插入到适当的子节点中
		for i := 0; i < 8; i++ {
			if node.Children[i] != nil && node.Children[i].Bounds.containsPolygon(polygon) {
				node.Children[i].Insert(polygon)
				return
			}
		}
	}
	// 否则，将多边形插入到当前节点
	node.Polygons = append(node.Polygons, polygon)
}

// 检查节点是否包含多边形
func (rect Rect) containsPolygon(polygon PolygonO) bool {
	for _, point := range polygon.Points {
		if point.X < rect.MinX || point.X > rect.MaxX || point.Y < rect.MinY || point.Y > rect.MaxY {
			return false
		}
	}
	return true
}

// QueryPoint 查询包含给定点的多边形
func (node *OctreeNode) QueryPoint(point Coord) *PolygonO {
	if len(node.Children) > 0 {
		// 如果有子节点，递归查询
		for _, child := range node.Children {
			if child.Bounds.containsPoint(point) {
				return child.QueryPoint(point)
			}
		}
	} else {
		// 否则，检查当前节点中的多边形
		for _, polygon := range node.Polygons {
			if pointInPolygon(point, polygon.Points) {
				return &polygon
			}
		}
	}
	return nil // 找不到包含点的多边形
}

// 判断点是否在多边形内部，射线法判断
func pointInPolygon(point Coord, polygon []Coord) bool {
	intersections := 0
	for i := range polygon {
		j := (i + 1) % len(polygon)
		pi := polygon[i]
		pj := polygon[j]
		
		if pi.Y == pj.Y {
			continue // 水平边，跳过
		}
		
		if point.Y < math.Min(pi.Y, pj.Y) || point.Y >= math.Max(pi.Y, pj.Y) {
			continue // 不相交的边，跳过
		}
		
		// 计算射线与边的交点的 X 坐标
		x := (point.Y-pi.Y)*(pj.X-pi.X)/(pj.Y-pi.Y) + pi.X
		
		if x > point.X {
			intersections++ // 有奇数个交点，点在多边形内部
		}
	}
	
	return intersections%2 == 1
}

// 检查节点是否包含点
func (rect Rect) containsPoint(point Coord) bool {
	return point.X >= rect.MinX && point.X <= rect.MaxX && point.Y >= rect.MinY && point.Y <= rect.MaxY
}
