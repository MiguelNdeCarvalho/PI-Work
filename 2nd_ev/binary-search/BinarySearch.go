package main

import "fmt"

func pesquisaBinaria(slice int, key []int) bool {

    low := 0
    high := len(key) - 1

    for low <= high{
        median := (low + high) / 2

        if key[median] < slice {
            low = median + 1
        }else{
            high = median - 1
        }
    }

    if low == len(key) || key[low] != slice {
        return false
    }
    return true
}

