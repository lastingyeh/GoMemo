### GRPC

#### [Prerequisite install](http://google.github.io/proto-lens/installing-protoc.html)

    $ brew install protobuf
    
#### [Install pkg install](https://github.com/grpc/grpc-go)

    $ go get -u google.golang.org/grpc
    
    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    
#### create *.proto

```protobuf
// format version
syntax = "proto3";  
// package defined.
package pb;

// Service defined
service Calculator {  
    // Plus service
    rpc Plus (CalcRequest) returns (CalcReply) {}
}

// CalcRequest (params)
message CalcRequest {  
    int32 number_a = 1;
    int32 number_b = 2;
}

// CalcReply (result)
message CalcReply {  
    int32 result = 1;
}
```

#### File structure (original)

    |- client
    |- pb
    |  |-*.proto
    |- server

#### Translate *.proto to *.go

    $ ./pb protoc --go_out=plugins=grpc:. *.proto
       
#### File structure (new)
 
    |- client
    |- pb
    |  |-*.pb.go
    |  |_*.proto
    |- server
    
### Implement client/server

    see code

---

1.[API 文件就是你的伺服器，REST 的另一個選擇：gRPC](https://yami.io/grpc/)

2.[比起 JSON 更方便、更快速、更簡短的 Protobuf 格式](https://yami.io/protobuf/)

3.[grpc org.doc](https://godoc.org/google.golang.org/grpc) 

4.[Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3)
  