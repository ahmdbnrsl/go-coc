package board

func Board_five() {
    const (
        /* VARIABLE YANG VALUENYA TIDAK MENGACU PADA VARIABEL LAIN*/
        /* MAKA KITA CARI VALUE NYA DI AWAL */
        k int = 8// K = 4² - 2³ (VAR K)
        
        /** VARIABEL YANG MENGACU PADA VARIABLE DIATAS **/
        d int = 4// D = 4! ÷ 8 + 1 (VAR D)
        h int = 9// H = 75% × (4 + 8) (VAR H)
        o int = 3// O = 4 + 8 - 9 (VAR O)
        t int = 6// T + 3 = 9² ÷ 3² (VAR T)
        p int = 12// 0,25 × P = (8 - 6)³ - 5 (VAR P)
        u int = 10// U = (9 - 4)! ÷ 12 (VAR U)
        x int = 3// X + 1 = 1 ÷ 4 × (12 + 4) (VAR X)
        cc int = 4// cc³ = (12 - 4) × 8 (VAR CC)
        e int = 5// 9 + 4 - E = 3² - 1 (VAR E)
        a int = 5// A³ = 12 × 10 + 5 (VAR A)
        g int = 16// 5 × 5 = G + 9 (VAR G)
        m int = 2// 4 × 5 = (3! + 4) × m (VAR M)
        aa int = 17 // aa - 5 = 16 × 3 ÷ 2² (VAR AA)
        r int = 6// (12 ÷ 3)² × 2 - 17 = 5 + 4 + R (VAR R)
        s int = 7// 8 - 6 + S = 3⁴ ÷ 9 (VAR S)
        //z int = // 5 × 7 - Z = (12 ÷ 3) ÷ (1 ÷ 3)
        /** KOREKSI **/
        /** HARUS MENGEMBALIKAN NILAI TRUE **/
    )
    compare(k, exponen(4, 2) - exponen(2, 3)) // var K
    compare(d, factorial(4) / k + 1) // var D
    compare(h, int(float32(75) / 100 * float32(d + k))) // var H
    compare(o, d + k - h) // var O
    compare(t + 3, exponen(h, 2) / exponen(o, 2)) // var T
    compare(int(0.25 * float64(p)), exponen(k - t, 3) - 5) // var P
    compare(u, factorial(h - d) / p) // var U
    compare(x + 1, int((float32(1) / 4) * (float32(p) + float32(d)))) // var X
    compare(exponen(cc, 3), (p - d) * k) // var CC
    compare(h + d - e, exponen(o, 2) - 1) // var E
    compare(exponen(a, 3), p * u + e) // var A
    compare(a * e, g + h) // var G
    compare(d * e, (factorial(x) + 4) * m) // var M
    compare(aa - a, g * x / exponen(m, 2)) // var AA
    compare(exponen(p / x, 2) * m - aa, a + d + r) //var R
    compare(k - r + s, exponen(x, 4) / h) // var S
}