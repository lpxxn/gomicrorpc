
generate proto model

```
  protoc --proto_path=$GOPATH/src:. --go_out=. example2/proto/model/*.proto 
  
  protoc --proto_path=$GOPATH/src:. --micro_out=. example2/proto/rpcapi/*.proto 
```

