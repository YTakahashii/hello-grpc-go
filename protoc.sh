PROTO_SRC=./idl
PROTO_DEST=./interfaces/rpc/

mkdir -p $PROTO_DEST

# Generate User
protoc -I idl/ --plugin=$HOME/golang/bin/protoc-gen-go --go_out=plugins=grpc:$PROTO_DEST $PROTO_SRC/*