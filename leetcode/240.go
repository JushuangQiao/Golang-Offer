package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	row, col := rows-1, 0
	for row >= 0 && col < cols {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row -= 1
		} else {
			col += 1
		}
	}
	return false
}
