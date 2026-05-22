package main

func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := make([][]int, numCourses)
    for _, p := range prerequisites {
        adj[p[1]] = append(adj[p[1]], p[0])
    }
    
    state := make([]int, numCourses)
    var hasCycle func(int) bool
    hasCycle = func(node int) bool {
        if state[node] == 1 { return true }
        if state[node] == 2 { return false }
        state[node] = 1
        for _, next := range adj[node] {
            if hasCycle(next) { return true }
        }
        state[node] = 2
        return false
    }
    
    for i := 0; i < numCourses; i++ {
        if state[i] == 0 && hasCycle(i) { return false }
    }
    return true
}
