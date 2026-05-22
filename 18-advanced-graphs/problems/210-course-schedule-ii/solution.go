package main

func findOrder(numCourses int, prerequisites [][]int) []int {
    adj := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, p := range prerequisites {
        adj[p[1]] = append(adj[p[1]], p[0])
        indegree[p[0]]++
    }
    
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 { queue = append(queue, i) }
    }
    
    result := []int{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)
        for _, next := range adj[node] {
            indegree[next]--
            if indegree[next] == 0 { queue = append(queue, next) }
        }
    }
    
    if len(result) != numCourses { return []int{} }
    return result
}
