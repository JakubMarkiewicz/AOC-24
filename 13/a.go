package main

func A(games []Game) int {
	res := 0

	for _, g := range games {
		res += countGameA(g)
	}

	return res
}

// TODO: determine which button should be pressed more times based on price
func countGameA(g Game) int {

	for a := g.P[0] / g.A[0]; a > 0; a-- {

		lMoveX, lMoveY := g.A[0]*a, g.A[1]*a
		xLeft, yLeft := g.P[0]-lMoveX, g.P[1]-lMoveY

		b := xLeft / g.B[0]

		if g.B[1]*b == yLeft {
			return a*3 + b*1
		}
	}

	return 0

}
