package comparisons

import (
	"killjvm/instructions/base"
	"killjvm/rtda"
)

type IFEQ struct{ base.BranchInstruction } //ifeq:x == 0
type IFNE struct{ base.BranchInstruction } //ifeq:x != 0
type IFLT struct{ base.BranchInstruction } //ifeq:x < 0
type IFLE struct{ base.BranchInstruction } //ifeq:x <= 0
type IFGT struct{ base.BranchInstruction } //ifeq:x > 0
type IFGE struct{ base.BranchInstruction } //ifeq:x >= 0
func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
