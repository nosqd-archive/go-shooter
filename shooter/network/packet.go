package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Packet struct {
	buffer        bytes.Buffer
	ContentLength int32
	PacketID      int32
}

func PacketCreate(packetId int32) *Packet {
	p := Packet{}

	p.buffer = bytes.Buffer{}
	p.PacketID = packetId

	return &p
}

func (self *Packet) Write(writer io.Writer) {
	var b bytes.Buffer

	err1 := binary.Write(&b, binary.LittleEndian, int32(self.PacketID))
	err2 := binary.Write(&b, binary.LittleEndian, int32(self.buffer.Len()))
	err3 := binary.Write(&b, binary.LittleEndian, self.buffer.Bytes())
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("error at writing")
	}
	_, err := writer.Write(b.Bytes())
	if err != nil {
		return
	}

}

func ReadPacket(reader io.Reader) *Packet {
	var contentLength int32
	var packetId int32
	err1 := binary.Read(reader, binary.LittleEndian, &packetId)
	err2 := binary.Read(reader, binary.LittleEndian, &contentLength)
	data := make([]byte, contentLength)
	err3 := binary.Read(reader, binary.LittleEndian, &data)
	if err1 != nil || err2 != nil || err3 != nil {
		return nil
	}
	p := Packet{ContentLength: contentLength, PacketID: packetId}
	p.buffer = bytes.Buffer{}
	p.buffer.Write(data)

	return &p
}

func (self *Packet) ReadInt8() int8 {
	var value int8

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteInt8(value int8) int8 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadInt16() int16 {
	var value int16

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteInt16(value int16) int16 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadInt32() int32 {
	var value int32

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteInt32(value int32) int32 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadInt64() int64 {
	var value int64

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteInt64(value int64) int64 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadUInt8() uint8 {
	var value uint8

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteUInt8(value uint8) uint8 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadUInt16() uint16 {
	var value uint16

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteUInt16(value uint16) uint16 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadUInt32() uint32 {
	var value uint32

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteUInt32(value uint32) uint32 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadUInt64() uint64 {
	var value uint64

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteUInt64(value uint64) uint64 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadFloat32() float32 {
	var value float32

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteFloat32(value float32) float32 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadFloat64() float64 {
	var value float64

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteFloat64(value float64) float64 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadComplex64() complex64 {
	var value complex64

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteComplex64(value complex64) complex64 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadComplex128() complex128 {
	var value complex128

	err := binary.Read(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}
func (self *Packet) WriteComplex128(value complex128) complex128 {

	err := binary.Write(&self.buffer, binary.LittleEndian, &value)
	if err != nil {
		return 0
	}

	return value
}

func (self *Packet) ReadString() string {
	var length int32 = self.ReadInt32()
	if length <= 0 {
		return ""
	}

	data := make([]byte, length)
	err := binary.Read(&self.buffer, binary.LittleEndian, &data)
	if err != nil {
		return ""
	}

	return string(data)
}
func (self *Packet) WriteString(value string) {
	length := int32(len(value))
	self.WriteInt32(length)

	if length > 0 {
		err := binary.Write(&self.buffer, binary.LittleEndian, []byte(value))
		if err != nil {
			fmt.Println("error at writing string")
		}
	}
}

func (self *Packet) ReadBool() bool {
	return self.ReadUInt8() == 255
}
func (self *Packet) WriteBool(value bool) bool {
	if value {
		self.WriteUInt8(255)
	} else {
		self.WriteUInt8(0)
	}
	return value
}
