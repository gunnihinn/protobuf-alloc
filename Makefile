protobuf-alloc: main.go foo/foo.pb.go
	go build -o protobuf-alloc

foo/foo.pb.go: foo.proto foo
	protoc --gofast_out=foo foo.proto

foo:
	mkdir foo

clean:
	rm -rf protobuf-alloc foo
