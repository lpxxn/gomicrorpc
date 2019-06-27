
generate proto model

```
  protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. grpcexample/proto/*.proto 
```

