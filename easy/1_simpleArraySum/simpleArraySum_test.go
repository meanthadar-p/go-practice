package simpleArraySum

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input       []int32
    expected    int32
}{
    {[]int32 {1, 2, 3, 4}, 10},
    {[]int32 {1, 2, 3, 4, 10, 11}, 31},
}

func TestSum(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %d when input %#v", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := simpleArraySum(test.input); r != test.expected {
                t.Errorf("expected %d but got %d", test.expected, r)
            }
        })
    }
}