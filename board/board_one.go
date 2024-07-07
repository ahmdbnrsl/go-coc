package board

import "math"
import "fmt"

func Board_one() {
    a := 8
    b := 9
    c := 2
    d := 26
    e := 4
    f := -13
    g := 3
    h := 9
    i := int(math.Sqrt(math.Pow(8, 2) + math.Pow(6, 2)))
    
    // THE VALUE BELOW MUST BE TRUE
    
    println(a / (-8) == int(math.Pow(float64(g), 2)) - 1 - h)
    println(b == -3 * g * (-1))
    println(float32(float32(2) / 64) == float32(1 / math.Pow(float64(c), 5)))
    println(f + d == 2 * a - 12 / e)
    println(6 * c == e * g)
    println(c * f == -2 * (b + e))
    println(2 + g == 5)
    println(g == i - h + c)
    println(i == int(math.Sqrt(math.Pow(8, 2) + math.Pow(6, 2))))
    
    result := map[string]any {
        "a": a,
        "b": b,
        "c": c,
        "d": d,
        "f": f,
        "g": g,
        "h": h,
        "i": i,
    }
    
    fmt.Println(result)
}