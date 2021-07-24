package countingValleys

import (
    "fmt"
    "testing"
)

var tests = []struct {
    steps       int32
    path        string
    expected    int32
}{
    {8, "DDUUUUDD", 1},
    {8, "UDDDUDUU", 1},
}

func TestSum(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %d when input %d, %s", test.expected, test.steps, test.path)
        t.Run(name, func(t *testing.T){
            if r := countingValleys(test.steps, test.path); r != test.expected {
                t.Errorf("expected %d but got %d", test.expected, r)
            }
        })
    }
}