protobuf-alloc: main.go
	go build -gcflags=all='-l -N' -o protobuf-alloc

proto: foo.proto foo
	protoc --gofast_out=foo foo.proto

foo:
	mkdir foo

clean:
	rm -rf protobuf-alloc foo
