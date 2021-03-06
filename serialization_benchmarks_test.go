package goserbench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"unsafe"

	"zombiezen.com/go/capnproto2"

	"github.com/DeDiS/protobuf"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/binary"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/glycerine/go-capnproto"
	"github.com/gogo/protobuf/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hprose/hprose-go"
	//"github.com/tinylib/msgp/msgp"
	"github.com/ugorji/go/codec"
	//vitessbson "github.com/youtube/vitess/go/bson"
	//"gopkg.in/mgo.v2/bson"
	vmihailenco "gopkg.in/vmihailenco/msgpack.v2"
)

var (
	validate = os.Getenv("VALIDATE")
)

var commonData []*A

func init() {
	commonData = generate()
}

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate() []*A {
	rand.Seed(42)
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

type Serializer interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{}) error
	String() string
}

func benchMarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	//	data := generate()
	data := commonData
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Marshal(data[rand.Intn(len(data))])
	}
}

func cmpTags(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

func cmpAliases(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.StopTimer()
	//data := generate()
	data := commonData
	ser := make([][]byte, len(data))
	for i, d := range data {
		o := s.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

func TestMessage(t *testing.T) {
	println(`
A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.
`)

}

// gopkg.in/vmihailenco/msgpack.v2

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) []byte {
	d, _ := vmihailenco.Marshal(o)
	return d
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func (m VmihailencoMsgpackSerializer) String() string {
	return "vmihailenco-msgpack"
}

func BenchmarkVmihailencoMsgpackMarshal(b *testing.B) {
	benchMarshal(b, VmihailencoMsgpackSerializer{})
}

func BenchmarkVmihailencoMsgpackUnmarshal(b *testing.B) {
	benchUnmarshal(b, VmihailencoMsgpackSerializer{})
}

// encoding/json

type JsonSerializer struct{}

func (m JsonSerializer) Marshal(o interface{}) []byte {
	d, _ := json.Marshal(o)
	return d
}

func (m JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func (j JsonSerializer) String() string {
	return "json"
}

func BenchmarkJsonMarshal(b *testing.B) {
	benchMarshal(b, JsonSerializer{})
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	benchUnmarshal(b, JsonSerializer{})
}

// encoding/gob

type GobSerializer struct {
	b   bytes.Buffer
	enc *gob.Encoder
	dec *gob.Decoder
}

func (g *GobSerializer) Marshal(o interface{}) []byte {
	g.b.Reset()
	err := g.enc.Encode(o)
	if err != nil {
		panic(err)
	}
	return g.b.Bytes()
}

func (g *GobSerializer) Unmarshal(d []byte, o interface{}) error {
	g.b.Reset()
	g.b.Write(d)
	err := g.dec.Decode(o)
	return err
}

func (g GobSerializer) String() string {
	return "gob"
}

func NewGobSerializer() *GobSerializer {
	s := &GobSerializer{}
	s.enc = gob.NewEncoder(&s.b)
	s.dec = gob.NewDecoder(&s.b)
	err := s.enc.Encode(A{})
	if err != nil {
		panic(err)
	}
	var a A
	err = s.dec.Decode(&a)
	if err != nil {
		panic(err)
	}
	return s
}

func BenchmarkGobMarshal(b *testing.B) {
	s := NewGobSerializer()
	benchMarshal(b, s)
}

func BenchmarkGobUnmarshal(b *testing.B) {
	s := NewGobSerializer()
	benchUnmarshal(b, s)
}

// github.com/davecgh/go-xdr/xdr

type XdrSerializer struct{}

func (x XdrSerializer) Marshal(o interface{}) []byte {
	d, _ := xdr.Marshal(o)
	return d
}

func (x XdrSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func (x XdrSerializer) String() string {
	return "xdr"
}

func BenchmarkXdrMarshal(b *testing.B) {
	benchMarshal(b, XdrSerializer{})
}

func BenchmarkXdrUnmarshal(b *testing.B) {
	benchUnmarshal(b, XdrSerializer{})
}

// github.com/ugorji/go/codec

type UgorjiCodecSerializer struct {
	name string
	h    codec.Handle
}

func NewUgorjiCodecSerializer(name string, h codec.Handle) *UgorjiCodecSerializer {
	return &UgorjiCodecSerializer{
		name: name,
		h:    h,
	}
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) []byte {
	var bs []byte
	codec.NewEncoderBytes(&bs, u.h).Encode(o)
	return bs
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	return codec.NewDecoderBytes(d, u.h).Decode(o)
}

func (u *UgorjiCodecSerializer) String() string {
	return "ugorjicodec-" + u.name
}

func BenchmarkUgorjiCodecMsgpackMarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchMarshal(b, s)
}

func BenchmarkUgorjiCodecMsgpackUnmarshal(b *testing.B) {
	s := NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{})
	benchUnmarshal(b, s)
}

func BenchmarkUgorjiCodecBincMarshal(b *testing.B) {
	h := &codec.BincHandle{}
	h.AsSymbols = codec.AsSymbolNone
	s := NewUgorjiCodecSerializer("binc", h)
	benchMarshal(b, s)
}

func BenchmarkUgorjiCodecBincUnmarshal(b *testing.B) {
	h := &codec.BincHandle{}
	h.AsSymbols = codec.AsSymbolNone
	s := NewUgorjiCodecSerializer("binc", h)
	benchUnmarshal(b, s)
}

// github.com/Sereal/Sereal/Go/sereal

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) []byte {
	d, _ := sereal.Marshal(o)
	return d
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func (s SerealSerializer) String() string {
	return "sereal"
}

func BenchmarkSerealMarshal(b *testing.B) {
	benchMarshal(b, SerealSerializer{})
}

func BenchmarkSerealUnmarshal(b *testing.B) {
	benchUnmarshal(b, SerealSerializer{})
}

// github.com/alecthomas/binary

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) []byte {
	d, _ := binary.Marshal(o)
	return d
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func (b BinarySerializer) String() string {
	return "binary"
}

func BenchmarkBinaryMarshal(b *testing.B) {
	benchMarshal(b, BinarySerializer{})
}

func BenchmarkBinaryUnmarshal(b *testing.B) {
	benchUnmarshal(b, BinarySerializer{})
}

// github.com/google/flatbuffers/go

type FlatBufferSerializer struct {
	builder *flatbuffers.Builder
}

var fbSer = &FlatBufferSerializer{flatbuffers.NewBuilder(0)}

func BenchmarkFlatBufferMarshal(b *testing.B) {
	b.StopTimer()
	data := commonData
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fbSer.Marshal(data[rand.Intn(len(data))])
	}
}

func (s *FlatBufferSerializer) Marshal(o interface{}) []byte {
	a := o.(*A)
	builder := s.builder

	builder.Reset()

	name := builder.CreateString(a.Name)
	phone := builder.CreateString(a.Phone)

	FlatBufferAStart(builder)
	FlatBufferAAddName(builder, name)
	FlatBufferAAddPhone(builder, phone)
	FlatBufferAAddBirthDay(builder, a.BirthDay.UnixNano())
	FlatBufferAAddSiblings(builder, int32(a.Siblings))
	var spouse byte
	if a.Spouse {
		spouse = byte(1)
	}
	FlatBufferAAddSpouse(builder, spouse)
	FlatBufferAAddMoney(builder, a.Money)
	builder.Finish(FlatBufferAEnd(builder))
	return builder.Bytes[builder.Head():]
}

func (s *FlatBufferSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	o := FlatBufferA{}
	o.Init(d, flatbuffers.GetUOffsetT(d))
	a.Name = string(o.Name())
	a.BirthDay = time.Unix(o.BirthDay(), 0)
	a.Phone = string(o.Phone())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse() == byte(1)
	a.Money = o.Money()
	return nil
}

func (s *FlatBufferSerializer) String() string {
	return "FlatBuffer"
}

func BenchmarkFlatBuffersUnmarshal(b *testing.B) {
	fbUn := &FlatBufferSerializer{flatbuffers.NewBuilder(0)}

	b.StopTimer()
	//data := generate()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		o := fbSer.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}

	o := &A{}
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		err := fbUn.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", fbUn, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/glycerine/go-capnproto

type CapNProtoSerializer struct {
	buf []byte
	out *bytes.Buffer
}

func (x *CapNProtoSerializer) Marshal(o interface{}) []byte {
	a := o.(*A)
	s := capn.NewBuffer(x.buf)
	c := NewRootCapnpA(s)
	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)
	x.out.Reset()
	s.WriteTo(x.out)
	x.buf = []byte(s.Data)[:0]
	return x.out.Bytes()
}

// the multi.Segment[0] here can be safely reused on each Unmarshal
var multi = capn.NewSingleSegmentMultiBuffer()

func (x *CapNProtoSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	capn.ReadFromMemoryZeroCopyNoAlloc(d, multi)
	s := multi.Segments[0]
	o := ReadRootCapnpA(s)

	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()

	// This is tuned for performance, to demonstrate how
	// to do zero-allocation deserialization in CapNProto, which
	// is the one of the big reasons CapNProto was created
	// as an alternative to Protocol Buffers.
	//
	// In general you should only use unsafe if you can
	// guarantee those bytes won't be changing
	// out from under you when they are accessed afterwards [1].
	// Quite frequently this is possible if you own the
	// []byte buffer, process each message in turn, and then
	// never refer to the string again. This can make a huge
	// perf difference when dealing with large data.
	//
	// reference
	// [1] https://groups.google.com/d/topic/golang-nuts/ENgbUzYvCuU

	// 1st possible allocation: Avoided using unsafe; could
	// also be avoided by storing Name as []byte instead
	// of string in the first place
	nameBytes := o.NameBytes()
	a.Name = *(*string)(unsafe.Pointer(&nameBytes))
	// a.Name = string(nameBytes) // would allocate for the string conversion.

	// 2nd possible allocation: also a string -- same comments
	// as above apply.
	phoneBytes := o.PhoneBytes()
	a.Phone = *(*string)(unsafe.Pointer(&phoneBytes))
	// a.Phone = string(phoneBytes) // would allocate for the string conversion.

	return nil
}

func (x *CapNProtoSerializer) String() string {
	return "CapNProto"
}

func BenchmarkCapNProtoMarshal(b *testing.B) {
	benchMarshal(b, &CapNProtoSerializer{nil, &bytes.Buffer{}})
}

var capnSer = &CapNProtoSerializer{nil, &bytes.Buffer{}}

func BenchmarkCapNProtoUnmarshal(b *testing.B) {
	benchUnmarshalCapn(b, capnSer) // 0 allocations
	// benchUnmarshal(b, capnSer) // costs 80 bytes allocated for the interface{}
}

// zombiezen.com/go/capnproto2

type CapNProto2Serializer struct {
	arena capnp.Arena
}

func (x *CapNProto2Serializer) Marshal(o interface{}) []byte {
	a := o.(*A)
	m, s, _ := capnp.NewMessage(x.arena)
	c, _ := NewRootCapnp2A(s)
	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)
	b, _ := m.Marshal()
	return b
}

func (x *CapNProto2Serializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*A)
	m, _ := capnp.Unmarshal(d)
	o, _ := ReadRootCapnp2A(m)
	a.Name, _ = o.Name()
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone, _ = o.Phone()
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func (x *CapNProto2Serializer) String() string {
	return "CapNProto2"
}

func BenchmarkCapNProto2Marshal(b *testing.B) {
	benchMarshal(b, &CapNProto2Serializer{capnp.SingleSegment(nil)})
}

func BenchmarkCapNProto2Unmarshal(b *testing.B) {
	benchUnmarshal(b, &CapNProto2Serializer{capnp.SingleSegment(nil)})
}

// github.com/hprose/hprose-go/io

type HproseSerializer struct {
	writer *hprose.Writer
	reader *hprose.Reader
}

func (s *HproseSerializer) Marshal(o interface{}) []byte {
	a := o.(*A)
	writer := s.writer
	buf := writer.Stream.(*bytes.Buffer)
	l := buf.Len()
	writer.WriteString(a.Name)
	writer.WriteTime(a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt64(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat64(a.Money)
	return buf.Bytes()[l:]
}

func (s *HproseSerializer) Unmarshal(d []byte, i interface{}) error {
	o := i.(*A)
	reader := s.reader
	reader.Stream = &hprose.BytesReader{d, 0}
	o.Name, _ = reader.ReadString()
	o.BirthDay, _ = reader.ReadDateTime()
	o.Phone, _ = reader.ReadString()
	o.Siblings, _ = reader.ReadInt()
	o.Spouse, _ = reader.ReadBool()
	o.Money, _ = reader.ReadFloat64()
	return nil
}

func (s *HproseSerializer) String() string {
	return "Hprose"
}

func BenchmarkHproseMarshal(b *testing.B) {
	buf := new(bytes.Buffer)
	writer := hprose.NewWriter(buf, true)
	benchMarshal(b, &HproseSerializer{writer: writer})
}

func BenchmarkHproseUnmarshal(b *testing.B) {
	buf := new(bytes.Buffer)
	reader := hprose.NewReader(buf, true)
	bufw := new(bytes.Buffer)
	writer := hprose.NewWriter(bufw, true)
	benchUnmarshal(b, &HproseSerializer{writer: writer, reader: reader})
}

// github.com/DeDiS/protobuf

type ProtobufSerializer struct{}

func (m ProtobufSerializer) Marshal(o interface{}) []byte {
	d, _ := protobuf.Encode(o)
	return d
}

func (m ProtobufSerializer) Unmarshal(d []byte, o interface{}) error {
	return protobuf.Decode(d, o)
}

func (j ProtobufSerializer) String() string {
	return "protobuf"
}

func BenchmarkProtobufMarshal(b *testing.B) {
	benchMarshal(b, ProtobufSerializer{})
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	benchUnmarshal(b, ProtobufSerializer{})
}

// github.com/golang/protobuf

func generateProto() []*ProtoBufA {
	rand.Seed(42)
	a := make([]*ProtoBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &ProtoBufA{
			Name:     proto.String(randString(16)),
			BirthDay: proto.Int64(time.Now().UnixNano()),
			Phone:    proto.String(randString(10)),
			Siblings: proto.Int32(rand.Int31n(5)),
			Spouse:   proto.Bool(rand.Intn(2) == 1),
			Money:    proto.Float64(rand.Float64()),
		})
	}
	return a
}

func BenchmarkGoprotobufMarshal(b *testing.B) {
	b.StopTimer()
	data := generateProto()
	//data := commonData
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(data[rand.Intn(len(data))])
	}
}

func BenchmarkGoprotobufUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateProto()
	//data := commonData
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = proto.Marshal(d)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &ProtoBufA{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := *o.Name == *i.Name && *o.Phone == *i.Phone && *o.Siblings == *i.Siblings && *o.Spouse == *i.Spouse && *o.Money == *i.Money && *o.BirthDay == *i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/gogo/protobuf/proto

func generateGogoProto() []*GogoProtoBufA {
	rand.Seed(42)
	a := make([]*GogoProtoBufA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GogoProtoBufA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: rand.Int31n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGogoprotobufMarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoProto()
	//data := commonData
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(data[rand.Intn(len(data))])
	}
}

func BenchmarkGogoprotobufUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGogoProto()
	//data := commonData
	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = proto.Marshal(d)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GogoProtoBufA{}
		err := proto.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/andyleap/gencode

func generateGencode() []*GencodeA {
	rand.Seed(42)
	a := make([]*GencodeA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &GencodeA{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Int63n(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGencodeMarshal(b *testing.B) {
	b.StopTimer()
	data := generateGencode()
	//data := commonData
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].Marshal(nil)
	}
}

func BenchmarkGencodeUnmarshal(b *testing.B) {
	b.StopTimer()
	data := generateGencode()
	//data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.Marshal(nil)
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &GencodeA{}
		_, err := o.Unmarshal(ser[n])
		if err != nil {
			b.Fatalf("goprotobuf failed to unmarshal: %s (%s)", err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// the interface conversion in benchUnmarshal allocates 80 bytes; we
// can avoid that overhead by avoiding the interface{} and
// using the pointer s *CapNProtoSerializer directly.
func benchUnmarshalCapn(b *testing.B, s *CapNProtoSerializer) {
	b.StopTimer()
	//data := generate()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		o := s.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}

	o := &A{}
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String() //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

// github.com/tinylib/msgp

func BenchmarkMsgpMarshal(b *testing.B) {
	b.StopTimer()
	//data := generateZebraPack()
	data := commonData

	//b.ReportAllocs()
	_, err := data[rand.Intn(len(data))].MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	/*
		// compute bytes written
		writ := 0
		for i := range data {
			o, _ := data[i].MarshalMsg(nil)
			writ += len(o)
		}
		b.SetBytes(int64(writ / len(data)))
	*/
	buf := make([]byte, 0, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].MarshalMsg(buf)
	}
}

func BenchmarkMsgpUnmarshal(b *testing.B) {
	b.StopTimer()
	//	data := generateZebraPack()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.MarshalMsg(nil)
	}
	o := &A{}
	z := A{}
	var n int
	var err error
	b.ResetTimer()
	//b.ReportAllocs()

	/*
		// compute bytes read
		red := 0
		for i := range ser {
			red += len(ser[i])
		}
		b.SetBytes(int64(red / len(ser)))
	*/
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rand.Intn(len(ser))
		*o = z // clear
		_, err = o.UnmarshalMsg(ser[n])
		if err != nil {
			b.Fatalf("zebrapack failed to unmarshal: %s (%s)", err, ser[n])
		}

		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}

	}
}

// github.com/glycerine/zebrapack

func BenchmarkZebraPackUnmarshal(b *testing.B) {
	b.StopTimer()
	//	data := generateZebraPack()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.ZMarshalMsg(nil)
		//ser[i], _ = proto.Marshal(d)
	}
	o := &A{}
	z := A{}
	var n int
	var err error
	b.ResetTimer()
	b.ReportAllocs()

	// compute bytes read
	red := 0
	for i := range ser {
		red += len(ser[i])
	}
	b.SetBytes(int64(red / len(ser)))

	b.SetBytes(int64(len(ser[0])))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rand.Intn(len(ser))
		*o = z // clear
		_, err = o.ZUnmarshalMsg(ser[n])
		if err != nil {
			b.Fatalf("zebrapack failed to unmarshal: %s (%s)", err, ser[n])
		}

		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}

	}
}

func generateZebraPack() []*A {
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkZebraPackMarshal(b *testing.B) {
	b.StopTimer()
	//data := generateZebraPack()
	data := commonData

	b.ReportAllocs()
	_, err := data[rand.Intn(len(data))].ZMarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	// compute bytes written
	writ := 0
	for i := range data {
		o, _ := data[i].ZMarshalMsg(nil)
		writ += len(o)
	}
	b.SetBytes(int64(writ / len(data)))
	buf := make([]byte, 0, 512)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].ZMarshalMsg(buf)
	}
}

func BenchmarkGreenPackUnmarshal(b *testing.B) {
	b.StopTimer()
	//	data := generateGreenPack()
	data := commonData

	ser := make([][]byte, len(data))
	for i, d := range data {
		ser[i], _ = d.ZMarshalMsg(nil)
		//ser[i], _ = proto.Marshal(d)
	}
	o := &A{}
	z := A{}
	var n int
	var err error
	b.ResetTimer()
	b.ReportAllocs()

	// compute bytes read
	red := 0
	for i := range ser {
		red += len(ser[i])
	}
	b.SetBytes(int64(red / len(ser)))

	b.SetBytes(int64(len(ser[0])))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		n = rand.Intn(len(ser))
		*o = z // clear
		_, err = o.ZUnmarshalMsg(ser[n])
		if err != nil {
			b.Fatalf("zebrapack failed to unmarshal: %s (%s)", err, ser[n])
		}

		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}

	}
}

func generateGreenPack() []*A {
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func BenchmarkGreenPackMarshal(b *testing.B) {
	b.StopTimer()
	//data := generateGreenPack()
	data := commonData

	b.ReportAllocs()
	_, err := data[rand.Intn(len(data))].GreenMarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	// compute bytes written
	writ := 0
	for i := range data {
		o, _ := data[i].GreenMarshalMsg(nil)
		writ += len(o)
	}
	b.SetBytes(int64(writ / len(data)))
	buf := make([]byte, 0, 512)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		data[rand.Intn(len(data))].GreenMarshalMsg(buf)
	}
}
