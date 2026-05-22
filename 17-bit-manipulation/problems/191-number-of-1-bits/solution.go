package main

// hammingWeight returns the number of 1 bits in the binary representation of num
func hammingWeight(num uint32) int {
    count := 0
    for num != 0 {
        count++
        num = num & (num - 1) // Remove rightmost set bit
    }
    return count
}
