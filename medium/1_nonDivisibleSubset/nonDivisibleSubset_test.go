package nonDivisibleSubset

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input_k     int32
    input_s     []int32
    expected    int32
}{
    {3, []int32 {1, 7, 2, 4}, 3},
    {7, []int32 {278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}, 11},
}

func TestNonDivisibleSubset(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %d when input k:%v s:%v", test.expected, test.input_k, test.input_s)
        t.Run(name, func(t *testing.T){
            if r := nonDivisibleSubset(test.input_k, test.input_s); r != test.expected {
                t.Errorf("expected %d but got %d", test.expected, r)
            }
        })
    }
}