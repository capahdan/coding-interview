package sewa

func Sewa(jmasuk, jkeluar int) int {
	if jmasuk < 6 || jkeluar > 23 {
		return 0
	}
	if jkeluar-jmasuk == 1 {
		return 350000
	}
	if jkeluar-jmasuk == 2 {
		return 500000
	}

	if jkeluar-jmasuk > 2 && jkeluar-jmasuk < 8 {
		return 150000*(jkeluar-jmasuk) - 100000
	}
	if jkeluar-jmasuk > 8 {
		return 200000*(jkeluar-jmasuk) - 200000
	}
	return 0
}
