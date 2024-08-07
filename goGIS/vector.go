package main

import "fmt"

func VectorTest() {
	// 多边形的顶点坐标
	polygon := []*PointR{{0, 0}, {0, 1158}, {346, 1345}, {750, 1118}, {750, 0}}
	
	bSuc1 := IsPointInConvexPolygon(polygon, &PointR{350, 1160})
	fmt.Println("bSuc1=", bSuc1)
	
	bSuc2 := IsPointInConvexPolygon(polygon, &PointR{2, 1190})
	fmt.Println("bSuc2=", bSuc2)
	
	bSuc3 := IsPointInConvexPolygon(polygon, &PointR{749, 1118})
	fmt.Println("bSuc3=", bSuc3)
}

// point1和point2的向量
func SubPoint(point1 *PointR, point2 *PointR) *PointR {
	return &PointR{
		X: point1.GetX() - point2.GetX(),
		Y: point1.GetY() - point2.GetY(),
	}
}

// 向量积（叉乘）
func CrossProduct(point1 *PointR, point2 *PointR) float64 {
	return point1.GetX()*point2.GetY() - point2.GetX()*point1.GetY()
}

// 判断一个点是否在多边形内
func IsPointInConvexPolygon(aPoints []*PointR, vTarget *PointR) bool {
	if len(aPoints) == 0 {
		return false
	}
	
	var nCurCrossProduct float64
	var nLastValue float64
	
	for i := 0; i < len(aPoints); i++ {
		vU := SubPoint(vTarget, aPoints[i])
		nNextIndex := (i + 1) % len(aPoints)
		vV := SubPoint(aPoints[nNextIndex], aPoints[i])
		nCurCrossProduct = CrossProduct(vU, vV)
		if i > 0 && nCurCrossProduct*nLastValue <= 0 {
			return false
		}
		nLastValue = nCurCrossProduct
	}
	
	return true
}
