package eval

import (
	"fmt"
	"strings"
)

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u Unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("incorrect Unary operator: %q", u.op)
	}
	return u.x.Check(vars)
}

func (b Binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("incorrect Binary operator: %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c Call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown func %q", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("Call %s has %d instead %d agruments",
			c.fn, len(c.args), arity)
	}

	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}

	return nil
}
