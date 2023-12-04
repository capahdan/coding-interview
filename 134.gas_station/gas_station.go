package gas_station

func CanCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	totalSurplus := 0
	surplus := 0
	start := 0

	for i := 0; i < n; i++ {
		totalSurplus += gas[i] - cost[i]
		surplus += gas[i] - cost[i]
		if surplus < 0 {
			surplus = 0
			start = i + 1
		}
	}

	if totalSurplus < 0 {
		return -1
	}

	return start
}

// yang perlu kita lakukan adalah
// menyimpan sisa gas yang ada kalo ada
// simpan index yang bisa sirkular
//
