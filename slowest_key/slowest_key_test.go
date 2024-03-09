package slowest_key

// import "fmt"

// func SlowestKey(arr [][]int) string {

// 	fmt.Println()

// }

// func slowestKey(matrix [][]int) string {
// 	slowestTimes := 0
// 	slowestChar := 0

// 	charMap := map[int]string{
// 		0:  "a",
// 		1:  "b",
// 		2:  "c",
// 		3:  "d",
// 		4:  "e",
// 		5:  "f",
// 		6:  "g",
// 		7:  "h",
// 		8:  "i",
// 		9:  "j",
// 		10: "k",
// 		11: "l",
// 		12: "m",
// 		13: "n",
// 		14: "o",
// 		15: "p",
// 		16: "q",
// 		17: "r",
// 		18: "s",
// 		19: "t",
// 		20: "u",
// 		21: "v",
// 		22: "w",
// 		23: "x",
// 		24: "y",
// 		25: "z",
// 	}

// 	for i := 0; i < len(matrix); i++ {
// 		if i > 0 {
// 			if matrix[i][1]-matrix[i-1][1] > slowestTimes {
// 				slowestTimes = matrix[i][1] - matrix[i-1][1]
// 				slowestChar = matrix[i][0]
// 			}
// 		} else {
// 			if matrix[i][1] > slowestTimes {
// 				slowestTimes = matrix[i][1]
// 				slowestChar = matrix[i][0]
// 			}
// 		}
// 	}

// 	return substitusiMap[slowestChar]
// }

// func main() {
// 	fmt.Println(slowestKey([][]int{{0, 2}, {1, 5}, {0, 9}, BadExpr}))
// }
