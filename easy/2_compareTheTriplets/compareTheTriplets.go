package compareTheTriplets

func compareTheTriplets(a []int32, b []int32) []int32 {
    scores := []int32{0, 0}
    for i, _ := range a {
        if a[i] > b[i] {
            scores[0] += 1
        } else if a[i] < b[i] {
            scores[1] += 1
        }
    }
    return scores
}