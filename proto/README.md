Protos live here.

Use https://github.com/golang/protobuf to setup proto

cd to repo base direcotry and compile using:

protoc --go_out=proto/gen --proto_path=proto/src proto/src/*proto
