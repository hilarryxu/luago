package binchunk

import (
	"encoding/binary"

	"github.com/hilarryxu/golua/types"
)

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *reader) readUint32() uint32 {
	val := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *reader) readUint64() uint64 {
	val := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	if size == 0 {
		return ""
	}
	if size == 0xFF {
		// size_t
		// n + 1
		size = uint(self.readUint64())
	}

	bytes := self.readBytes(size - 1)
	return string(bytes)
}

func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != types.LUA_SIGNATURE {
		panic("not a precompiled chunk")
	}
}
