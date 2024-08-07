package main

//
// // Region 定义区域的结构体
// type RegionOrb struct {
// 	ID       int
// 	Geometry orb.Geometry
// }
//
// func paulmachOrb() {
// 	// 示例：假设有100万个区域
// 	regions := make([]RegionOrb, 0)
//
// 	// 假设从数据库或文件加载这些区域数据，这里简化为手动创建示例数据
// 	regions = append(regions, RegionOrb{
// 		ID:       1,
// 		Geometry: orb.Polygon{{{0, 0}, {0, 3}, {3, 3}, {3, 0}, {0, 0}}},
// 	})
// 	// 添加更多区域数据，实际应用中从数据库或文件加载
//
// 	// 构建 四叉树
// 	quadtree.New(orb.Bound{
// 		Min: orb.Point{2, 5},
// 		Max: orb.Point{5, 8},
// 	})
// 	tree := quadtree.NewTree(25, 50) // 参数可以根据数据量进行调整
//
// 	// 将所有区域添加到 R-tree 中
// 	for _, region := range regions {
// 		tree.Insert(region.Geometry.Bound(), region)
// 	}
//
// 	// 需要查询的点
// 	point := orb.Point{0, 0}
//
// 	// 查找包含给定点的区域
// 	var foundRegion *RegionOrb
// 	tree.Search(point.Bound(), func(item rtree.Item) bool {
// 		region := item.(*RegionOrb)
// 		if planar.Contains(region.Geometry, point) {
// 			foundRegion = region
// 			return true // 找到即可退出搜索
// 		}
// 		return false
// 	})
//
// 	// 输出结果
// 	if foundRegion != nil {
// 		fmt.Printf("Point is in region with ID: %d\n", foundRegion.ID)
// 	} else {
// 		fmt.Println("Point is not within any region.")
// 	}
// }
//
// func paulmachOrb2() {
// 	// 假设这是你的多边形地理数据，这里示例一个简单的多边形
// 	polygonCoords := []orb.Point{
// 		{100.0, 0.0},
// 		{101.0, 0.0},
// 		{101.0, 1.0},
// 		{100.0, 1.0},
// 	}
// 	geo.
// 	// 假设这是你的点的坐标
// 	point := orb.Point{100.5, 0.5}
//
// 	// 创建一个多边形
// 	polygon := orb.Polygon{polygonCoords}
//
// 	// 判断点是否在多边形内部
// 	if geo.PointInPolygon(point, polygon) {
// 		fmt.Println("点位于多边形内部")
// 	} else {
// 		fmt.Println("点不位于多边形内部")
// 	}
//
// }
