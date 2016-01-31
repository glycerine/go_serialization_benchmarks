package goserbench

import (
	"math/rand"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	structdef "github.com/alecthomas/go_serialization_benchmarks/gen-go/structdef"
)

// git.apache.org/thrift.git/lib/go/thrift

type ThriftCompactSerializer struct {
	mem    *thrift.TMemoryBuffer
	pro    thrift.TProtocol
	tser   *thrift.TSerializer
	tdeser *thrift.TDeserializer
}

func (g *ThriftCompactSerializer) Marshal(a *structdef.ThriftA) []byte {
	by, _ := g.tser.Write(a)
	return by
}

func (g *ThriftCompactSerializer) Unmarshal(d []byte, a *structdef.ThriftA) error {
	g.mem.Close() // resets buffer
	return g.tdeser.Read(a, d)
}

func (g ThriftCompactSerializer) String() string {
	return "thrift-compact"
}

func NewThriftCompactSerializer() *ThriftCompactSerializer {
	s := &ThriftCompactSerializer{}
	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTCompactProtocolFactory().GetProtocol(t)
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

var thriftCompactSer = NewThriftCompactSerializer()

func BenchmarkThriftCompactMarshal(b *testing.B) {
	benchMarshalThriftCompact(b, thriftCompactSer)
}

func BenchmarkThriftCompactUnmarshal(b *testing.B) {
	benchUnmarshalThriftCompact(b, thriftCompactSer)
}

func benchUnmarshalThriftCompact(b *testing.B, s *ThriftCompactSerializer) {
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

func benchMarshalThriftCompact(b *testing.B, s *ThriftCompactSerializer) {
	b.StopTimer()
	data := generateThrift()
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Marshal(data[rand.Intn(len(data))])
	}
}
