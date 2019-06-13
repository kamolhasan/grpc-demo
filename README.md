# grpc-demo

This is a demo grpc server that shows how a grpc server communicates with clients.
### Code Gen

```console
$ protoc -I . ./src/hello.proto --go_out=plugins=grpc:.
```

## User Guide
* Create directory 
* Clone repository 
* Go inside 
    ```console
    $ cd /home/kamol/go/src/github.com/kamolhasan/grpc-demo/
    ```
* Start server
    ```console
    $  go run ./cmd/server/main.go
    ```
* Start client
    ```console
    # go run ./cmd/client/main.go "<your-message>"
    $ go run ./cmd/client/main.go "my name is kamol"
      Succeed! Code:  200
      I got your message! Thanks! :D
    ```