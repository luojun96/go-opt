package bufferpool

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBufferPool(t *testing.T) {
	c := qt.New(t)
	defer c.Done()

	buff := GetBuffer()
	buff.WriteString("do be do be do")
	c.Assert(buff.String(), qt.Equals, "do be do be do")
	PutBuffer(buff)

	c.Assert(buff.Len(), qt.Equals, 0)
}
