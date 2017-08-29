package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type Book struct {
        Name  string
        Price float64 // `json:"price,string"`
    }
    var person = struct {
        Name string
        Age  int
        // Book
        Book `json:"Book"`
    }{
        Name: "doit",
        Age:  26,
        Book: Book{
            Price: 13.4,
            Name:  "Go语言",
        },
    }

    buf, _ := json.Marshal(person)
    fmt.Println(string(buf))

    // Output:{"Name":"polaris","Age":30,"Price":3.4}
    // Book 中的 Name 被忽略了
}
