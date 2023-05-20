package rtda

import "killjvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
