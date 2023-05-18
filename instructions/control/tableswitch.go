package control

import "killjvm/instructions/base"

type TABLE_SWITCH struct {
	defaultOffset    int32
	low                int32
	high               int32
	jumpOffsets      []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 ! = 0 {
		self.ReadUint8()
	}
}