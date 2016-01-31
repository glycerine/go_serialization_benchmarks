package goserbench

import (
	"math/rand"
	"testing"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	structdef "github.com/alecthomas/go_serialization_benchmarks/gen-go/structdef"
)

// git.apache.org/thrift.git/lib/go/thrift

type ThriftBinarySerializer struct {
	mem    *thrift.TMemoryBuffer
	pro    thrift.TProtocol
	tser   *thrift.TSerializer
	tdeser *thrift.TDeserializer
}

func (g *ThriftBinarySerializer) Marshal(a *structdef.ThriftA) []byte {
	by, _ := g.tser.Write(a)
	return by
}

func (g *ThriftBinarySerializer) Unmarshal(d []byte, a *structdef.ThriftA) error {
	g.mem.Close() // resets buffer
	return g.tdeser.Read(a, d)
}

func (g ThriftBinarySerializer) String() string {
	return "thrift-binary"
}

func NewThriftBinarySerializer() *ThriftBinarySerializer {
	s := &ThriftBinarySerializer{}
	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t)
	s.mem = t
	s.pro = p

	s.tser = &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}
	s.tdeser = &thrift.TDeserializer{
		Transport: t,
		Protocol:  p,
	}

	return s
}

var thriftBinarySer = NewThriftBinarySerializer()

func BenchmarkThriftBinaryMarshal(b *testing.B) {
	benchMarshalThriftBinary(b, thriftBinarySer)
}

func BenchmarkThriftBinaryUnmarshal(b *testing.B) {
	benchUnmarshalThriftBinary(b, thriftBinarySer)
}

func benchUnmarshalThriftBinary(b *testing.B, s *ThriftBinarySerializer) {
	b.StopTimer()
	data := generateThrift()
	ser := make([][]byte, len(data))
	for i, d := range data {
		o := s.Marshal(d)
		t := make([]byte, len(o))
		copy(t, o)
		ser[i] = t
	}

	o := &structdef.ThriftA{}
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
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay == i.BirthDay //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}

func generateThrift() []*structdef.ThriftA {
	a := make([]*structdef.ThriftA, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &structdef.ThriftA{
			Name:     randString(16),
			BirthDay: time.Now().UnixNano(),
			Phone:    randString(10),
			Siblings: int32(rand.Intn(5)),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func benchMarshalThriftBinary(b *testing.B, s *ThriftBinarySerializer) {
	b.StopTimer()
	data := generateThrift()
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Marshal(data[rand.Intn(len(data))])
	}
}
