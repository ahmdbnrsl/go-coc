package board

import "fmt"
import "math"

func factorial(value int) int {
    if value == 1 {
        return 1
    } else {
        result := value * factorial(value - 1)
        return result
    }
}

func Board_two() {
    r := map[string]int {
        "a": 14,
        "b": 9,
        "c": 8,
        "d": 8,
        "e": 11,
        "f": 3,
        "g": -1,
        "h": 9,  
        "i": 3,
        "j": 9,  
        "k": 5,  
        "l": -1,
        "m": 2,
        "n": 4,
        "o": 22,
        "p": 6,
    }
    
    //VALUE BELOW MUST BE TRUE
    println(r["d"] == r["c"] / 2 + r["a"] - 2 * r["k"])
    println((-1) * r["b"] == r["g"] * r["j"])
    println(r["c"] == 8)
    println(r["n"] + int(math.Pow(2, 2)) == r["d"])
    println(r["i"] * (r["e"] - r["c"]) == r["b"])
    println(2 * r["f"] == r["p"])
    println(factorial(4) / 3 == r["g"] + r["h"])
    println(int(math.Pow(float64(r["p"]), 2)) == r["h"] * 4)
    println(r["n"] + r["k"] == r["f"] * r["i"])
    println(int(math.Pow(float64(3 * (r["k"] - r["p"])), 2)) == r["j"])
    println(r["k"] == 5)
    println(int(math.Pow(3, 2)) - r["d"] == r["l"] * -1)
    println(int(math.Pow(float64(r["k"]), 2)) + r["m"] == int(math.Pow(float64(r["i"]), 3)))
    println(r["n"] == int(math.Pow(3, 2) - 5))
    println(r["i"] - r["m"] == (2 * r["a"] - factorial(3)) / r["o"])
    println(r["p"] == 12 / 2)
    
    fmt.Println(r)
}