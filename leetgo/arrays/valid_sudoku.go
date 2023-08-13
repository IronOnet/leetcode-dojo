package arrays 


func IsValidSudoku(board [][]byte) bool{
	cols := make(map[byte]bool) 
	rows := make(map[byte]bool) 
	squares := make(map[byte]bool) 

	for r:= 0; r <= 9;  r++{
		for c:= 0; c <= 9; c++{
			if board[r][c] == '.'{
				continue
			}

			_, exists_rows := rows[board[r][c]] 
			_, exists_cols := cols[board[r][c]] 
			start_row, start_col := (r/3)*3, (c/3)*3
			_, exists_squares := squares[board[(start_row+r)%3][(start_col + c)%3]]

			if exists_rows || exists_cols || exists_squares{
				return false 
			}

			cols[board[r][c]] = true 
			rows[board[r][c]] = true 
			squares[board[(start_row + r)%3][(start_col + c)%3]] = false 
		}

		
	}
	return true 
}