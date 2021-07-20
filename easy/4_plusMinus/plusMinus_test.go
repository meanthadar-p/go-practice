package plusMinus

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input       []int32
    expected    string
}{
    {[]int32{1, 1, 0, -1, -1}, "0.400000, 0.400000, 0.200000"},
    {[]int32{-4, 3, -9, 0, 4, 1}, "0.500000, 0.333333, 0.166667"},
}

func TestPlusMinus(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %v when input %v", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := plusMinus(test.input); r != test.expected {
                t.Errorf("expected %s but got %s", test.expected, r)
            }
        })
    }
}