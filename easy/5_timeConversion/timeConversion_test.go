package timeConversion

import (
    "fmt"
    "testing"
)

var tests = []struct {
    input       string
    expected    string
}{
    {"12:01:00PM", "12:01:00"},
    {"12:01:00AM", "00:01:00"},
    {"07:05:45PM", "19:05:45"},
}

func TestSum(t *testing.T){
    for _, test := range tests {
        name := fmt.Sprintf("should return %s when input %s", test.expected, test.input)
        t.Run(name, func(t *testing.T){
            if r := timeConversion(test.input); r != test.expected {
                t.Errorf("expected %s but got %s", test.expected, r)
            }
        })
    }
}