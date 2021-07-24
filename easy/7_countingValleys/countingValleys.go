package countingValleys

var path_value = map[rune]int32 {
    'U': 1,
    'D': -1,
}

func countingValleys(steps int32, path string) int32 {
    var p_level, level, valley int32
    for _, p := range path {
        p_level = level
        level += path_value[p]

        if level < 0 && p_level >= 0 {
            valley += 1
        }
    }

    return valley
}