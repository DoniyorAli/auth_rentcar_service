protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./protoc/article/*.proto

protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./protoc/author/*.proto