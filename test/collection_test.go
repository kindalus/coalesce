package test

import (
	"testing"

	"github.com/kindalus/gofunc/pkg/collections"
	"github.com/stretchr/testify/assert"
)

func TestCreateCollectionFromLiterals(t *testing.T) {
	c := collections.From(1, 2, 3, 4, 5)

	assert.Equal(t, 5, len(c))
	assert.Equal(t, 4, c[3])
	assert.Equal(t, 3, c[2])
}

func TestCreateEmptyCollection(t *testing.T) {
	c := collections.From[int]()
	e := collections.Empty[string]()

	assert.Equal(t, 0, len(c))
	assert.Equal(t, 0, len(e))
}

func TestCreateCollectionFromSlice(t *testing.T) {
	s := []int{'a', 'b', 'c', 'd', 'e'}
	c := collections.From(s...)

	assert.Equal(t, 5, len(c))
	assert.NotEqual(t, &s, &c)
}

func TestCollectionFilter(t *testing.T) {

	c := collections.From(1, 2, 3, 4, 5)

	c = c.Filter(func(r int) bool {
		return r%2 == 0
	})

	assert.Equal(t, 2, len(c))
	assert.Equal(t, 2, c[0])
	assert.Equal(t, 4, c[1])
}

func TestCollectionAny(t *testing.T) {

	c := collections.From(1, 2, 3, 4, 5)

	assert.True(t, c.Any(func(r int) bool {
		return r%2 == 0
	}))

	assert.False(t, c.Any(func(r int) bool {
		return r == 0
	}))
}

func TestCollectionSort(t *testing.T) {

	c := collections.From(1, 2, 3, 4, 5)

	c = c.Sort(func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, 5, len(c))
	assert.Equal(t, 5, c[0])
	assert.Equal(t, 4, c[1])
}
