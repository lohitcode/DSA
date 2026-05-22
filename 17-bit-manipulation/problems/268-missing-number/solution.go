package main

func missingNumber(nums []int) int {
    result := len(nums)
    for i, num := range nums {
        result ^= i ^ num
    }
    return result
}
