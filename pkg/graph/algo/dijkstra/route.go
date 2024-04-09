// Package dijkstra
// @Description: 路径

package dijkstra

import "fmt"

// HasShortestRoute
//
//	@Description: 是否存在从 source 到 v的路径
//	@receiver d
//	@param v
//	@return bool
func (d *dijkstra) HasShortestRoute(dest string) bool {
	if !d.g.Contains(dest) {
		fmt.Println("不存在顶点\n", dest)
		return false
	}
	// _, find := d.distToSource[source]
	find := d.marked[dest]
	if !find {
		source := *d.source
		fmt.Printf("从 %s 到 %s 不存在路径\n", source, dest)
	}
	return find
}

// ShortestRoute
//
//	@Description: 路劲
//	@receiver d
//	@param dest
//	@return []string 顶点
//	@return []float64 权重
//	@return float64 最短距离
//	@return bool 是否存在
func (d *dijkstra) ShortestRoute(dest string) ([]string, []float64, float64, bool) {
	if !d.HasShortestRoute(dest) {
		return nil, nil, -1, false
	}

	source := *d.source
	if dest == source {
		return []string{dest}, []float64{0}, 0, true
	}

	route0, weights0 := make([]string, 0), make([]float64, 0)

	to := dest
	for from, has := d.whereFrom[to]; has; {
		route0 = append(route0, from)
		weight, _ := d.g.GetEdgeWeight(from, to)
		weights0 = append(weights0, weight)

		to = from
		from, has = d.whereFrom[to]
	}

	// reverse
	route, weights := make([]string, len(route0)), make([]float64, len(weights0))
	size := len(route0)
	for i := 0; i < size; i++ {
		route[size-i-1], weights[size-i-1] = route0[i], weights0[i]
	}
	return route, weights, d.distToSource[dest], true
}

// PrintShortestRoute
//
//	@Description: 路径打印
//	@receiver d
//	@param v
func (d *dijkstra) PrintShortestRoute(dest string) {
	if !d.HasShortestRoute(dest) {
		return
	}
	source := *d.source
	route, weights, distance, _ := d.ShortestRoute(dest)
	fmt.Printf("从 %s 到 %s 的最短距离为 : %g\n", source, dest, distance)

	fmt.Print("路径: ")

	for i, v := range route {
		fmt.Printf("%s -(%g)-> ", v, weights[i])
	}
	fmt.Println(dest)
	fmt.Println()
}
