package plusMinus

import "fmt"

func plusMinus(arr []int32) string {
    var positive, negative, zero, size float64
    size = float64(len(arr))
    for _, num := range arr {
        if num > 0 {
            positive += 1
        }else if num < 0 {
            negative += 1
        }else {
            zero += 1
        }
    }

    return fmt.Sprintf("%.6f, %.6f, %.6f", positive/size, negative/size, zero/size)
}