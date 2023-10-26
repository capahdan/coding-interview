package sqrt

func MySqrt(x int) int {
	////this is my bruteforce solution
	//i := 1
	//for x >= i*i {
	//	i++
	//}
	//
	//return i - 1

	var r int64
	r = int64(x)

	for r*r > int64(x) {
		r = (r + int64(x)/r) / 2
	}
	return int(r)
}
