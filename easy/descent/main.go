package main

import "fmt"
// import "os"


func main() {
    for {
        mountains := make([]int,8)
        tallestHeight := 0
        var tallestMount int
        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0.
            var mountainH int
            fmt.Scan(&mountainH)
            mountains[i] = mountainH
        }
        for i, v := range(mountains) {
            if v > tallestHeight {
                tallestHeight = v
                tallestMount = i
            }
        }

        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println(tallestMount) // The number of the mountain to fire on.
    }
}
