# MyCookBook

Dockerfile -> go linter

docker build -t <image_name>:<tag> .
	
docker run --rm -v $(pwd):/app -w /app <image_name>:<tag>
	
gRPC Installation

Install packages
Run -> dep ensure
To create protobuf
create a file names <filename>.protobuf
Run -> protoc --go_out=plugins=grpc:. protos/helloworld.proto

Run server
go run server/main.go
https://github.com/grpc/grpc/blob/d910557f22ab445a3618a358d7a6bb17836c7151/BUILDING.md												

Create grpc_cli for Client

git clone grpc repo -> git clone -b v1.34.1 https://github.com/grpc/grpc
git submodule update --init
Run from grpc repo
$ mkdir -p cmake/build
$ cd cmake/build
$ cmake ../..
$ make
$ make clean
$ cmake -DgRPC_BUILD_TESTS=ON ../..
$ make grpc_cli
$ grpc_cli ls localhost:50051 -l
$ grpc_cli call localhost:50051 SayHello 'name: "john"'	

