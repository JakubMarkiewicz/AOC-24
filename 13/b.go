package main

func B(games []Game) int {
	res := 0

	for _, g := range games {
		res += countGameB(g)
	}

	return res
}

// TODO: determine which button should be pressed more times based on price
func countGameB(g Game) int {
	g.P[0] += 10000000000000
	g.P[1] += 10000000000000

	d := g.A[0]*g.B[1] - g.B[0]*g.A[1]
	dx := g.P[0]*g.B[1] - g.P[1]*g.B[0]
	dy := g.A[0]*g.P[1] - g.P[0]*g.A[1]

	x := dx / d
	y := dy / d

	xRes := g.A[0]*x+g.B[0]*y == g.P[0]
	yRes := g.A[1]*x+g.B[1]*y == g.P[1]

	if xRes && yRes {
		return x*3 + y

	}

	return 0

}
