package search2d

func SearchMatrix(matrix [][]int, target int) bool {

	n := len(matrix)
	m := len(matrix[0])
	l := 0
	r := m*n - 1
	for l != r {
		mid := (l + r - 1) >> 1
		if matrix[mid/m][mid%m] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return matrix[r/m][r%m] == target
}
