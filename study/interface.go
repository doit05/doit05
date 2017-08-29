package study

import (
  "fmt"
)

type Exchanger interface {
       exchange()
}

type StringPair struct {
       first, second string
}

type Point[2]int

func (sp *StringPair) exchange() {
       sp.first, sp.second = sp.second, sp.first
}

func (p *Point) exchange() {
       p[0], p[1] = p[1], p[0]
}

func exchangeThese(exchangers ...Exchanger) {
       for i, exchanger := range exchangers {
            fmt.Printf("%d : %s  Type is = %T\n", i, exchanger, exchanger)
            exchanger.exchange()
       }
}

func main() {
       pair1 := StringPair{"abc","def"}
       pair2 := StringPair{"ghi","jkl"}
       point := Point{5, 7}

       fmt.Println(pair1, pair2, point)
       pair1.exchange()
       pair2.exchange()
       point.exchange()
       fmt.Println(pair1, pair2, point)

       // exchangeThese(pair1, pair2) //wrong
       exchangeThese(&pair1, &pair2)
       fmt.Println(pair1, pair2)
}
