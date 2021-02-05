# hello-grpc
A gRPC programming exercise.

## Instructions
Install Protocl Buffer compiler `protoc`:

    > sudo apt install protobuf-compiler
    # Ensure that the version is newer than 3.0
    > protoc --version
    libprotoc 3.6.1

Use `protoc` to generate go code:

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hzglexamplesvc/myservice.proto
