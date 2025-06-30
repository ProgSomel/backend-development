package main

import "fmt"



func main() {
    numbers := []int{25, 99, -3, 99, 56, 78, 42}

    if(len(numbers) == 0){
        fmt.Println("Array is empty")
        return
    }

    max := numbers[0]
    maxIndex := 0

    for i, num := range numbers{
        if num >= max{
            max = num
            maxIndex = i
        }
        
        
    }
    fmt.Printf("Max value: %d found at index: %d\n", max, maxIndex)
}
