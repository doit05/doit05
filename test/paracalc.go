package test

import (
	"fmt"

	"github.com/doit05/stringutil"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println("Result: ", sum1, sum2, sum1+sum2)
	fmt.Printf(stringutil.Reverse("!dlroW ,olleH"))
}
