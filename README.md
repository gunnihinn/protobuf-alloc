# protobuf-alloc

This program demonstrates excessive memory use in 
[gogo/protobuf's](https://github.com/gogo/protobuf)
deserialization of repeated fields. It's only purpose is to be pointed to in a
Github issue discussing the problem.

## Build and run

Run:

    $ make proto
    $ make

I didn't make the generated proto files a dependency of the default target so I
could edit them by hand without `make` overriding my changes.

## Sample results

```
## Commit 36a0ed7 (current behaviour):

$ ./protobuf-alloc
$ go tool pprof protobuf-alloc memory.prof
File: protobuf-alloc
Type: inuse_space
Time: Jul 24, 2018 at 12:01pm (CEST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top                                                                                                                                                                  
Showing nodes accounting for 5175.25MB, 100% of 5175.25MB total                                                                                                              
      flat  flat%   sum%        cum   cum%
 2871.24MB 55.48% 55.48%  2871.24MB 55.48%  github.com/gunnihinn/protobuf-alloc/foo.(*Foo).Unmarshal
    2048MB 39.57% 95.05%  5175.25MB   100%  main.main
  256.01MB  4.95%   100%   256.01MB  4.95%  github.com/gunnihinn/protobuf-alloc/foo.(*Foo).Marshal
         0     0%   100%  5175.25MB   100%  runtime.main

## Commit 97163c0 (with ad-hoc allocation patch):

$ ./protobuf-alloc
$ go tool pprof protobuf-alloc memory.prof                                                                                         
File: protobuf-alloc                                                                                                                                                         
Type: inuse_space
Time: Jul 24, 2018 at 12:03pm (CEST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top                                                                                                                                                                  
Showing nodes accounting for 2304.01MB, 100% of 2304.01MB total                                                                                                              
      flat  flat%   sum%        cum   cum%
    2048MB 88.89% 88.89%  2304.01MB   100%  main.main
  256.01MB 11.11%   100%   256.01MB 11.11%  github.com/gunnihinn/protobuf-alloc/foo.(*Foo).Marshal
         0     0%   100%  2304.01MB   100%  runtime.main
```
