package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *PoolDecommissionInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "st":
			z.StartTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "ss":
			z.StartSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "StartSize")
				return
			}
		case "ts":
			z.TotalSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "TotalSize")
				return
			}
		case "cs":
			z.CurrentSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "CurrentSize")
				return
			}
		case "cmp":
			z.Complete, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fl":
			z.Failed, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		case "cnl":
			z.Canceled, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Canceled")
				return
			}
		case "bkts":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "QueuedBuckets")
				return
			}
			if cap(z.QueuedBuckets) >= int(zb0002) {
				z.QueuedBuckets = (z.QueuedBuckets)[:zb0002]
			} else {
				z.QueuedBuckets = make([]string, zb0002)
			}
			for za0001 := range z.QueuedBuckets {
				z.QueuedBuckets[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "QueuedBuckets", za0001)
					return
				}
			}
		case "dbkts":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "DecommissionedBuckets")
				return
			}
			if cap(z.DecommissionedBuckets) >= int(zb0003) {
				z.DecommissionedBuckets = (z.DecommissionedBuckets)[:zb0003]
			} else {
				z.DecommissionedBuckets = make([]string, zb0003)
			}
			for za0002 := range z.DecommissionedBuckets {
				z.DecommissionedBuckets[za0002], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "DecommissionedBuckets", za0002)
					return
				}
			}
		case "bkt":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "obj":
			z.Object, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "id":
			z.ItemsDecommissioned, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ItemsDecommissioned")
				return
			}
		case "idf":
			z.ItemsDecommissionFailed, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ItemsDecommissionFailed")
				return
			}
		case "bd":
			z.BytesDone, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BytesDone")
				return
			}
		case "bf":
			z.BytesFailed, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "BytesFailed")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolDecommissionInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "st"
	err = en.Append(0x8f, 0xa2, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteTime(z.StartTime)
	if err != nil {
		err = msgp.WrapError(err, "StartTime")
		return
	}
	// write "ss"
	err = en.Append(0xa2, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.StartSize)
	if err != nil {
		err = msgp.WrapError(err, "StartSize")
		return
	}
	// write "ts"
	err = en.Append(0xa2, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.TotalSize)
	if err != nil {
		err = msgp.WrapError(err, "TotalSize")
		return
	}
	// write "cs"
	err = en.Append(0xa2, 0x63, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.CurrentSize)
	if err != nil {
		err = msgp.WrapError(err, "CurrentSize")
		return
	}
	// write "cmp"
	err = en.Append(0xa3, 0x63, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Complete)
	if err != nil {
		err = msgp.WrapError(err, "Complete")
		return
	}
	// write "fl"
	err = en.Append(0xa2, 0x66, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Failed)
	if err != nil {
		err = msgp.WrapError(err, "Failed")
		return
	}
	// write "cnl"
	err = en.Append(0xa3, 0x63, 0x6e, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Canceled)
	if err != nil {
		err = msgp.WrapError(err, "Canceled")
		return
	}
	// write "bkts"
	err = en.Append(0xa4, 0x62, 0x6b, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.QueuedBuckets)))
	if err != nil {
		err = msgp.WrapError(err, "QueuedBuckets")
		return
	}
	for za0001 := range z.QueuedBuckets {
		err = en.WriteString(z.QueuedBuckets[za0001])
		if err != nil {
			err = msgp.WrapError(err, "QueuedBuckets", za0001)
			return
		}
	}
	// write "dbkts"
	err = en.Append(0xa5, 0x64, 0x62, 0x6b, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.DecommissionedBuckets)))
	if err != nil {
		err = msgp.WrapError(err, "DecommissionedBuckets")
		return
	}
	for za0002 := range z.DecommissionedBuckets {
		err = en.WriteString(z.DecommissionedBuckets[za0002])
		if err != nil {
			err = msgp.WrapError(err, "DecommissionedBuckets", za0002)
			return
		}
	}
	// write "bkt"
	err = en.Append(0xa3, 0x62, 0x6b, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "obj"
	err = en.Append(0xa3, 0x6f, 0x62, 0x6a)
	if err != nil {
		return
	}
	err = en.WriteString(z.Object)
	if err != nil {
		err = msgp.WrapError(err, "Object")
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ItemsDecommissioned)
	if err != nil {
		err = msgp.WrapError(err, "ItemsDecommissioned")
		return
	}
	// write "idf"
	err = en.Append(0xa3, 0x69, 0x64, 0x66)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ItemsDecommissionFailed)
	if err != nil {
		err = msgp.WrapError(err, "ItemsDecommissionFailed")
		return
	}
	// write "bd"
	err = en.Append(0xa2, 0x62, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BytesDone)
	if err != nil {
		err = msgp.WrapError(err, "BytesDone")
		return
	}
	// write "bf"
	err = en.Append(0xa2, 0x62, 0x66)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.BytesFailed)
	if err != nil {
		err = msgp.WrapError(err, "BytesFailed")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolDecommissionInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "st"
	o = append(o, 0x8f, 0xa2, 0x73, 0x74)
	o = msgp.AppendTime(o, z.StartTime)
	// string "ss"
	o = append(o, 0xa2, 0x73, 0x73)
	o = msgp.AppendInt64(o, z.StartSize)
	// string "ts"
	o = append(o, 0xa2, 0x74, 0x73)
	o = msgp.AppendInt64(o, z.TotalSize)
	// string "cs"
	o = append(o, 0xa2, 0x63, 0x73)
	o = msgp.AppendInt64(o, z.CurrentSize)
	// string "cmp"
	o = append(o, 0xa3, 0x63, 0x6d, 0x70)
	o = msgp.AppendBool(o, z.Complete)
	// string "fl"
	o = append(o, 0xa2, 0x66, 0x6c)
	o = msgp.AppendBool(o, z.Failed)
	// string "cnl"
	o = append(o, 0xa3, 0x63, 0x6e, 0x6c)
	o = msgp.AppendBool(o, z.Canceled)
	// string "bkts"
	o = append(o, 0xa4, 0x62, 0x6b, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.QueuedBuckets)))
	for za0001 := range z.QueuedBuckets {
		o = msgp.AppendString(o, z.QueuedBuckets[za0001])
	}
	// string "dbkts"
	o = append(o, 0xa5, 0x64, 0x62, 0x6b, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.DecommissionedBuckets)))
	for za0002 := range z.DecommissionedBuckets {
		o = msgp.AppendString(o, z.DecommissionedBuckets[za0002])
	}
	// string "bkt"
	o = append(o, 0xa3, 0x62, 0x6b, 0x74)
	o = msgp.AppendString(o, z.Bucket)
	// string "obj"
	o = append(o, 0xa3, 0x6f, 0x62, 0x6a)
	o = msgp.AppendString(o, z.Object)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.ItemsDecommissioned)
	// string "idf"
	o = append(o, 0xa3, 0x69, 0x64, 0x66)
	o = msgp.AppendInt64(o, z.ItemsDecommissionFailed)
	// string "bd"
	o = append(o, 0xa2, 0x62, 0x64)
	o = msgp.AppendInt64(o, z.BytesDone)
	// string "bf"
	o = append(o, 0xa2, 0x62, 0x66)
	o = msgp.AppendInt64(o, z.BytesFailed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolDecommissionInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "st":
			z.StartTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "ss":
			z.StartSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartSize")
				return
			}
		case "ts":
			z.TotalSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TotalSize")
				return
			}
		case "cs":
			z.CurrentSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentSize")
				return
			}
		case "cmp":
			z.Complete, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Complete")
				return
			}
		case "fl":
			z.Failed, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Failed")
				return
			}
		case "cnl":
			z.Canceled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Canceled")
				return
			}
		case "bkts":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "QueuedBuckets")
				return
			}
			if cap(z.QueuedBuckets) >= int(zb0002) {
				z.QueuedBuckets = (z.QueuedBuckets)[:zb0002]
			} else {
				z.QueuedBuckets = make([]string, zb0002)
			}
			for za0001 := range z.QueuedBuckets {
				z.QueuedBuckets[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "QueuedBuckets", za0001)
					return
				}
			}
		case "dbkts":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DecommissionedBuckets")
				return
			}
			if cap(z.DecommissionedBuckets) >= int(zb0003) {
				z.DecommissionedBuckets = (z.DecommissionedBuckets)[:zb0003]
			} else {
				z.DecommissionedBuckets = make([]string, zb0003)
			}
			for za0002 := range z.DecommissionedBuckets {
				z.DecommissionedBuckets[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "DecommissionedBuckets", za0002)
					return
				}
			}
		case "bkt":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "obj":
			z.Object, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "id":
			z.ItemsDecommissioned, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ItemsDecommissioned")
				return
			}
		case "idf":
			z.ItemsDecommissionFailed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ItemsDecommissionFailed")
				return
			}
		case "bd":
			z.BytesDone, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BytesDone")
				return
			}
		case "bf":
			z.BytesFailed, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BytesFailed")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolDecommissionInfo) Msgsize() (s int) {
	s = 1 + 3 + msgp.TimeSize + 3 + msgp.Int64Size + 3 + msgp.Int64Size + 3 + msgp.Int64Size + 4 + msgp.BoolSize + 3 + msgp.BoolSize + 4 + msgp.BoolSize + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.QueuedBuckets {
		s += msgp.StringPrefixSize + len(z.QueuedBuckets[za0001])
	}
	s += 6 + msgp.ArrayHeaderSize
	for za0002 := range z.DecommissionedBuckets {
		s += msgp.StringPrefixSize + len(z.DecommissionedBuckets[za0002])
	}
	s += 4 + msgp.StringPrefixSize + len(z.Bucket) + 4 + msgp.StringPrefixSize + len(z.Object) + 3 + msgp.Int64Size + 4 + msgp.Int64Size + 3 + msgp.Int64Size + 3 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PoolStatus) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "cl":
			z.CmdLine, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "CmdLine")
				return
			}
		case "lu":
			z.LastUpdate, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "dec":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Decommission")
					return
				}
				z.Decommission = nil
			} else {
				if z.Decommission == nil {
					z.Decommission = new(PoolDecommissionInfo)
				}
				err = z.Decommission.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Decommission")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PoolStatus) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "id"
	err = en.Append(0x84, 0xa2, 0x69, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt(z.ID)
	if err != nil {
		err = msgp.WrapError(err, "ID")
		return
	}
	// write "cl"
	err = en.Append(0xa2, 0x63, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteString(z.CmdLine)
	if err != nil {
		err = msgp.WrapError(err, "CmdLine")
		return
	}
	// write "lu"
	err = en.Append(0xa2, 0x6c, 0x75)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastUpdate")
		return
	}
	// write "dec"
	err = en.Append(0xa3, 0x64, 0x65, 0x63)
	if err != nil {
		return
	}
	if z.Decommission == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Decommission.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Decommission")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PoolStatus) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "id"
	o = append(o, 0x84, 0xa2, 0x69, 0x64)
	o = msgp.AppendInt(o, z.ID)
	// string "cl"
	o = append(o, 0xa2, 0x63, 0x6c)
	o = msgp.AppendString(o, z.CmdLine)
	// string "lu"
	o = append(o, 0xa2, 0x6c, 0x75)
	o = msgp.AppendTime(o, z.LastUpdate)
	// string "dec"
	o = append(o, 0xa3, 0x64, 0x65, 0x63)
	if z.Decommission == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Decommission.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Decommission")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PoolStatus) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ID")
				return
			}
		case "cl":
			z.CmdLine, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CmdLine")
				return
			}
		case "lu":
			z.LastUpdate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "dec":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Decommission = nil
			} else {
				if z.Decommission == nil {
					z.Decommission = new(PoolDecommissionInfo)
				}
				bts, err = z.Decommission.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Decommission")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PoolStatus) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 3 + msgp.StringPrefixSize + len(z.CmdLine) + 3 + msgp.TimeSize + 4
	if z.Decommission == nil {
		s += msgp.NilSize
	} else {
		s += z.Decommission.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *decomError) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Err":
			z.Err, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z decomError) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Err"
	err = en.Append(0x81, 0xa3, 0x45, 0x72, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Err)
	if err != nil {
		err = msgp.WrapError(err, "Err")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z decomError) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Err"
	o = append(o, 0x81, 0xa3, 0x45, 0x72, 0x72)
	o = msgp.AppendString(o, z.Err)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *decomError) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Err":
			z.Err, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z decomError) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Err)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *poolMeta) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, err = dc.ReadInt()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "pls":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Pools")
				return
			}
			if cap(z.Pools) >= int(zb0002) {
				z.Pools = (z.Pools)[:zb0002]
			} else {
				z.Pools = make([]PoolStatus, zb0002)
			}
			for za0001 := range z.Pools {
				err = z.Pools[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Pools", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *poolMeta) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "v"
	err = en.Append(0x82, 0xa1, 0x76)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "pls"
	err = en.Append(0xa3, 0x70, 0x6c, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Pools)))
	if err != nil {
		err = msgp.WrapError(err, "Pools")
		return
	}
	for za0001 := range z.Pools {
		err = z.Pools[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Pools", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *poolMeta) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "v"
	o = append(o, 0x82, 0xa1, 0x76)
	o = msgp.AppendInt(o, z.Version)
	// string "pls"
	o = append(o, 0xa3, 0x70, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Pools)))
	for za0001 := range z.Pools {
		o, err = z.Pools[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Pools", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *poolMeta) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "v":
			z.Version, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "pls":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Pools")
				return
			}
			if cap(z.Pools) >= int(zb0002) {
				z.Pools = (z.Pools)[:zb0002]
			} else {
				z.Pools = make([]PoolStatus, zb0002)
			}
			for za0001 := range z.Pools {
				bts, err = z.Pools[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Pools", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *poolMeta) Msgsize() (s int) {
	s = 1 + 2 + msgp.IntSize + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Pools {
		s += z.Pools[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *poolSpaceInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Free":
			z.Free, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Free")
				return
			}
		case "Total":
			z.Total, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Total")
				return
			}
		case "Used":
			z.Used, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Used")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z poolSpaceInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Free"
	err = en.Append(0x83, 0xa4, 0x46, 0x72, 0x65, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Free)
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	// write "Total"
	err = en.Append(0xa5, 0x54, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Total)
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	// write "Used"
	err = en.Append(0xa4, 0x55, 0x73, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Used)
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z poolSpaceInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Free"
	o = append(o, 0x83, 0xa4, 0x46, 0x72, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.Free)
	// string "Total"
	o = append(o, 0xa5, 0x54, 0x6f, 0x74, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.Total)
	// string "Used"
	o = append(o, 0xa4, 0x55, 0x73, 0x65, 0x64)
	o = msgp.AppendInt64(o, z.Used)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *poolSpaceInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Free":
			z.Free, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Free")
				return
			}
		case "Total":
			z.Total, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Total")
				return
			}
		case "Used":
			z.Used, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Used")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z poolSpaceInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int64Size + 6 + msgp.Int64Size + 5 + msgp.Int64Size
	return
}
