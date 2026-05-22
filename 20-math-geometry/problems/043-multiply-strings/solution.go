package main

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" { return "0" }
    m, n := len(num1), len(num2)
    result := make([]int, m+n)
    
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            mul := int(num1[i]-'0') * int(num2[j]-'0')
            p1, p2 := i+j, i+j+1
            sum := mul + result[p2]
            result[p2] = sum % 10
            result[p1] += sum / 10
        }
    }
    
    sb := ""
    for _, v := range result {
        if !(len(sb) == 0 && v == 0) {
            sb += string(rune('0' + v))
        }
    }
    if sb == "" { return "0" }
    return sb
}
