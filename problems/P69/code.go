package P69

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	y := x / 2
	for y*y-x > 0 {
		y = (y + x/y) / 2
	}
	return int(y)
}
