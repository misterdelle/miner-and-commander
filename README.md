# miner-and-commander

To turn on the logs for debugging, run the code with the following environment variable: GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info

In a CMD shell type:
SET GRPC_GO_LOG_VERBOSITY_LEVEL=99
SET GRPC_GO_LOG_SEVERITY_LEVEL=info
go run .

In a linux shell type:
GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run .

In a POWERSHELL shell type:
$Env:GRPC_GO_LOG_VERBOSITY_LEVEL=99
$Env:GRPC_GO_LOG_SEVERITY_LEVEL='info'
go run .


To Generate gRPC code: Generate the necessary Go code for gRPC from the .proto definitions:
c:\ProtocolBuffers\protoc-31.1-win64\bin\protoc -I bos-plus-api\proto bos-plus-api\proto\bos\*.proto bos-plus-api\proto\bos\v1\*.proto --go_out=pb --go-grpc_out=pb

## To build docker image for Raspberry PI
docker buildx build --platform linux/arm64/v8 --output type=docker -t misterdelle/miner-and-commander:latest .

## To build docker image for Linux
docker buildx build --platform linux/amd64 --output type=docker -t misterdelle/miner-and-commander:latest .
