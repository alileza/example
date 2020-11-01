#/bin/sh

rm -rf $2/*
mkdir $2/ts
mkdir $2/pb
mkdir $2/docs

protoc -I /proto --proto_path=/third_party \
                 --plugin=/usr/local/bin/protoc-gen-ts_proto \
                 --go_out=plugins=grpc:$2/pb \
                 --grpc-gateway_out=logtostderr=true:$2/pb \
                 --swagger_out=logtostderr=true:$2/docs /proto/$1 \
                 --validate_out=lang=go:$2/pb \
                 --js_out=import_style=commonjs,binary:$2/ts \
                 --ts_proto_out=$2/ts  \
                 --ts_proto_opt=outputClientImpl=false \
                 --ts_proto_opt=forceLong=long \
                 --ts_proto_opt=env=browser \
                 --ts_proto_opt=lowerCaseServiceMethods=true \
                 --ts_proto_opt=useOptionals=false
                 