package radix

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Edge) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Label":
			z.Label, err = dc.ReadByte()
			if err != nil {
				return
			}
		case "Node":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Node = nil
			} else {
				if z.Node == nil {
					z.Node = new(Node)
				}
				err = z.Node.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Edge) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Label"
	err = en.Append(0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteByte(z.Label)
	if err != nil {
		return
	}
	// write "Node"
	err = en.Append(0xa4, 0x4e, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	if z.Node == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Node.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Edge) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Label"
	o = append(o, 0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
	o = msgp.AppendByte(o, z.Label)
	// string "Node"
	o = append(o, 0xa4, 0x4e, 0x6f, 0x64, 0x65)
	if z.Node == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Node.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Edge) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Label":
			z.Label, bts, err = msgp.ReadByteBytes(bts)
			if err != nil {
				return
			}
		case "Node":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Node = nil
			} else {
				if z.Node == nil {
					z.Node = new(Node)
				}
				bts, err = z.Node.UnmarshalMsg(bts)
				if err != nil {
					return
				}
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
func (z *Edge) Msgsize() (s int) {
	s = 1 + 6 + msgp.ByteSize + 5
	if z.Node == nil {
		s += msgp.NilSize
	} else {
		s += z.Node.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Edges) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Edges, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, err = dc.ReadMapHeader()
		if err != nil {
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			switch msgp.UnsafeString(field) {
			case "Label":
				(*z)[zb0001].Label, err = dc.ReadByte()
				if err != nil {
					return
				}
			case "Node":
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					(*z)[zb0001].Node = nil
				} else {
					if (*z)[zb0001].Node == nil {
						(*z)[zb0001].Node = new(Node)
					}
					err = (*z)[zb0001].Node.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			default:
				err = dc.Skip()
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Edges) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0004 := range z {
		// map header, size 2
		// write "Label"
		err = en.Append(0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
		if err != nil {
			return err
		}
		err = en.WriteByte(z[zb0004].Label)
		if err != nil {
			return
		}
		// write "Node"
		err = en.Append(0xa4, 0x4e, 0x6f, 0x64, 0x65)
		if err != nil {
			return err
		}
		if z[zb0004].Node == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0004].Node.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Edges) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		// map header, size 2
		// string "Label"
		o = append(o, 0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
		o = msgp.AppendByte(o, z[zb0004].Label)
		// string "Node"
		o = append(o, 0xa4, 0x4e, 0x6f, 0x64, 0x65)
		if z[zb0004].Node == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0004].Node.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Edges) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Edges, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			switch msgp.UnsafeString(field) {
			case "Label":
				(*z)[zb0001].Label, bts, err = msgp.ReadByteBytes(bts)
				if err != nil {
					return
				}
			case "Node":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].Node = nil
				} else {
					if (*z)[zb0001].Node == nil {
						(*z)[zb0001].Node = new(Node)
					}
					bts, err = (*z)[zb0001].Node.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Edges) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + 6 + msgp.ByteSize + 5
		if z[zb0004].Node == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].Node.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LeafNode) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Val":
			z.Val, err = dc.ReadIntf()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LeafNode) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Key"
	err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "Val"
	err = en.Append(0xa3, 0x56, 0x61, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteIntf(z.Val)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LeafNode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Key"
	o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
	o = msgp.AppendString(o, z.Key)
	// string "Val"
	o = append(o, 0xa3, 0x56, 0x61, 0x6c)
	o, err = msgp.AppendIntf(o, z.Val)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LeafNode) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Key":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Val":
			z.Val, bts, err = msgp.ReadIntfBytes(bts)
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
func (z LeafNode) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 4 + msgp.GuessSize(z.Val)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Node) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Leaf":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Leaf = nil
			} else {
				if z.Leaf == nil {
					z.Leaf = new(LeafNode)
				}
				var zb0002 uint32
				zb0002, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Leaf.Key, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Val":
						z.Leaf.Val, err = dc.ReadIntf()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "Prefix":
			z.Prefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Edges":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Edges) >= int(zb0003) {
				z.Edges = (z.Edges)[:zb0003]
			} else {
				z.Edges = make(Edges, zb0003)
			}
			for za0001 := range z.Edges {
				var zb0004 uint32
				zb0004, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0004 > 0 {
					zb0004--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Label":
						z.Edges[za0001].Label, err = dc.ReadByte()
						if err != nil {
							return
						}
					case "Node":
						if dc.IsNil() {
							err = dc.ReadNil()
							if err != nil {
								return
							}
							z.Edges[za0001].Node = nil
						} else {
							if z.Edges[za0001].Node == nil {
								z.Edges[za0001].Node = new(Node)
							}
							err = z.Edges[za0001].Node.DecodeMsg(dc)
							if err != nil {
								return
							}
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Node) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Leaf"
	err = en.Append(0x83, 0xa4, 0x4c, 0x65, 0x61, 0x66)
	if err != nil {
		return err
	}
	if z.Leaf == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 2
		// write "Key"
		err = en.Append(0x82, 0xa3, 0x4b, 0x65, 0x79)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Leaf.Key)
		if err != nil {
			return
		}
		// write "Val"
		err = en.Append(0xa3, 0x56, 0x61, 0x6c)
		if err != nil {
			return err
		}
		err = en.WriteIntf(z.Leaf.Val)
		if err != nil {
			return
		}
	}
	// write "Prefix"
	err = en.Append(0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Prefix)
	if err != nil {
		return
	}
	// write "Edges"
	err = en.Append(0xa5, 0x45, 0x64, 0x67, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Edges)))
	if err != nil {
		return
	}
	for za0001 := range z.Edges {
		// map header, size 2
		// write "Label"
		err = en.Append(0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
		if err != nil {
			return err
		}
		err = en.WriteByte(z.Edges[za0001].Label)
		if err != nil {
			return
		}
		// write "Node"
		err = en.Append(0xa4, 0x4e, 0x6f, 0x64, 0x65)
		if err != nil {
			return err
		}
		if z.Edges[za0001].Node == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Edges[za0001].Node.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Node) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Leaf"
	o = append(o, 0x83, 0xa4, 0x4c, 0x65, 0x61, 0x66)
	if z.Leaf == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 2
		// string "Key"
		o = append(o, 0x82, 0xa3, 0x4b, 0x65, 0x79)
		o = msgp.AppendString(o, z.Leaf.Key)
		// string "Val"
		o = append(o, 0xa3, 0x56, 0x61, 0x6c)
		o, err = msgp.AppendIntf(o, z.Leaf.Val)
		if err != nil {
			return
		}
	}
	// string "Prefix"
	o = append(o, 0xa6, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78)
	o = msgp.AppendString(o, z.Prefix)
	// string "Edges"
	o = append(o, 0xa5, 0x45, 0x64, 0x67, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Edges)))
	for za0001 := range z.Edges {
		// map header, size 2
		// string "Label"
		o = append(o, 0x82, 0xa5, 0x4c, 0x61, 0x62, 0x65, 0x6c)
		o = msgp.AppendByte(o, z.Edges[za0001].Label)
		// string "Node"
		o = append(o, 0xa4, 0x4e, 0x6f, 0x64, 0x65)
		if z.Edges[za0001].Node == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Edges[za0001].Node.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Node) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Leaf":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Leaf = nil
			} else {
				if z.Leaf == nil {
					z.Leaf = new(LeafNode)
				}
				var zb0002 uint32
				zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0002 > 0 {
					zb0002--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Key":
						z.Leaf.Key, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Val":
						z.Leaf.Val, bts, err = msgp.ReadIntfBytes(bts)
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
			}
		case "Prefix":
			z.Prefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Edges":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Edges) >= int(zb0003) {
				z.Edges = (z.Edges)[:zb0003]
			} else {
				z.Edges = make(Edges, zb0003)
			}
			for za0001 := range z.Edges {
				var zb0004 uint32
				zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0004 > 0 {
					zb0004--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Label":
						z.Edges[za0001].Label, bts, err = msgp.ReadByteBytes(bts)
						if err != nil {
							return
						}
					case "Node":
						if msgp.IsNil(bts) {
							bts, err = msgp.ReadNilBytes(bts)
							if err != nil {
								return
							}
							z.Edges[za0001].Node = nil
						} else {
							if z.Edges[za0001].Node == nil {
								z.Edges[za0001].Node = new(Node)
							}
							bts, err = z.Edges[za0001].Node.UnmarshalMsg(bts)
							if err != nil {
								return
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
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
func (z *Node) Msgsize() (s int) {
	s = 1 + 5
	if z.Leaf == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 4 + msgp.StringPrefixSize + len(z.Leaf.Key) + 4 + msgp.GuessSize(z.Leaf.Val)
	}
	s += 7 + msgp.StringPrefixSize + len(z.Prefix) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Edges {
		s += 1 + 6 + msgp.ByteSize + 5
		if z.Edges[za0001].Node == nil {
			s += msgp.NilSize
		} else {
			s += z.Edges[za0001].Node.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Tree) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Root":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Root = nil
			} else {
				if z.Root == nil {
					z.Root = new(Node)
				}
				err = z.Root.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Size":
			z.Size, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Tree) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Root"
	err = en.Append(0x82, 0xa4, 0x52, 0x6f, 0x6f, 0x74)
	if err != nil {
		return err
	}
	if z.Root == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Root.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Size"
	err = en.Append(0xa4, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Size)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tree) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Root"
	o = append(o, 0x82, 0xa4, 0x52, 0x6f, 0x6f, 0x74)
	if z.Root == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Root.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt(o, z.Size)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tree) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Root":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Root = nil
			} else {
				if z.Root == nil {
					z.Root = new(Node)
				}
				bts, err = z.Root.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Size":
			z.Size, bts, err = msgp.ReadIntBytes(bts)
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
func (z *Tree) Msgsize() (s int) {
	s = 1 + 5
	if z.Root == nil {
		s += msgp.NilSize
	} else {
		s += z.Root.Msgsize()
	}
	s += 5 + msgp.IntSize
	return
}
