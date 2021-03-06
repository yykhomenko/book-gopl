package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u Unary) String() string {
	return string(u.op) + u.x.String()
}

func (b Binary) String() string {
	switch b.op {
	case '+', '-':
		return fmt.Sprintf("%s %s %s", b.x.String(), string(b.op), b.y.String())
	case '*', '/':
		return fmt.Sprintf("%s%s%s", b.x.String(), string(b.op), b.y.String())
	default:
		panic(fmt.Sprintf("unsupported Binary operator: %q", b.op))
	}
}

func (c Call) String() string {
	var b bytes.Buffer
	b.WriteString(c.fn)
	b.WriteRune('(')
	for i, arg := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(arg.String())
	}
	b.WriteRune(')')
	return b.String()
}
