package solve

func Solve(board [][]byte) [][]byte {
	if len(board) == 0 {
		return [][]byte{}
	}

	m, n := len(board), len(board[0])

	// Fungsi bantu untuk melakukan pencarian DFS dari sebuah sel
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != 'O' {
			return
		}
		board[row][col] = 'T' // Tanda sementara untuk sel "O" yang dikunjungi
		// Jelajahi keempat sel tetangga
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	// Tandai "O" yang terhubung dengan pinggiran sebagai "T"
	for i := 0; i < m; i++ {
		dfs(i, 0)   // Pinggiran kiri
		dfs(i, n-1) // Pinggiran kanan
	}
	for j := 0; j < n; j++ {
		dfs(0, j)   // Pinggiran atas
		dfs(m-1, j) // Pinggiran bawah
	}

	// Balikkan "O" menjadi "X" dan kembalikan "T" menjadi "O"
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}

	return board
}

// jika bord[i][j] di ujung maka gak usah di apa-apain
// jika board disampingnya adalah 0 maka gak usah juga di apa-apain

//      0,    1,    2,     3   i
// 0    00,  01,  0,2    0,3
// 1    10,  11,  1,2    1,3
// 2    20,  21,  2,2    2,3
// 3    30,  31,  3,2    3,3

// j
