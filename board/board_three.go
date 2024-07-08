package board

import "fmt"
import "math"

func compare(xx, yy any) {
    fmt.Println(xx == yy)
}

func exponen(xx, yy int) int {
    return int(math.Pow(float64(xx), float64(yy)))
}

func Board_three() {
    const (
        /* VARIABLE YANG VALUENYA TIDAK MENGACU PADA VARIABEL LAIN*/
        /* MAKA KITA CARI VALUE NYA DI AWAL */
        d int = 15 // 3 × 5 = d
        g int = 2 // 10 -2 × 4 = g
        l int = 5 // 2 + L = 7
        m int = 6 // M ÷ 3 = 2
        s int = 11 // 4 + 7 = S
        /** VARIABEL YANG MENGACU PADA VARIABLE DIATAS **/
        b int = -9 // B + 15 = 6
        i int = 31 // 77 = I + (12 + 11) × 2
        j int = 12 // 6 × 2
        n int = 77 // N = 11 × 7
        t int = 7 // T - 2 = 5
        /** SISA VARIABEL **/
        a int = 76// 26 - (-50)
        c int = 125// 5³
        e int = 26// 77 + 1 = E × 3
        f int = -50 // F + 81 = 31
        h int = 3// H × 9 = 3³
        k int = 168// 44 = 29 - 76 + K - 77
        o int = 44// O + 3⁴ = 125
        p int = 29// P - 26 = 3
        q int = 81 // B²
        r int = 3 // 31 = 4 × 7 + R
    )
    // VARIABLE YANG VALUENYA TIDAK MEMGACU PADA VARIABLE LAIN
    compare(3 * 5, d) // var D
    compare(10 - 2 * 4, g) // var G
    compare(2 + l, 7) //var L
    compare(m / 3, 2) // var M
    compare(4 + 7, s) // var S
    
    // VARIABLE YANG VALUENYA MENGACU PADA VARIABEL DIATAS
    compare(b + d, m) // var B
    compare(n, i + (j + s) * g) //var I
    compare(j, m * g) // var J
    compare(n, s * t) // var N
    compare(t - g, l) //var T
    
    // Variable sisa
    compare(e - f, a) // var A
    compare(c, exponen(l, r)) //var C
    compare(n + 1, e * r) // var E
    compare(f + q, i) //var F
    compare(h * 9, exponen(r, 3)) // var H
    compare(o, p - a + k - n) //var K
    compare(o + exponen(r, 4), c) //var O
    compare(p - e, h) //var P
    compare(q, exponen(b, 2)) //var Q
    compare(i, 4 * t + r) //var R
}