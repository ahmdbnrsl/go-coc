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
        z int = 23// 5 × 7 - Z = (12 ÷ 3) ÷ (1 ÷ 3) (var Z)
        dd int = 11// 2 ÷ 5 = 4 ÷ (3DD - 23) (var DD)
        c int = 10// 11 + 8 + 5 = (C - 6)! (var C)
        v int = 14// V × 10 = 12² - 4 (VAR V)
        n int = 21// (14 + 10) × 1 ÷ 2 = N - 9 (VAR N)
        f int = 1// (21 ÷ 7)³ = 14 + 12 + F (VAR F)
        j int = 6// 23 + 5 - 16 = J³ × 1 ÷ 18 (VAR J)
        y int = 2// (10 - 1) × 6 = 21 + 23 + Y³ + 2 (VAR Y)
        l int = 9// 5² + L = 2 × 10 + 14 (VAR L)
        w int = 15// (14 - 9)! ÷ W = 4 × 12 ÷ 6 (VAR W)
        bb int = 3// 15 × 6 ÷ 9 = (5 + 1) × 5 ÷ BB (VAR BB)
        q int = 11// (15 × 3 × 2 ÷ 9)² = (Q + 9) × 5 (VAR Q)
        b int = 9// ((11 - 3) × 50%)² = 4 + 21 - B (var B)
        i float32 = -3.75// 5 × 1 + I = (10 + 15) ÷ (17 + 3)
    )
    /** KOREKSI **/
    /** HARUS MENGEMBALIKAN NILAI TRUE **/
    
    compare(k, exponen(4, 2) - exponen(2, 3)) // 1.var K 
    compare(d, factorial(4) / k + 1) // 2.var D
    compare(h, int(float32(75) / 100 * float32(d + k))) // 3.var H
    compare(o, d + k - h) // 4.var O
    compare(t + 3, exponen(h, 2) / exponen(o, 2)) // 5.var T
    compare(int(0.25 * float64(p)), exponen(k - t, 3) - 5) // 6.var P
    compare(u, factorial(h - d) / p) // 7.var U
    compare(x + 1, int((float32(1) / 4) * (float32(p) + float32(d)))) // 8.var X
    compare(exponen(cc, 3), (p - d) * k) // 9.var CC
    compare(h + d - e, exponen(o, 2) - 1) // 10.var E
    compare(exponen(a, 3), p * u + e) // 11.var A
    compare(a * e, g + h) // 12.var G
    compare(d * e, (factorial(x) + 4) * m) // 13.var M
    compare(aa - a, g * x / exponen(m, 2)) // 14.var AA
    compare(exponen(p / x, 2) * m - aa, a + d + r) // 15.var R
    compare(k - r + s, exponen(x, 4) / h) // 16.var S
    compare(5 * s - z, int((float32(p) / 3) / (float32(1) / float32(o)))) // 17.var Z
    compare(2 / float32(a), float32(d) / (3 * float32(dd) - float32(z))) // 18.var DD
    compare(dd + k + e, factorial(c - r)) // 19. var C
    compare(v * c, exponen(p, 2) - cc) // 20. var V
    compare((v + u) * 1 / 2, n - h) // 21. var N
    compare(exponen(n / s, 3), v + p + f) //22. var F
    compare(z + e - g, exponen(j, 3) * 1 / 18) // 23. var J
    compare((c - f) * r, n + z + exponen(y, 3) + 2) // 24. var Y
    compare(exponen(e, y) + l, 2 * c + v) // 25. var L
    compare(factorial(v - l) / w, 4 * p / j) // 26. var W
    compare(w * r / 9, (e + f) * a / bb) // 27. var BB
    compare(exponen(w * x * y / l, m), (q + l) * 5) // 28. var Q
    compare(exponen((q - bb) * 50 / 100, 2), d + n - b) // 29. var B
    compare(float32(e * f) + i, float32(u + w) / float32(aa + bb)) // 30. var I
}