package diagonalDifference

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input       [][]int32
    expected    int32
}{
    {[][]int32{[]int32{1, 2, 3}, []int32{4, 5, 6}, []int32{9, 8, 9}}, 2},
    {[][]int32{[]int32{11, 2, 4}, []int32{4, 5, 6}, []int32{10, 8, -12}}, 15},
}

func TestDiagonalDifference(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %d when input %v", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := diagonalDifference(test.input); r != test.expected {
                t.Errorf("expected %d but got %d", test.expected, r)
            }
        })
    }
}