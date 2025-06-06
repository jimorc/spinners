package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpinnerBase_NewSpinnerBase(t *testing.T) {
	s := NewSpinnerBase(2, 4, 1)
	assert.Equal(t, 2., s.Min)
	assert.Equal(t, 4., s.Max)
	assert.Equal(t, 1., s.Step)
	assert.Equal(t, 2., s.Value)
}
