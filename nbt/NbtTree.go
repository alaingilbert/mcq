package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
)

// NbtTree represents an NbtTree.
type NbtTree struct {
	Stream   io.Reader
	root     *TagNodeCompound
	rootName string
}

// NewNbtTree creates a new NbtTree.
func NewNbtTree(r io.Reader) *NbtTree {
	tree := new(NbtTree)
	tree.Stream = r
	tree.root = tree.ReadRoot()
	return tree
}

// Root get the root element of the NbtTree.
func (n *NbtTree) Root() *TagNodeCompound {
	return n.root
}

// ReadRoot ...
func (n *NbtTree) ReadRoot() *TagNodeCompound {
	tagType := TagType(ReadByte(n.Stream))
	if tagType == TagCompound {
		n.rootName = ReadString(n.Stream)
		return n.ReadValue(tagType).(*TagNodeCompound)
	}
	return new(TagNodeCompound)
}

// ReadValue ...
func (n *NbtTree) ReadValue(pTagType TagType) ITagNode {
	var tagNode ITagNode

	switch pTagType {
	case TagByte:
		tagNode = n.ReadByte()
	case TagCompound:
		tagNode = n.ReadCompound()
	case TagList:
		tagNode = n.ReadList()
	case TagByteArray:
		tagNode = n.ReadByteArray()
	case TagLong:
		tagNode = n.ReadLong()
	case TagInt:
		tagNode = n.ReadInt()
	case TagIntArray:
		tagNode = n.ReadIntArray()
	case TagShort:
		tagNode = n.ReadShort()
	case TagFloat:
		tagNode = n.ReadFloat()
	case TagDouble:
		tagNode = n.ReadDouble()
	case TagString:
		tagNode = n.ReadString()
	case TagLongArray:
		tagNode = n.ReadLongArray()
	default:
		fmt.Println("Unknown TagNode", pTagType)
		tagNode = NewTagNodeUnknown()
	}

	return tagNode
}

// ReadFloat ...
func (n *NbtTree) ReadFloat() *TagNodeFloat {
	floatRead := ReadFloat(n.Stream)
	return NewTagNodeFloat(floatRead)
}

// ReadString ...
func (n *NbtTree) ReadString() *TagNodeString {
	stringRead := ReadString(n.Stream)
	return NewTagNodeString(stringRead)
}

// ReadDouble ...
func (n *NbtTree) ReadDouble() *TagNodeDouble {
	doubleRead := ReadDouble(n.Stream)
	return NewTagNodeDouble(doubleRead)
}

// ReadShort ...
func (n *NbtTree) ReadShort() *TagNodeShort {
	shortRead := ReadShort(n.Stream)
	return NewTagNodeShort(shortRead)
}

// ReadIntArray read an array of TAG_Int's payloads.
func (n *NbtTree) ReadIntArray() *TagNodeIntArray {
	size := ReadInt(n.Stream)
	intArray := make([]int32, size)
	for i := int32(0); i < size; i++ {
		tmpInt := ReadInt(n.Stream)
		intArray[i] = tmpInt
	}
	return NewTagNodeIntArray(intArray)
}

// ReadLongArray read an array of TAG_Long's payloads.
func (n *NbtTree) ReadLongArray() *TagNodeLongArray {
	size := ReadInt(n.Stream)
	longArray := make([]int64, size)
	for i := int64(0); i < int64(size); i++ {
		tmpInt := ReadLong(n.Stream)
		longArray[i] = tmpInt
	}
	return NewTagNodeLongArray(longArray)
}

// ReadByte read a signed integral type. Sometimes used for booleans.
func (n *NbtTree) ReadByte() *TagNodeByte {
	byteRead := ReadByte(n.Stream)
	return NewTagNodeByte(byteRead)
}

// ReadInt read a signed integral type.
func (n *NbtTree) ReadInt() *TagNodeInt {
	intRead := ReadInt(n.Stream)
	return NewTagNodeInt(intRead)
}

// ReadLong read a signed integral type.
func (n *NbtTree) ReadLong() *TagNodeLong {
	longRead := ReadLong(n.Stream)
	return NewTagNodeLong(longRead)
}

// ReadByteArray read an array of bytes.
func (n *NbtTree) ReadByteArray() *TagNodeByteArray {
	size := ReadInt(n.Stream)
	if size < 0 {
		log.Fatal("Read Neg")
	}
	byteArray := make([]byte, size)
	_, _ = n.Stream.Read(byteArray)

	return NewTagNodeByteArray(byteArray)
}

// ReadList read a list of tag payloads, without repeated tag IDs or any tag
// names.
func (n *NbtTree) ReadList() *TagNodeList {
	tagID := TagType(ReadByte(n.Stream))
	length := ReadInt(n.Stream)
	list := make([]ITagNode, length)
	val := NewTagNodeList(tagID, length, list)
	if val.ValueType() == TagEnd {
		return NewTagNodeList(TagByte, length, list)
	}
	for i := 0; int32(i) < length; i++ {
		val.Add(n.ReadValue(val.ValueType()), i)
	}
	return val
}

// ReadCompound read a list of fully formed tags, including their IDs, names,
// and payloads. No two tags may have the same name
func (n *NbtTree) ReadCompound() *TagNodeCompound {
	entriesMap := make(map[string]ITagNode)
	tagNodeCompound := NewTagNodeCompound(entriesMap)
	for n.ReadTag(tagNodeCompound) {
	}
	return tagNodeCompound
}

// ReadTag read a tag from the nbtTree stream.
// Return false if tagEnd was found, true otherwise.
func (n *NbtTree) ReadTag(pParent *TagNodeCompound) bool {
	tagType := TagType(ReadByte(n.Stream))
	if tagType != TagEnd {
		name := ReadString(n.Stream)
		value := n.ReadValue(tagType)
		pParent.Entries[name] = value
		return true
	}
	return false
}

// ReadByte ...
func ReadByte(r io.Reader) (i byte) {
	b := make([]byte, 1)
	_, _ = r.Read(b)
	i = b[0]
	return
}

// ReadShort ...
func ReadShort(r io.Reader) (i int16) {
	_ = binary.Read(r, binary.BigEndian, &i)
	return
}

// ReadInt ...
func ReadInt(r io.Reader) (i int32) {
	_ = binary.Read(r, binary.BigEndian, &i)
	return
}

// ReadLong ...
func ReadLong(r io.Reader) (i int64) {
	_ = binary.Read(r, binary.BigEndian, &i)
	return
}

// ReadFloat ...
func ReadFloat(r io.Reader) (i float32) {
	b := make([]byte, 4)
	_, _ = r.Read(b)
	i = math.Float32frombits(binary.BigEndian.Uint32(b))
	return
}

// ReadDouble ...
func ReadDouble(r io.Reader) (i float64) {
	b := make([]byte, 8)
	_, _ = r.Read(b)
	i = math.Float64frombits(binary.BigEndian.Uint64(b))
	return
}

// ReadByteArray ...
func ReadByteArray(r io.Reader) (i []byte) {
	i = make([]byte, ReadInt(r))
	_, _ = r.Read(i)
	return
}

// ReadString ...
func ReadString(r io.Reader) string {
	result := make([]byte, ReadShort(r))
	_, _ = r.Read(result)
	return string(result)
}

// ReadIntArray ...
func ReadIntArray(r io.Reader) (list []int32) {
	length := int(ReadInt(r))
	for i := 0; i < length; i++ {
		list = append(list, ReadInt(r))
	}
	return
}
