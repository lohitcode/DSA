package main

func isHappy(n int) bool {
    slow, fast := n, n
    for {
        slow = sumSquares(slow)
        fast = sumSquares(sumSquares(fast))
        if slow == fast { break }
    }
    return slow == 1
}

func sumSquares(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}
