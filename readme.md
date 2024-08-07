This is a practice project.

Stack: Protobuf, gRPC, Go, React, Envoy, Nginx


Simply run these commands:

```
mkdir client/generated
mkdir server/generated
mkdir fe_client/src/generated
```

Setting up protobuf and clients (Mac)
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```
export PATH="$PATH:$(go env GOPATH)/bin"
source ~/.zshrc
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
cd fe_react
npm install -g protobufjs
```

```
brew install protoc-gen-grpc-web
```

```
protoc -I=proto proto/hello.proto \
  --js_out=import_style=commonjs:fe_client/src/generated \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:fe_client/src/generated
```


For Mac:
Inside the project directory
```
docker compose up --build
```
