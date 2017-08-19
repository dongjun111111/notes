* 定义结构 hello.proto
* 编译hello.proto 执行命令：  protoc -I . --go_out=plugins=grpc:. ./hello.proto
* 运行服务端与客户端