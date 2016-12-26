package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// fieldsNotEmpty supports omitempty tags
func (z *A) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[0] = (len(z.Name) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.BirthDay.IsZero()) // time.Time, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Siblings == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (!z.Spouse) // bool, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.Money == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// ZMarshalMsg implements msgp.Marshaler
func (z *A) ZMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.ZMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "Name"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[1] {
		// zid 1 for "BirthDay"
		o = append(o, 0x1)
		o = msgp.AppendTime(o, z.BirthDay)
	}

	if !empty[2] {
		// zid 2 for "Phone"
		o = append(o, 0x2)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// zid 3 for "Siblings"
		o = append(o, 0x3)
		o = msgp.AppendInt(o, z.Siblings)
	}

	if !empty[4] {
		// zid 4 for "Spouse"
		o = append(o, 0x4)
		o = msgp.AppendBool(o, z.Spouse)
	}

	if !empty[5] {
		// zid 5 for "Money"
		o = append(o, 0x5)
		o = msgp.AppendFloat64(o, z.Money)
	}

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
	const maxFields0zujs = 6

	// -- templateZUnmarshalMsgZid starts here--
	var totalEncodedFields0zujs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zujs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zujs := totalEncodedFields0zujs
	missingFieldsLeft0zujs := maxFields0zujs - totalEncodedFields0zujs

	var nextMiss0zujs int = -1
	var found0zujs [maxFields0zujs]bool
	var curField0zujs int

doneWithStruct0zujs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zujs > 0 || missingFieldsLeft0zujs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zujs, missingFieldsLeft0zujs, msgp.ShowFound(found0zujs[:]), unmarshalMsgFieldOrder0zujs)
		if encodedFieldsLeft0zujs > 0 {
			encodedFieldsLeft0zujs--
			curField0zujs, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zujs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zujs = 0
			}
			for nextMiss0zujs < maxFields0zujs && (found0zujs[nextMiss0zujs] || unmarshalMsgFieldSkip0zujs[nextMiss0zujs]) {
				nextMiss0zujs++
			}
			if nextMiss0zujs == maxFields0zujs {
				// filled all the empty fields!
				break doneWithStruct0zujs
			}
			missingFieldsLeft0zujs--
			curField0zujs = nextMiss0zujs
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zujs)
		switch curField0zujs {
		// -- templateZUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Name"
			found0zujs[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "BirthDay"
			found0zujs[1] = true
			z.BirthDay, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Phone"
			found0zujs[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 3:
			// zid 3 for "Siblings"
			found0zujs[3] = true
			z.Siblings, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 4:
			// zid 4 for "Spouse"
			found0zujs[4] = true
			z.Spouse, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case 5:
			// zid 5 for "Money"
			found0zujs[5] = true
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
	if nextMiss0zujs != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zujs = []string{"Name", "BirthDay", "Phone", "Siblings", "Spouse", "Money"}

var unmarshalMsgFieldSkip0zujs = []bool{false, false, false, false, false, false}

// ZMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) ZMsgsize() (s int) {
	s = 1 + 1 + msgp.StringPrefixSize + len(z.Name) + 1 + msgp.TimeSize + 1 + msgp.StringPrefixSize + len(z.Phone) + 1 + msgp.IntSize + 1 + msgp.BoolSize + 1 + msgp.Float64Size
	return
}
