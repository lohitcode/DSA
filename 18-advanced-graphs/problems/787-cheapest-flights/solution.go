package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prices := make([]int, n)
    for i := range prices { prices[i] = 1<<31 - 1 }
    prices[src] = 0
    
    for i := 0; i <= k; i++ {
        temp := make([]int, n)
        copy(temp, prices)
        changed := false
        for _, f := range flights {
            from, to, price := f[0], f[1], f[2]
            if prices[from] != 1<<31-1 && prices[from]+price < temp[to] {
                temp[to] = prices[from] + price
                changed = true
            }
        }
        prices = temp
        if !changed { break }
    }
    
    if prices[dst] == 1<<31-1 { return -1 }
    return prices[dst]
}
