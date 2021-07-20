package breakingtheRecords

func breakingtheRecords(scores []int32) []int32 {
    min := scores[0]
    max := scores[0]
    records := []int32{0, 0}
    for _, score := range scores {
        if score < min {
            min = score
            records[1] += 1
        } else if score > max {
            max = score
            records[0] += 1
        }
    }
    return records
}