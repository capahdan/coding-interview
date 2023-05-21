package faktorial

func Faktorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	return num * Faktorial(num-1)
}
