package study

import (
	"fmt"
	"sort"
)

type Test struct {
	name string
}

func (t Test) Test() {
	fmt.Printf("test")
}

func (t Test) GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

type in interface {
	Test()
}

func main() {
	var a int = 4
	b := 4
	var arr = []int{1, 2, 3, 4}
	var array = [5]float32{1, 2, 3, 4, 5}
	test := Test{}
	var intest in
	intest = test
	test.GuessingGame()

	fmt.Printf("%v \n", a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", b)
	fmt.Printf("%v \n", arr)
	fmt.Printf("%T \n", arr)
	fmt.Printf("%v \n", array)
	fmt.Printf("%T \n", array)
	fmt.Printf("%T \n", test.Test)
	fmt.Printf("%+v \n", test)
	fmt.Printf("%T \n", intest)
	fmt.Printf("%+v \n", intest)
}
