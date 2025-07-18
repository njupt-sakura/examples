# nest-demo

## Step 1: item-server

```bash
pwd # /path/to/nest-demo

pnpm i -g @nestjs/cli
nest new item-server
mkdir ./client ./order-server

brew install protobuf # MacOS

cd ./client
pnpm init # "type": "module"

cd ../order-server
go mod init com.github.njupt-sakura/nest-demo/order-server

cd ../item-server
pnpm add @grpc/grpc-js @grpc/proto-loader
pnpm add ts-proto -D

mkdir ./src/codegen
protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto \
       --ts_proto_out=./src/codegen                     \
       --ts_proto_opt=nestJs=true                       \
       --proto_path=../proto                            \
       common.proto order.proto item.proto
```

## Step2: order-server

```bash
pwd # /path/to/nest-demo

cd ./order-server
brew install protoc-gen-go-grpc # MacOS
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

mkdir ./codegen
protoc --go_out=./codegen                  \
       --go_opt=paths=source_relative      \
       --proto_path=../proto               \
       --go-grpc_out=./codegen             \
       --go-grpc_opt=paths=source_relative \
       common.proto order.proto item.proto

go mod tidy
```

## Step 3: client

```bash
pwd # /path/to/nest-demo

cd ./client
# "type": "module"
pnpm add @types/node ts-proto -D
pnpm add @grpc/grpc-js @grpc/proto-loader
```

## Test

```bash
# Item Service (Powered by Node.js)
cd ./item-server && pnpm start

# Order Service (Powered by Go)
cd ./order-server && go run .

# Client
cd ./client && node ./main.js

# [orderClient] findOne res: {"id":1,"price":1.1100000143051147}
# [itemClient] findOne res: {"id":1,"name":"161043261","url":"https://github.com/161043261"}
# [orderClient] findOneWithItem res: {"id":1,"price":1.1100000143051147,"item":{"id":1,"name":"161043261"}}
# [itemClient] findOneWithOrder res: {"id":1,"name":"161043261","order":{"id":1,"price":1.1100000143051147}}
# [itemClient] findMany res: {"list":[{"id":1,"name":"161043261","url":"https://github.com/161043261"},{"id":2,"name":"tianchenghang","url":"https://github.com/tianchenghang"}]}
```
