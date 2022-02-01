Ref: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/

## How To
*Project root active directory
Modification:
1. Create folder service
2. Create folder proto (to save all the proto file)
3. Create folder service
4. Run this command
```
protoc --go_out=plugins=grpc:. proto/chat.proto
```
5. Create the chat service file (chat.go)

## GUI "Postman" for gRPC
https://github.com/bloomrpc/bloomrpc