package rtda

import (
	"jvm-go/ch04/rtda"
)

type Frame struct {
	lower        *Frame
	localVars    rtda.LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVers(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
