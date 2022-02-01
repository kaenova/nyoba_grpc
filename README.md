Ref: https://tutorialedge.net/golang/go-grpc-beginners-tutorial/

## How To
*Project root active directory
Modification:
1. Create folder service
2. Create folder proto (to save all the proto file)
3. Create proto file in the proto folder
4. Run this command to compile the proto file to desired language, in this case is Golang
```
protoc --go_out=plugins=grpc:. proto/chat.proto
```
5. Create the chat service file (chat.go)

## GUI "Postman" for gRPC
https://github.com/bloomrpc/bloomrpc