package reserved

import "killjvm/instructions/base"
import "killjvm/rtda"
import "killjvm/native"
import _ "killjvm/native/java/io"
import _ "killjvm/native/java/lang"
import _ "killjvm/native/java/security"
import _ "killjvm/native/java/util/concurrent/atomic"
import _ "killjvm/native/sun/io"
import _ "killjvm/native/sun/misc"
import _ "killjvm/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
