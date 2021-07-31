package common

import (
    "fmt"
    "testing"
)

var tests = []struct{
    input     []int32
    expected    []int32
}{
    {[]int32{3, 2, 5, 4, 1}, []int32{1, 2, 3, 4, 5}},
    {[]int32{3, 2, 2, 1, 4}, []int32{1, 2, 2, 3, 4}},
}

func TestCompare(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %#v when input %#v", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := Quicksort(test.input); !IsTwoSlicesEqual(r, test.expected) {
                t.Errorf("expected %#v but got %#v", test.expected, r)
            }
        })
    }
}