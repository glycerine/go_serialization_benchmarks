# Benchmarks of Go serialization methods

This is a test suite for benchmarking various Go serialization methods.

## Tested serialization methods

- [encoding/gob](http://golang.org/pkg/encoding/gob/)
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [github.com/alecthomas/binary](https://github.com/alecthomas/binary)
- [github.com/davecgh/go-xdr/xdr](https://github.com/davecgh/go-xdr)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/ugorji/go/codec](https://github.com/ugorji/go/tree/master/codec)
- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [github.com/youtube/vitess/go/bson](https://github.com/youtube/vitess/tree/master/go/bson) *(using the bsongen code generator)*
- [labix.org/v2/mgo/bson](https://labix.org/v2/mgo/bson)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp) *(code generator for msgpack)*
- [github.com/glycerine/zebrapack](https://github.com/glycerine/zebrapack) *(code generator)*
- [github.com/golang/protobuf](https://github.com/golang/protobuf) (generated code)
- [github.com/gogo/protobuf](https://gogo.github.io/) (generated code, optimized version of `goprotobuf`)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf) (reflection based)
- [github.com/google/flatbuffers](https://github.com/google/flatbuffers)
- [github.com/hprose/hprose-go/io](https://github.com/hprose/hprose-go)
- [github.com/glycerine/go-capnproto](https://github.com/glycerine/go-capnproto)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)
- [git.apache.org/thrift.git/lib/go/thrift](https://thrift.apache.org/lib/go) (Thrift CompactBinary and Thrift Binary formats)

## Running the benchmarks

```bash
go get -u -t
go test -bench='.*' ./
```

## Recommendation

Adapt these benchmarks to your data, and make your own measurements.

## Data

The data being serialized is the following structure with randomly generated values:

```go
type A struct {
    Name     string
    BirthDay time.Time
    Phone    string
    Siblings int
    Spouse   bool
    Money    float64
}
```


## Results

Results with Go 1.8beta2 on a 3.1 GHz Intel Core i7 (MacBook Pro 13"):

* sorted by read speed

```
benchmark                                       iter           time/iter         bytes alloc       allocs
---------                                       ----           ---------         -----------       ------
BenchmarkGencodeUnmarshal-4              	10000000	       196 ns/op	     112 B/op	       3 allocs/op
BenchmarkCapNProtoUnmarshal-4            	10000000	       210 ns/op	       0 B/op	       0 allocs/op
BenchmarkGreenPackUnmarshal-4            	10000000	       224 ns/op	       0 B/op	       0 allocs/op
BenchmarkFlatBuffersUnmarshal-4          	10000000	       236 ns/op	      32 B/op	       2 allocs/op
BenchmarkZebraPackUnmarshal-4            	10000000	       236 ns/op	      0 B/op	       0 allocs/op
BenchmarkGogoprotobufUnmarshal-4         	10000000	       240 ns/op	      96 B/op	       3 allocs/op
BenchmarkMsgpUnmarshal-4                 	 5000000	       300 ns/op	      32 B/op	       2 allocs/op
BenchmarkProtobufUnmarshal-4             	 2000000	       690 ns/op	     192 B/op	      10 allocs/op
BenchmarkGoprotobufUnmarshal-4           	 2000000	       701 ns/op	     432 B/op	       9 allocs/op
BenchmarkGobUnmarshal-4                  	 2000000	       943 ns/op	     112 B/op	       3 allocs/op
BenchmarkHproseUnmarshal-4               	 1000000	      1094 ns/op	     320 B/op	      10 allocs/op
BenchmarkCapNProto2Unmarshal-4           	 1000000	      1426 ns/op	     608 B/op	      12 allocs/op
BenchmarkBinaryUnmarshal-4               	 1000000	      1579 ns/op	     336 B/op	      22 allocs/op
BenchmarkXdrUnmarshal-4                  	 1000000	      1635 ns/op	     239 B/op	      11 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-4   	  500000	      2406 ns/op	     384 B/op	      13 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-4   	  500000	      2525 ns/op	    3008 B/op	       6 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-4      	  500000	      2740 ns/op	    3168 B/op	       9 allocs/op
BenchmarkSerealUnmarshal-4               	  500000	      3356 ns/op	    1008 B/op	      34 allocs/op
BenchmarkJsonUnmarshal-4                 	  300000	      4380 ns/op	     495 B/op	       8 allocs/op
```

* sorted by write speed


```
benchmark                                       iter           time/iter          bytes alloc      allocs
---------                                       ----           ---------          -----------      ------
BenchmarkZebraPackMarshal-4              	10000000	       120 ns/op	      0 B/op	       0 allocs/op
BenchmarkGreenPackMarshal-4              	10000000	       129 ns/op	      0 B/op	       0 allocs/op
BenchmarkGogoprotobufMarshal-4           	10000000	       152 ns/op	      64 B/op	       1 allocs/op
BenchmarkMsgpMarshal-4                   	10000000	       168 ns/op	     128 B/op	       1 allocs/op
BenchmarkGencodeMarshal-4                	10000000	       174 ns/op	      80 B/op	       2 allocs/op
BenchmarkFlatBufferMarshal-4             	 5000000	       379 ns/op	       0 B/op	       0 allocs/op
BenchmarkCapNProtoMarshal-4              	 3000000	       498 ns/op	      56 B/op	       2 allocs/op
BenchmarkGoprotobufMarshal-4             	 3000000	       500 ns/op	     312 B/op	       4 allocs/op
BenchmarkProtobufMarshal-4               	 2000000	       936 ns/op	     200 B/op	       7 allocs/op
BenchmarkGobMarshal-4                    	 2000000	       943 ns/op	      48 B/op	       2 allocs/op
BenchmarkCapNProto2Marshal-4             	 1000000	      1053 ns/op	     436 B/op	       7 allocs/op
BenchmarkHproseMarshal-4                 	 1000000	      1053 ns/op	     479 B/op	       8 allocs/op
BenchmarkBinaryMarshal-4                 	 1000000	      1418 ns/op	     256 B/op	      16 allocs/op
BenchmarkXdrMarshal-4                    	 1000000	      1902 ns/op	     455 B/op	      20 allocs/op
BenchmarkVmihailencoMsgpackMarshal-4     	  500000	      2115 ns/op	     368 B/op	       6 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-4     	 1000000	      2364 ns/op	    2752 B/op	       8 allocs/op
BenchmarkJsonMarshal-4                   	 1000000	      2387 ns/op	     536 B/op	       6 allocs/op
BenchmarkUgorjiCodecBincMarshal-4        	  500000	      2410 ns/op	    2784 B/op	       8 allocs/op
BenchmarkSerealMarshal-4                 	  500000	      2816 ns/op	     912 B/op	      21 allocs/op
```

## Issues


The benchmarks can also be run with validation enabled. Validation does memory allocations so the memory usage will be wrong during a validation run.

```bash
VALIDATE=1 go test -bench='.*' ./
```

Unfortunately, several of the serializers exhibit issues:

1. **(minor)** BSON drops sub-microsecond precision from `time.Time`.
2. **(minor)** Vitess BSON drops sub-microsecond precision from `time.Time`.
3. **(minor)** Ugorji Binc Codec drops the timezone name (eg. "EST" -> "-0500") from `time.Time`.

```
--- FAIL: BenchmarkBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20b999e3621bd773 2016-01-19 14:05:02.469416459 -0800 PST f017c8e9de 4 true 0.20887343719329818}
        &{20b999e3621bd773 2016-01-19 14:05:02.469 -0800 PST f017c8e9de 4 true 0.20887343719329818}
--- FAIL: BenchmarkVitessBsonUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{64204cc8fe408156 2016-01-19 14:05:04.068418034 -0800 PST 95d04e2d49 3 false 0.011931519836294776}
        &{64204cc8fe408156 2016-01-19 22:05:04.068 +0000 UTC 95d04e2d49 3 false 0.011931519836294776}
--- FAIL: BenchmarkUgorjiCodecBincUnmarshal-8
    serialization_benchmarks_test.go:115: unmarshaled object differed:
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 PST 71f3bf4233 0 false 0.8712180830484527}
        &{20a1757ced6b488e 2016-01-19 14:05:15.69474534 -0800 -0800 71f3bf4233 0 false 0.8712180830484527}
```

All other fields are correct however.

Additionally, while not a correctness issue, FlatBuffers, ProtoBuffers and Cap'N'Proto do not
support time types directly. In the benchmarks an int64 value is used to hold a UnixNano timestamp.
