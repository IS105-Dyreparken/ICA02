package algorithms

import (
  "fmt"
)

var toBe [10]int = [10]int{1,3,4,7,9,10,11,16,20,40}

func benchmarkBSortModified(input [10]int) {
  size := len(input)
  swapped := true
  for swapped {
    swapped = false
    for index := 1; index < size-1; index++ {
      if input[index-1] > input[index] {
        fmt.Println("Swapping")
        //Swaps the values
        input[index], input[index-1] = input[index-1], input[index]
        swapped = true
      }
    }
  }
  fmt.Println(input)
}

func main(){
benchmarkBSortModified(toBe)
}
