package eval

import (
	"fmt"
	"math"
)

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
	}
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
	}
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Sprintf("unsupported function call: %q", c.fn))
	}
}

func (m min) Eval(env Env) float64 {
	min := m.args[0].Eval(env)
	for _, arg := range m.args[1:] {
		v := arg.Eval(env)
		if min > v {
			min = v
		}
	}
	return min
}
