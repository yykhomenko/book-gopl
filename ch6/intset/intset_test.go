package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, "{2 4}", x.String())
}

func TestHas(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.True(t, x.Has(4))
	assert.True(t, !x.Has(3))
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	assert.Equal(t, "{1 9 42 144}", x.String())
}

func TestLen(t *testing.T) {
	x := IntSet{}
	x.Add(4)
	x.Add(2)
	assert.Equal(t, 2, x.Len())
}