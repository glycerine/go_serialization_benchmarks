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
  BenchmarkGencodeUnmarshal-4              	10000000	       207 ns/op	     112 B/op	       3 allocs/op
* BenchmarkZebraPackUnmarshal-4            	10000000	       222 ns/op	       0 B/op	       0 allocs/op
  BenchmarkGogoprotobufUnmarshal-4         	 5000000	       253 ns/op	     112 B/op	       3 allocs/op
  BenchmarkCapNProtoUnmarshal-4            	 5000000	       267 ns/op	       0 B/op	       0 allocs/op
  BenchmarkFlatBuffersUnmarshal-4          	 5000000	       277 ns/op	     112 B/op	       3 allocs/op
  BenchmarkMsgpUnmarshal-4                 	 5000000	       354 ns/op	     112 B/op	       3 allocs/op
  BenchmarkGoprotobufUnmarshal-4           	 2000000	       655 ns/op	     432 B/op	       9 allocs/op
  BenchmarkProtobufUnmarshal-4             	 2000000	       734 ns/op	     192 B/op	      10 allocs/op
  BenchmarkGobUnmarshal-4                  	 1000000	      1066 ns/op	     112 B/op	       3 allocs/op
  BenchmarkHproseUnmarshal-4               	 1000000	      1066 ns/op	     320 B/op	      10 allocs/op
  BenchmarkCapNProto2Unmarshal-4           	 1000000	      1399 ns/op	     608 B/op	      12 allocs/op
  BenchmarkBinaryUnmarshal-4               	 1000000	      1562 ns/op	     336 B/op	      22 allocs/op
  BenchmarkXdrUnmarshal-4                  	 1000000	      1571 ns/op	     239 B/op	      11 allocs/op
  BenchmarkVmihailencoMsgpackUnmarshal-4   	 1000000	      2084 ns/op	     384 B/op	      13 allocs/op
  BenchmarkUgorjiCodecMsgpackUnmarshal-4   	  500000	      2593 ns/op	    3008 B/op	       6 allocs/op
  BenchmarkUgorjiCodecBincUnmarshal-4      	  500000	      2769 ns/op	    3168 B/op	       9 allocs/op
  BenchmarkSerealUnmarshal-4               	  300000	      3426 ns/op	    1008 B/op	      34 allocs/op
  BenchmarkJsonUnmarshal-4                 	  300000	      4224 ns/op	     495 B/op	       8 allocs/op
```

* sorted by write speed


```
benchmark                                       iter           time/iter          bytes alloc      allocs
---------                                       ----           ---------          -----------      ------
  BenchmarkGogoprotobufMarshal-4           	10000000	       151 ns/op	      64 B/op	       1 allocs/op
* BenchmarkZebraPackMarshal-4              	10000000	       151 ns/op	      80 B/op	       1 allocs/op
  BenchmarkGencodeMarshal-4                	10000000	       168 ns/op	      80 B/op	       2 allocs/op
  BenchmarkMsgpMarshal-4                   	10000000	       191 ns/op	     128 B/op	       1 allocs/op
  BenchmarkFlatbuffersMarshal-4            	 5000000	       352 ns/op	       0 B/op	       0 allocs/op
  BenchmarkCapNProtoMarshal-4              	 3000000	       510 ns/op	      56 B/op	       2 allocs/op
  BenchmarkGoprotobufMarshal-4             	 2000000	       535 ns/op	     312 B/op	       4 allocs/op
  BenchmarkGobMarshal-4                    	 2000000	       918 ns/op	      48 B/op	       2 allocs/op
  BenchmarkProtobufMarshal-4               	 2000000	       920 ns/op	     200 B/op	       7 allocs/op
  BenchmarkCapNProto2Marshal-4             	 1000000	      1008 ns/op	     436 B/op	       7 allocs/op
  BenchmarkHproseMarshal-4                 	 1000000	      1013 ns/op	     476 B/op	       8 allocs/op
  BenchmarkBinaryMarshal-4                 	 1000000	      1370 ns/op	     256 B/op	      16 allocs/op
  BenchmarkXdrMarshal-4                    	 1000000	      1810 ns/op	     456 B/op	      21 allocs/op
  BenchmarkVmihailencoMsgpackMarshal-4     	 1000000	      1997 ns/op	     368 B/op	       6 allocs/op
  BenchmarkJsonMarshal-4                   	 1000000	      2328 ns/op	     536 B/op	       6 allocs/op
  BenchmarkUgorjiCodecBincMarshal-4        	  500000	      2523 ns/op	    2785 B/op	       8 allocs/op
  BenchmarkUgorjiCodecMsgpackMarshal-4     	  500000	      2578 ns/op	    2752 B/op	       8 allocs/op
  BenchmarkSerealMarshal-4                 	  500000	      3262 ns/op	     912 B/op	      21 allocs/op
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
