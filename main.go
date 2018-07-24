package main

import (
	"os"
	"runtime/pprof"

	"github.com/gunnihinn/protobuf-alloc/foo"
)

func main() {
	f := &foo.Foo{
		Stuff: make([]int64, 1<<28), // Around 2GiB of stuff
	}

	bs, err := f.Marshal()
	if err != nil {
		panic(err)
	}

	g := new(foo.Foo)
	if err := g.Unmarshal(bs); err != nil {
		panic(err)
	}

	fh, err := os.Create("memory.prof")
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	if err := pprof.WriteHeapProfile(fh); err != nil {
		panic(err)
	}
}
