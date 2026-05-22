package main

import "math"

func reverse(x int) int {
    var result int64
    for x != 0 {
        digit := x % 10
        x /= 10
        result = result*10 + int64(digit)
        if result > math.MaxInt32 || result < math.MinInt32 {
            return 0
        }
    }
    return int(result)
}
