package valid_sudoku

import "strconv"

func IsValidSudoku(board [][]byte) bool {
	rowsMap := [9][9]bool{}
	colsMap := [9][9]bool{}
	gridMap := [9][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val, err := strconv.Atoi(string(board[row][col]))
			if err != nil {
				continue
			}
			val--
			gridIndex := col/3 + (row/3)*3
			if rowsMap[row][val] || colsMap[col][val] || gridMap[gridIndex][val] {
				return false
			}
			rowsMap[row][val] = true
			colsMap[col][val] = true
			gridMap[gridIndex][val] = true
		}
	}
	return true
}
