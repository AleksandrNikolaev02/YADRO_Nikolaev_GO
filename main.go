// Copyright Nikolaev Alexander

package main
import (
    "fmt"
    "os"
    "sort"
)

func init_array(n int) [][]int {
    var array = make([][]int, n)
    
    for i := 0; i < n; i++ {
        array[i] = make([]int, n)
    }
    
    return array
}

func checkSortBalls(balls [][]int, n int) bool {
    var sum_rows = make([]int, n)
    var sum_col = make([]int, n)

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            sum_rows[i] += balls[i][j]
            sum_col[j] += balls[i][j]
        }
    }
    
    sort.Ints(sum_rows)
    sort.Ints(sum_col)
    
    for i := 0; i < n; i++ {
        if sum_rows[i] != sum_col[i] {return false}
    }
    return true
}

func main() {
    var n int
    
    fmt.Fscan(os.Stdin, &n)
    
    var balls = init_array(n)
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Fscan(os.Stdin, &balls[i][j])
        }
    }
    
    if checkSortBalls(balls, n) {
        fmt.Println("YES")        
    } else {
        fmt.Println("NO")   
    }
}
