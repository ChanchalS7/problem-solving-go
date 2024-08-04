package main


func canPlaceFlowers(flowerbed []int, n int) bool {
    l := len(flowerbed)
    if n == 0 { return true }
    if l == 1 && flowerbed[0] == 0 && n == 1 { return true }
    if l == 2 && flowerbed[0] == 0 && flowerbed[1] == 0 && n == 1 { return true }
    for i := 1; i < l - 1; i++ {
        if i-1 == 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
            flowerbed[i-1] = 1
            n--
        }
        if i+1 == l-1 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
            flowerbed[i+1] = 1
            n--
        }
        if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            n--
        }
        if n <= 0 { return true }
    }
    return false
}

/*
func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 { return true }
    for i, c := 0, 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            var prev, next int
            if i != 0 { prev = flowerbed[i - 1] }
            if i != len(flowerbed) - 1 { next = flowerbed[i + 1] }
            if prev == 0 && next == 0 { flowerbed[i] = 1; c++ }
            if c == n { return true }
        }
    }
    return false
}
*/