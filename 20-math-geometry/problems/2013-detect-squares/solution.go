package main

type DetectSquares struct {
    points map[[2]int]int
    addPts [][]int
}

func Constructor() DetectSquares {
    return DetectSquares{points: make(map[[2]int]int)}
}

func (this *DetectSquares) Add(point []int) {
    this.points[[2]int{point[0], point[1]}]++
    this.addPts = append(this.addPts, point)
}

func (this *DetectSquares) Count(point []int) int {
    x1, y1 := point[0], point[1]
    result := 0
    for _, p := range this.addPts {
        x3, y3 := p[0], p[1]
        if x3 != x1 && abs(x3-x1) == abs(y3-y1) {
            result += this.points[[2]int{x1, y3}] * this.points[[2]int{x3, y1}]
        }
    }
    return result
}

func abs(x int) int { if x < 0 { return -x }; return x }
