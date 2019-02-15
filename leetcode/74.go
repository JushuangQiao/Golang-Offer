package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	for rows > 0 {
		if matrix[rows-1][0] == target {
			return true
		} else if matrix[rows-1][0] < target {
			return binSearch(matrix[rows-1], target)
		} else {
			rows -= 1
		}
	}
	return false
}

func binSearch(data []int, target int) bool {
	low, high := 1, len(data)
	for low < high {
		mid := low + (high-low)/2
		if data[mid] == target {
			return true
		} else if data[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return false
}

func main() {
	tests := [][]int{{}, {}}
	_ := searchMatrix(tests, 1)
}
