This is a practice project.

It has a Go gRPC server, a Go client to access the gRPC server, envoy proxy, React client to access gRPC server via envoy proxy.


Simply run these commands:

```
mkdir client/generated
mkdir server/generated
mkdir fe_cleint/src/generated
```

Setting up protobuf and clients (Mac)
```
brew install protobuf

```

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```
npm install protobufjs
```

```
brew install protoc-gen-grpc-web
```

For the server directory
```
protoc --go_out=./server/generated --go-grpc_out=./server/generated proto/hello.proto
```

For the client directory
```
protoc --go_out=./client/generated --go-grpc_out=./client/generated proto/hello.proto
```

For the React app directory
```
protoc -I=proto proto/hello.proto \
  --js_out=import_style=commonjs:fe_client/src/generated \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:fe_client/src/generated
```


For Mac:
```
docker compose up --build
```
