package test

import (
	"testing"

	"github.com/kindalus/gofunc/pkg/collections"
	"github.com/stretchr/testify/assert"
)

func TestCollectionsMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	m := collections.Map(s, func(r int) int {
		return r * 2
	})

	assert.Equal(t, 5, len(m))
	assert.Equal(t, 10, m[4])
	assert.Equal(t, 6, m[2])
}

func TestCollectionsMapInCollection(t *testing.T) {
	s := collections.From(10, 20, 30, 40, 50)

	f := func(r int) int {
		return r * 2
	}

	m := collections.Map(s, f)

	assert.Equal(t, 5, len(m))
	assert.Equal(t, 100, m[4])
	assert.Equal(t, 60, m[2])
}
