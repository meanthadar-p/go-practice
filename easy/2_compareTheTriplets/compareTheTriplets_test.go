package compareTheTriplets

import (
    "fmt"
    "testing"
    "go-practice/common"
)

var tests = []struct{
    input_a     []int32
    input_b     []int32
    expected    []int32
}{
    {[]int32{5, 6, 7}, []int32{3, 6, 10}, []int32{1, 1}},
    {[]int32{17, 28, 30}, []int32{99, 16, 8}, []int32{2, 1}},
}

func TestCompare(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %#v when input %#v, %#v", test.expected, test.input_a, test.input_b)
        t.Run(name, func(t *testing.T){
            if r := compareTheTriplets(test.input_a, test.input_b); !common.IsTwoSlicesEqual(r, test.expected) {
                t.Errorf("expected %#v but got %#v", test.expected, r)
            }
        })
    }
}