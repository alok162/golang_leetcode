// Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

// Input: 
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// Output: 
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

func setZeroes(matrix [][]int)  {
    var colFlag bool
    var rowFlag bool
    for i := 0; i < len(matrix); i ++ {
        for j := 0; j < len(matrix[0]); j ++ {
            if i == 0 && matrix[0][j] == 0 {
                rowFlag = true
            }
            if j == 0 && matrix[i][0] == 0 {
                colFlag = true
            }
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }
    
     for i := 1; i < len(matrix); i ++ {
            for j := 1; j < len(matrix[0]); j ++ {
                if matrix[0][j] == 0 {
                    matrix[i][j] = 0
                }
                if matrix[i][0] == 0 {
                    matrix[i][j] = 0
                }
            }
        }
    
        if colFlag {
            for i := 0; i < len(matrix) ; i ++ {
                matrix[i][0] = 0
            }
        }
    
        if rowFlag {
            for i := 0; i < len(matrix[0]) ; i ++ {
                matrix[0][i] = 0
            }
        }
}
