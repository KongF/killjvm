package constants

import "killjvm/instructions/base"
import "killjvm/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
