package rectangles

type Point struct{ x, y int }

func Count(diagram []string) int {
	if len(diagram) == 0 || len(diagram[0]) == 0 {
		return 0
	}

	isVertex := func(x, y int) bool { return diagram[y][x] == '+' }
	isEdgeH := func(x, y int) bool { return isVertex(x, y) || diagram[y][x] == '-' }
	isEdgeV := func(x, y int) bool { return isVertex(x, y) || diagram[y][x] == '|' }

	vertices := make([]Point, 0, (len(diagram)*len(diagram[0]))/2)
	for y := range diagram {
		for x := range []byte(diagram[y]) {
			if isVertex(x, y) {
				vertices = append(vertices, Point{x, y})
			}
		}
	}

	count := 0
	for i, p := range vertices {
	qLoop:
		for _, q := range vertices[i+1:] {
			if p.y >= q.y || p.x >= q.x || !isVertex(p.x, q.y) || !isVertex(q.x, p.y) {
				continue
			}
			for x := p.x + 1; x < q.x; x++ {
				if !isEdgeH(x, p.y) || !isEdgeH(x, q.y) {
					continue qLoop
				}
			}
			for y := p.y + 1; y < q.y; y++ {
				if !isEdgeV(p.x, y) || !isEdgeV(q.x, y) {
					continue qLoop
				}
			}
			count++
		}
	}
	return count
}
