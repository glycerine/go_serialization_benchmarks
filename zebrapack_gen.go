package goserbench

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

// MarshalMsg implements msgp.Marshaler
func (z *A) ZebraPackMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.ZebraPackMsgsize())
	// map header, size 6
	// string "Name"
	o = append(o, 0x86, 0x00) // 0x86 is the map head for fixmap
	//o = append(o, 0x86, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "BirthDay"
	o = append(o, 0x01)
	//o = append(o, 0xa8, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79)
	o = msgp.AppendTime(o, z.BirthDay)
	// string "Phone"
	o = append(o, 0x02)
	//o = append(o, 0xa5, 0x50, 0x68, 0x6f, 0x6e, 0x65)
	o = msgp.AppendString(o, z.Phone)
	// string "Siblings"
	o = append(o, 0x03)
	//o = append(o, 0xa8, 0x53, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendInt(o, z.Siblings)
	// string "Spouse"
	o = append(o, 0x04)
	//o = append(o, 0xa6, 0x53, 0x70, 0x6f, 0x75, 0x73, 0x65)
	o = msgp.AppendBool(o, z.Spouse)
	// string "Money"
	o = append(o, 0x05)
	//o = append(o, 0xa5, 0x4d, 0x6f, 0x6e, 0x65, 0x79)
	o = msgp.AppendFloat64(o, z.Money)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *A) ZebraPackUnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var fieldId int64
	var zxvk uint32
	zxvk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		fieldId, bts, err = msgp.ReadInt64Bytes(bts)
		//field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		//switch msgp.UnsafeString(field) {
		switch fieldId {
		case 0: //"Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case 1: // "BirthDay":
			z.BirthDay, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case 2: //"Phone":
			z.Phone, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case 3: //"Siblings":
			z.Siblings, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case 4: //"Spouse":
			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case 5: //"Money":
			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)
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
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) ZebraPackMsgsize() (s int) {
	//	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 9 + msgp.IntSize + 7 + msgp.BoolSize + 6 + msgp.Float64Size
	s = 1 + 1 + msgp.StringPrefixSize + len(z.Name) + 1 + msgp.TimeSize + 1 + msgp.StringPrefixSize + len(z.Phone) + 1 + msgp.IntSize + 1 + msgp.BoolSize + 1 + msgp.Float64Size
	return
}
