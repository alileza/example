#/bin/sh

mkdir $2/pb
mkdir $2/docs

protoc -I /proto --proto_path=/third_party \
                 --go_out=plugins=grpc:$2/pb \
                 --grpc-gateway_out=logtostderr=true:$2/pb \
                 --swagger_out=logtostderr=true:$2/docs /proto/$1 \
                 --validate_out=lang=go:$2/pb