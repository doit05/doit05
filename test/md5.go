package main

import "fmt"
func main() {
				fmt.Println(p())
				fmt.Println(p())
				fmt.Println(p())
				fmt.Println(p())
				for i := 4; i > 0; i-- {
								//tmp :=
								fmt.Println(p())
				}
}

func p() *int {
				x := 344
				return &x
}