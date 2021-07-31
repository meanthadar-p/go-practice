package organizingContainers

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input       [][]int32
    expected    string
}{
    {[][]int32 {[]int32{1, 4}, []int32{2, 3}}, "Impossible"},
    {[][]int32 {[]int32{1, 3, 1}, []int32{2, 1, 2}, []int32{3, 3, 3}}, "Impossible"},
    {[][]int32 {[]int32{0, 2, 1}, []int32{1, 1, 1}, []int32{2, 0, 0}}, "Possible"},
}

func TestOrganizingContainers(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %s when input %v\n", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := organizingContainers(test.input); r != test.expected {
                t.Errorf("expected %s but got %s\n", test.expected, r)
            }
        })
    }
}