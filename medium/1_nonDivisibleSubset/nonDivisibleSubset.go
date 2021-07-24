package nonDivisibleSubset

import "math"

func nonDivisibleSubset(k int32, s []int32) int32 {
    var cnt []int32
    var result int32
    halfK := int(math.Ceil(float64(k)/2.0))

    for i := 0 ; i <= int(k) ; i++ {
        cnt = append(cnt, 0)
    }

    for _, value := range s {
        m :=  value % k
        cnt[m] += 1
    }

    for i, value := range cnt[1:halfK] {
        pair := cnt[len(cnt)-i-2]
        if value >= pair {
            result += value
        }else {
            result += pair
        }
    }

    if cnt[0] > 0 {
        result += 1
    }
    if k % 2 == 0 && cnt[k/2] > 0 {
        result += 1
    }

    return result
}
