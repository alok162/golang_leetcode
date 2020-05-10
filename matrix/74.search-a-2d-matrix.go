// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.

// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 3
// Output: true

// Input:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 13
// Output: false

func searchMatrix(matrix [][]int, target int) bool {
    
    if len(matrix) == 0 { return false }
    i, j := 0, len(matrix[0])-1
    
    for i < len(matrix) && j >= 0 {
        if matrix[i][j] == target {
            return true
        }
        if matrix[i][j] > target {
            j --
        } else {
            i ++
        }
    }
    return false
}
