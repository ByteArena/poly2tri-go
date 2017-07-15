package poly2tri

type Edge struct {
	P *Point
	Q *Point
}

func NewEdge(p1, p2 *Point) *Edge {

	actualP := p1
	actualQ := p2

	ycomp := XYCompareFloat(p1.Y, p2.Y)

	if ycomp > 0 {
		actualP = p2
		actualQ = p1
	} else if ycomp == 0 {
		xcomp := XYCompareFloat(p1.X, p2.X)
		if xcomp > 0 {
			actualP = p2
			actualQ = p1
		} else if xcomp == 0 {
			panic("poly2tri Invalid Edge constructor: repeated points!")
		}
	}

	res := &Edge{
		P: actualP,
		Q: actualQ,
	}

	res.P.Edges = append(res.P.Edges, res)
	return res
}
