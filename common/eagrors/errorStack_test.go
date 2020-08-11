package eagrors

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

type dummyELogger struct {
	msg []string
}

func (d *dummyELogger) Error(args ...interface{}) {
	e, _ := args[0].(error)
	d.msg = append(d.msg, e.Error())
}

func TestStackError(t *testing.T) {

	e := NewError("message")
	e = NewErrorWithCause(e, "is")
	e = NewErrorWithCause(e, "This")

	t.Run("Log", func(t *testing.T) {
		c := qt.New(t)
		l := &dummyELogger{msg: make([]string, 0)}
		FprintStackError(l, e)
		c.Assert(l.msg, qt.CmpEquals(), []string{"This", "is", "message"})
	})

	t.Run("Message", func(t *testing.T) {
		c := qt.New(t)
		msg := GetErrorString(e)
		c.Assert(msg, qt.Equals, "This\nis\nmessage\n")
	})
}
