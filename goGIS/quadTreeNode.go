package main

import "fmt"

type PointQ struct {
	X, Y float64
}

type Polygon []PointQ

type QuadTreeNode struct {
	boundaryX, boundaryY float64
	subNodes             [4]*QuadTreeNode
	polygons             []Polygon
}

func quaTreeTest() {
	polygon1 := Polygon{
		{X: 0, Y: 0},
		{X: 8, Y: 1},
		{X: 4, Y: 7},
		{X: -2, Y: 6},
		{X: 1, Y: 9},
	}
	
	// 创建一个四叉树节点，并设置其边界（这里简单设置为整个平面的一半）
	rootNode := NewQuadTreeNode(0, 0, 100) // 假设整个平面是从(0,0)到(10,10)
	
	// 将多边形插入到四叉树中
	rootNode.Insert(polygon1)
	
	// 判断点是否在多边形内部
	testPoint1 := PointQ{X: 2, Y: 2}
	fmt.Println("Point (2, 2) is inside polygon1:", rootNode.Contains(testPoint1)) // 应该输出true
	
	testPoint2 := PointQ{X: 5, Y: 5}
	fmt.Println("Point (5, 5) is inside polygon1:", rootNode.Contains(testPoint2)) // 应该输出true
	
	testPoint3 := PointQ{X: 6, Y: 6}
	fmt.Println("Point (6, 6) is inside polygon1:", rootNode.Contains(testPoint3)) // 应该输出false，因为点在多边形外部
	
}

func NewQuadTreeNode(x, y, width float64) *QuadTreeNode {
	return &QuadTreeNode{
		boundaryX: x + width/2,
		boundaryY: y + width/2,
	}
}

func (node *QuadTreeNode) Split() {
	half := 0.5 * node.boundaryX
	node.subNodes[0] = NewQuadTreeNode(node.boundaryX, node.boundaryY, half)
	node.subNodes[1] = NewQuadTreeNode(half, node.boundaryY, half)
	node.subNodes[2] = NewQuadTreeNode(node.boundaryX, half, half)
	node.subNodes[3] = NewQuadTreeNode(half, half, half)
}

func (node *QuadTreeNode) Insert(polygon Polygon) {
	if node.subNodes[0] != nil {
		index := node.getIndex(polygon[0])
		if index != -1 {
			node.subNodes[index].Insert(polygon)
			return
		}
	}
	node.polygons = append(node.polygons, polygon)
	if len(node.polygons) > 1 {
		node.Split()
		for i, subNode := range node.subNodes {
			for j := range node.polygons {
				index := node.getIndex(node.polygons[j][0])
				if index == i || index == -1 {
					subNode.Insert(node.polygons[j])
				}
			}
		}
		node.polygons = nil
	}
}

func (node *QuadTreeNode) getIndex(p PointQ) int {
	if p.X < node.boundaryX && p.Y < node.boundaryY {
		return 0
	} else if p.X >= node.boundaryX && p.Y < node.boundaryY {
		return 1
	} else if p.X < node.boundaryX && p.Y >= node.boundaryY {
		return 2
	} else if p.X >= node.boundaryX && p.Y >= node.boundaryY {
		return 3
	}
	return -1
}

func (node *QuadTreeNode) Contains(p PointQ) bool {
	if node.subNodes[0] != nil {
		index := node.getIndex(p)
		if index != -1 {
			return node.subNodes[index].Contains(p)
		}
	}
	for _, polygon := range node.polygons {
		inside := true
		j := len(polygon) - 1
		for i := 0; i < len(polygon); i++ {
			pi, pj := polygon[i], polygon[j]
			if (pi.Y > p.Y) != (pj.Y > p.Y) &&
				(p.X < (pj.X-pi.X)*(p.Y-pi.Y)/(pj.Y-pi.Y)+pi.X) {
				inside = !inside
			}
			j = i
		}
		if inside {
			return true
		}
	}
	return false
}
