package study

import (
	"fmt"

)

 func Count(ch chan int) {
 	fmt.Println("Countint !")
 	ch <- 1

 }

 func ceshi() {

 	chs := make([]chan int, 10)

 	for i := 0; i < 10; i++ {
 		chs[i] = make(chan int)
 		go Count(chs[i])

 	}

 	for _, ch := range chs {
 		fmt.Println(<-ch)
 	}
 }

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum

}

func ceshi2() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultchan := make(chan int, 2)
	go sum(values[0:len(values)/2], resultchan)
	go sum(values[len(values)/2:], resultchan)
	sum1, sum2 := <-resultchan, <-resultchan
	fmt.Println("result : ", sum1, sum2, sum1+sum2)

}


//
//func text1()  {
//				nums := [] int {2, 7, 11, 15}
//				target := 9
//				ret,exist = twoSum(nums,target)
//				fmt.Println(ret, exist)
//}


//}




