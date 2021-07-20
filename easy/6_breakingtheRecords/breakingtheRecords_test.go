package breakingtheRecords

import (
    "fmt"
    "testing"
    "go-practice/common"
)

var tests = []struct {
    input       []int32
    expected    []int32
}{
    {[]int32{12, 24, 10, 24}, []int32{1, 1}},
    {[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, []int32{2, 4}},
    {[]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}, []int32{4, 0}},
}

func TestSum(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %v when input %v", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := breakingtheRecords(test.input); !common.IsTwoSlicesEqual(r, test.expected) {
                t.Errorf("expected %v but got %v", test.expected, r)
            }
        })
    }
}