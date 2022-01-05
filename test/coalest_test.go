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

func TestEmptyWitStrings(t *testing.T) {
	testCases := []EmptyTestCase[string]{
		{[]string{"1", "2"}, "1"},
		{[]string{"", "", "7", "2"}, "7"},
	}

	for _, test := range testCases {
		result := coalesce.Empty(test.Args...)

		assert.Equal(t, test.Result, result)
	}
}
