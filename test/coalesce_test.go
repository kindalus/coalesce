package test

import (
	"testing"

	"github.com/kindalus/coalesce/pkg/coalesce"
	"github.com/stretchr/testify/assert"
)

type EmptyTestCase[T any] struct {
	Args   []T
	Result T
}

type EqualTestCase[T any] struct {
	Target T
	Args   []T
	Result T
}

func TestEmptyWithStrings(t *testing.T) {
	testCases := []EmptyTestCase[string]{
		{[]string{"1", "2"}, "1"},
		{[]string{"", "", "7", "2"}, "7"},
	}

	for _, test := range testCases {
		result := coalesce.Empty(test.Args...)

		assert.Equal(t, test.Result, result)
	}
}

func TestEqual(t *testing.T) {

	type Car struct {
		Model string
		Year  int
	}

	testCases := []EqualTestCase[interface{}]{
		{
			"1",
			[]interface{}{"1", "2"},
			"1",
		},
		{
			Car{"LR4", 2016},
			[]interface{}{
				Car{"Rav4", 2019},
				Car{"Jetour", 2020},
				Car{"LR4", 2016},
			},
			Car{"LR4", 2016},
		},
	}

	for _, test := range testCases {
		result := coalesce.Equal(test.Target, test.Args...)

		assert.Equal(t, test.Result, result)
	}
}
