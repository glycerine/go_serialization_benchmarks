package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// GreenfieldsNotEmpty supports omitempty tags
func (z *A) GreenFieldsNotEmpty(isempty []bool) uint32 {
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

// GreenMarshalMsg implements msgp.Marshaler
func (z *A) GreenMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.GreenMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.GreenFieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "Name_zid00_str"
		o = append(o, 0xae, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[1] {
		// string "BirthDay_zid01_tim"
		o = append(o, 0xb2, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
		o = msgp.AppendTime(o, z.BirthDay)
	}

	if !empty[2] {
		// string "Phone_zid02_str"
		o = append(o, 0xaf, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Siblings_zid03_int"
		o = append(o, 0xb2, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Siblings)
	}

	if !empty[4] {
		// string "Spouse_zid04_boo"
		o = append(o, 0xb0, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.Spouse)
	}

	if !empty[5] {
		// string "Money_zid05_f64"
		o = append(o, 0xaf, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x66, 0x36, 0x34)
		o = msgp.AppendFloat64(o, z.Money)
	}

	return
}

// GreenUnmarshalMsg implements msgp.Unmarshaler
func (z *A) GreenUnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.GreenUnmarshalMsgWithCfg(bts, cfg)
}
func (z *A) GreenUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zslc = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zslc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zslc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zslc := totalEncodedFields0zslc
	missingFieldsLeft0zslc := maxFields0zslc - totalEncodedFields0zslc

	var nextMiss0zslc int32 = -1
	var found0zslc [maxFields0zslc]bool
	var curField0zslc string

doneWithStruct0zslc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zslc > 0 || missingFieldsLeft0zslc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zslc, missingFieldsLeft0zslc, msgp.ShowFound(found0zslc[:]), unmarshalMsgFieldOrder0zslc)
		if encodedFieldsLeft0zslc > 0 {
			encodedFieldsLeft0zslc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zslc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zslc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zslc = 0
			}
			for nextMiss0zslc < maxFields0zslc && (found0zslc[nextMiss0zslc] || unmarshalMsgFieldSkip0zslc[nextMiss0zslc]) {
				nextMiss0zslc++
			}
			if nextMiss0zslc == maxFields0zslc {
				// filled all the empty fields!
				break doneWithStruct0zslc
			}
			missingFieldsLeft0zslc--
			curField0zslc = unmarshalMsgFieldOrder0zslc[nextMiss0zslc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zslc)
		switch curField0zslc {
		// -- templateUnmarshalMsg ends here --

		case "Name_zid00_str":
			found0zslc[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "BirthDay_zid01_tim":
			found0zslc[1] = true
			z.BirthDay, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "Phone_zid02_str":
			found0zslc[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Siblings_zid03_int":
			found0zslc[3] = true
			z.Siblings, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Spouse_zid04_boo":
			found0zslc[4] = true
			z.Spouse, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "Money_zid05_f64":
			found0zslc[5] = true
			z.Money, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zslc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of A
var unmarshalMsgFieldOrder0zslc = []string{"Name_zid00_str", "BirthDay_zid01_tim", "Phone_zid02_str", "Siblings_zid03_int", "Spouse_zid04_boo", "Money_zid05_f64"}

var unmarshalMsgFieldSkip0zslc = []bool{false, false, false, false, false, false}

// GreenMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) GreenMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.Name) + 19 + msgp.TimeSize + 16 + msgp.StringPrefixSize + len(z.Phone) + 19 + msgp.IntSize + 17 + msgp.BoolSize + 16 + msgp.Float64Size
	return
}
