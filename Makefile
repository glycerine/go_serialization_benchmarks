all: FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go structdef.capnp.go structdef.capnp2.go gencode.schema.gen.go zebrapack_gen.go

FlatBufferA.go: flatbuffers-structdef.fbs
	flatc -g flatbuffers-structdef.fbs
	mv flatbuffersmodels/FlatBufferA.go FlatBufferA.go
	rmdir flatbuffersmodels
	sed -i '' 's/flatbuffersmodels/goserbench/' FlatBufferA.go

msgp_gen.go zebrapack_gen.go: structdef.go
	go generate
	perl -pi -e 's/MarshalMsg/ZMarshalMsg/g' zebrapack_gen.go
	perl -pi -e 's/Msgsize/ZMsgsize/g' zebrapack_gen.go
	perl -pi -e 's/UnmarshalMsg/ZUnmarshalMsg/g' zebrapack_gen.go

structdef-gogo.pb.go: structdef-gogo.proto
	protoc --gogofaster_out=. -I. -I${GOPATH}/src  -I${GOPATH}/src/github.com/gogo/protobuf/protobuf structdef-gogo.proto

structdef.pb.go: structdef.proto
	protoc --go_out=. structdef.proto


structdef.capnp2.go: structdef.capnp2
	go get -u zombiezen.com/go/capnproto2/...
	capnp compile -I${GOPATH}/src -ogo structdef.capnp2

structdef.capnp.go: structdef.capnp
	go get -u github.com/glycerine/go-capnproto/capnpc-go
	capnp compile -I${GOPATH}/src -ogo structdef.capnp

gencode.schema.gen.go: gencode.schema
	go get -u github.com/andyleap/gencode
	gencode go -schema=gencode.schema -package=goserbench

.PHONY: clean
clean:
	rm -f FlatBufferA.go msgp_gen.go structdef-gogo.pb.go structdef.pb.go vitess_test.go structdef.capnp.go structdef.capnp2.go gencode.schema.gen.go

.PHONY: install
install:
	go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
	go get -u github.com/gogo/protobuf/gogoproto
	go get -u github.com/golang/protobuf/protoc-gen-go
	#go get -u github.com/tinylib/msgp
	#go get -u github.com/andyleap/gencode

	go get -u github.com/DeDiS/protobuf
	go get -u github.com/Sereal/Sereal/Go/sereal
	go get -u github.com/alecthomas/binary
	go get -u github.com/davecgh/go-xdr/xdr
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/google/flatbuffers/go
	#go get -u github.com/tinylib/msgp/msgp
	go get -u github.com/ugorji/go/codec
	go get -u gopkg.in/vmihailenco/msgpack.v2
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/hprose/hprose-go/io

benchit:
	go test -v -bench=. -benchtime 1s -timeout 60m -run - |tee res
	cat res|grep Un | sort -nk 3 > read.speed.txt
	cat res|grep -v Un | sort -nk 3 > write.speed.txt
