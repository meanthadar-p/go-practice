package diagonalDifference

import (
    "fmt"
    "math"
)

func diagonalDifference(arr [][]int32) int32 {
    var primary, secondary int32
    for r, row := range arr {
        for c, value := range row {
            if r == c {
                primary += value
            }
            if c == len(row) - r - 1 {
                secondary += value
            }
        }
    }
    return int32(math.Abs(float64(primary - secondary)))
}