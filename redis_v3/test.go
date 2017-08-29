package main

import (
  "fmt"
  "gopkg.in/redis.v3"
)

func main() {
  redis_ip := "192.168.0.106"
  redis_port := "6379"
  client := NewClient(redis_ip, redis_port)
  ExampleClient(client)
}

func NewClient(redis_ip, port string) (client *redis.Client) {

    client = redis.NewClient(&redis.Options{
        Addr:     redis_ip + ":" + port,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    return
}

func ExampleClient(client *redis.Client) {
    err := client.Set("key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := client.Get("key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := client.Get("key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exists")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
}
