package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", ":8086")
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()

    //循环接收客户端的连接，没有连接时会阻塞，出错则跳出循环
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("err : ", err)
            break
        }

        fmt.Println("[Server] Accept new conn.")

        go handler(conn) //启动新的goroutime处理连接
    }

}

func handler(conn net.Conn)  {
    defer conn.Close()
    //循环从连接中，读取请求内容，没有请求时会阻塞，出错则跳出
    for {
        request := make([]byte, 20)
        readLen, err := conn.Read(request)

        if err != nil {
            fmt.Println("handler err : ", err)
            break
        }

        if readLen == 0 {
            fmt.Println("Read 0 character.")
            break
        }


        //控制台输出读到的请求内容，并在请求的内容前加上Copy和时间后向客户端输出
        fmt.Println("[server] request from : ", string(request))
        conn.Write([]byte("Copy " + string(request) + " time : " + time.Now().Format("2015-10-11 15:30:21") ))
    }

}
