package organizingContainers

import "go-practice/common"

const (
    IMPOSSIBLE = "Impossible"
    POSSIBLE   = "Possible"
)

func organizingContainers(containers [][]int32) string {
    var ballsAmount, containerSize []int32

    for i, container := range containers {
        for j, ball := range container {
            if len(ballsAmount)-1 < j {
                ballsAmount = append(ballsAmount, 0)
            }
            ballsAmount[j] += ball

            if len(containerSize)-1 < i {
                containerSize = append(containerSize, 0)
            }
            containerSize[i] += ball
        }
    }

    ballsAmount = common.Quicksort(ballsAmount)
    containerSize = common.Quicksort(containerSize)

    for i, ball := range ballsAmount {
        if ball != containerSize[i] {
            return IMPOSSIBLE
        }
    }

    return POSSIBLE
}