package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

//todo 补充完整MUTF-8解码
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
