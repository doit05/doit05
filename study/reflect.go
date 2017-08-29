package study

import (
   "fmt"
   "reflect"
)

// http://speakingbaicai.blog.51cto.com/5667326/1707637

func main() {
   var x float64 = 1.1
   fmt.Println("reflect.Value:", reflect.ValueOf(x))
   fmt.Println("reflect.Type:", reflect.TypeOf(x))
   v := reflect.ValueOf(x)
   fmt.Println("reflect.Type:",v.Type())
   fmt.Println("actual value:", v.Float())
   fmt.Println("kind is float64?", v.Kind() == reflect.Float64)
   fmt.Println("----------------------")
   test1()
   fmt.Println("----------------------")
   test2()
   fmt.Println("----------------------")
   test3()
   fmt.Println("----------------------")
   test4()
   fmt.Println("----------------------")
   test5()
}

func test1()  {
  type MyIntStructStruct struct {
    value int
  }
   var x MyIntStructStruct = MyIntStructStruct{value:1}
   v := reflect.ValueOf(x)
   fmt.Println("reflect.Type:", v.Type())
   fmt.Println("kind is struct ? ", v.Kind() == reflect.Struct)
   fmt.Printf("v type is : %v\n", v)
}

func test2() {
  type MyInt int
   var x MyInt = 1
   v := reflect.ValueOf(x)
   fmt.Println("reflect.Type:", v.Type())
   fmt.Println("kind is int?", v.Kind() == reflect.Int)
}

func test3() {
   var x float64 = 1.1
   fmt.Println("reflect.Value:", reflect.ValueOf(x))
   fmt.Println("reflect.Type:", reflect.TypeOf(x))
   v := (reflect.ValueOf(x))
   fmt.Println("reflect.Type:", v.Type())
   fmt.Println("actual value(interface):", v.Interface())
   fmt.Println("kind is float64?", v.Kind() == reflect.Float64)
}

func test4() {
    var x float64 = 1.1
    v := reflect.ValueOf(x)
    fmt.Println("settability of v:",v.CanSet())
    // v.SetFloat(1.2)
}

func (v Value)Elem() Value
//Elem returns the value that the interface v contains or that the pointer vpoints to. It panics if v's Kind is not Interface or Ptr. It returns the zeroValue if v is nil.

func test5() {
    var x float64 = 1.1
    p := reflect.ValueOf(&x)
    fmt.Println("type of p:",p.Type())
    fmt.Println("type of p:",p.Interface())
    v := p.Elem()
    fmt.Println("type of v:",v.Type())
    fmt.Println("settability of v:",v.CanSet())
}
