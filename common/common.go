package common

func IsTwoSlicesEqual(a, b []int32) bool {
    for i, _ := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}