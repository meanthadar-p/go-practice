package timeConversion

import (
    "fmt"
    "strconv"
)

func timeConversion(s string) string {
    hour, _ := strconv.ParseInt(s[0:2], 10, 32)
    minute := s[3:5]
    second := s[6:8]

    if t := s[8:10]; (t == "PM" && hour != 12) || (hour == 12 && t == "AM"){
        hour = (hour + 12) % 24
    }

    return fmt.Sprintf("%02d:%s:%s", hour, minute, second)
}