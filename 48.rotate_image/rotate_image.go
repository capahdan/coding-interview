package rotate

// Create a function pow that takes two argument, a and b. And then it will calculate the a to the power of b. Please do not use built in function to calculate the power

func RotateImage(matrix [][]int) [][]int {

	var stackMatrix []int

	for i := range matrix {
		for j := range matrix[i] {
			stackMatrix = append(stackMatrix, matrix[i][j])
		}
	}

	for i := range matrix {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			matrix[j][i] = stackMatrix[len(stackMatrix)-1]
			stackMatrix = stackMatrix[:len(stackMatrix)-1]
		}
	}

	return matrix

}

//  00, 01, 02
// 10, 11, 12
// 20, 21, 22

// 1 2 3
// 4 5 6
// 7 8 9

// 7 4 1
// 8 5 2
// 9 6 3

//  00 -> 02
//  01 -> 12
//  02 -> 22

// 10 -> 01
// 11 -> 11
// 12 -> 21

// 20 -> 00
// 21 -> 10
// 22 ->
