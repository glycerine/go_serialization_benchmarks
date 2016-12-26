package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// fieldsNotEmpty supports omitempty tags
func (z *A) fieldsNotEmpty(isempty []bool) uint32 {
	return 6
}

// ZMarshalMsg implements msgp.Marshaler
func (z *A) ZMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.ZMsgsize())
	// map header, size 6
	// zid 0 for "Name"
	o = append(o, 0x86, 0x0)
	o = msgp.AppendString(o, z.Name)
	// zid 1 for "BirthDay"
	o = append(o, 0x1)
	o = msgp.AppendTime(o, z.BirthDay)
	// zid 2 for "Phone"
	o = append(o, 0x2)
	o = msgp.AppendString(o, z.Phone)
	// zid 3 for "Siblings"
	o = append(o, 0x3)
	o = msgp.AppendInt(o, z.Siblings)
	// zid 4 for "Spouse"
	o = append(o, 0x4)
	o = msgp.AppendBool(o, z.Spouse)
	// zid 5 for "Money"
	o = append(o, 0x5)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// ZUnmarshalMsg implements msgp.Unmarshaler
func (z *A) ZUnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.ZUnmarshalMsgWithCfg(bts, cfg)
}
func (z *A) ZUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0znxg = 6

	// -- templateZUnmarshalMsgZid starts here--
	var totalEncodedFields0znxg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0znxg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0znxg := totalEncodedFields0znxg
	missingFieldsLeft0znxg := maxFields0znxg - totalEncodedFields0znxg

	var nextMiss0znxg int = -1
	var found0znxg [maxFields0znxg]bool
	var curField0znxg int

doneWithStruct0znxg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0znxg > 0 || missingFieldsLeft0znxg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0znxg, missingFieldsLeft0znxg, msgp.ShowFound(found0znxg[:]), unmarshalMsgFieldOrder0znxg)
		if encodedFieldsLeft0znxg > 0 {
			encodedFieldsLeft0znxg--
			curField0znxg, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0znxg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0znxg = 0
			}
			for nextMiss0znxg < maxFields0znxg && (found0znxg[nextMiss0znxg] || unmarshalMsgFieldSkip0znxg[nextMiss0znxg]) {
				nextMiss0znxg++
			}
			if nextMiss0znxg == maxFields0znxg {
				// filled all the empty fields!
				break doneWithStruct0znxg
			}
			missingFieldsLeft0znxg--
			curField0znxg = nextMiss0znxg
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0znxg)
		switch curField0znxg {
		// -- templateZUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Name"
			found0znxg[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "BirthDay"
			found0znxg[1] = true
			z.BirthDay, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Phone"
			found0znxg[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 3:
			// zid 3 for "Siblings"
			found0znxg[3] = true
			z.Siblings, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "Spouse"
			found0znxg[4] = true
			z.Spouse, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "Money"
			found0znxg[5] = true
			z.Money, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0znxg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0znxg = []string{"Name", "BirthDay", "Phone", "Siblings", "Spouse", "Money"}

var unmarshalMsgFieldSkip0znxg = []bool{false, false, false, false, false, false}

// ZMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) ZMsgsize() (s int) {
	s = 1 + 1 + msgp.StringPrefixSize + len(z.Name) + 1 + msgp.TimeSize + 1 + msgp.StringPrefixSize + len(z.Phone) + 1 + msgp.IntSize + 1 + msgp.BoolSize + 1 + msgp.Float64Size
	return
}
