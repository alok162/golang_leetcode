func myPow(x float64, n int) float64 {
    res := 1.0
    if n > 0 {
        return positive(res, x, n)
    } else {
        return negative(res, x, n)
    }
}

func positive(res float64, x float64, n int) float64 {
     for n > 0 {
        if n & 1 == 1 {
            res = res * x
        }
        x = x * x
        n = n >> 1
    }
    return res
}

func negative(res float64, x float64, n int) float64 {
    n = - n
    for n > 0 {
        if n & 1 == 1 {
            res = res / x
        }
        x = x * x
        n = n >> 1
    }
    return res
}
