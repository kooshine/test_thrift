package main

import (
    "fmt"
    "github.com/kooshine/test-rpc/hello"
    "github.com/apache/thrift/lib/go/thrift"
    "golang.org/x/net/context"
)

var ctx = context.Background()

func main(){
    var transport thrift.TTransport
    var err error

    // 传输方式（需要与服务端一致）
    transport, err = thrift.NewTSocket("localhost:8000")
    if err != nil {
        fmt.Printf("open socket err, %s\n", err)
        return
    }

    // 传输协议（需要与服务端一致）
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    iProtocol := protocolFactory.GetProtocol(transport)
    oProtocol := protocolFactory.GetProtocol(transport)
    tClient := thrift.NewTStandardClient(iProtocol, oProtocol)

    transmit := hello.NewTransmitClient(tClient)
    if err := transport.Open(); err != nil {
        fmt.Printf("Error opening socket to :8000 %s", err)
        return
    }
    defer transport.Close()
    data := "{heelo:123}"
    res, err := transmit.SayMsg(ctx,data)

    if err != nil {
        fmt.Printf("get user err, %s", err)
        return
    }
    fmt.Printf("%s", res)
}
