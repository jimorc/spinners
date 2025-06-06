package widget

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/stretchr/testify/assert"
)

type spinner struct {
	SpinnerBase
}

func (s *spinner) CreateRenderer() fyne.WidgetRenderer {
	return nil
}

func (s *spinner) Destroy() {}

func (s *spinner) Layout() {}

func (s *spinner) MinSize() fyne.Size {
	return fyne.Size{}
}

func (s *spinner) Refresh() {}

func TestSpinnerBase_NewSpinnerBase(t *testing.T) {
	sp := &spinner{}
	s := NewSpinnerBase(sp, 2, 4, 1)

	assert.NotNil(t, s.spinner)
	assert.Equal(t, 2., s.Min)
	assert.Equal(t, 4., s.Max)
	assert.Equal(t, 1., s.Step)
	assert.Equal(t, 2., s.Value)
}

func TestSpinnerBase_clampValueToRange_InRange(t *testing.T) {
	sp := &spinner{}
	s := NewSpinnerBase(sp, 2, 4, 1)
	s.clampValueToRange()
	assert.Equal(t, 2., s.Value)

	s.Value = 3.
	s.clampValueToRange()
	assert.Equal(t, 3., s.Value)

	s.Value = 4.
	s.clampValueToRange()
	assert.Equal(t, 4., s.Value)
}

func TestSpinnerBase_clampValueToRange_GreaterThanMax(t *testing.T) {
	sp := &spinner{}
	s := NewSpinnerBase(sp, 2, 4, 1)
	s.Value = 5.
	s.clampValueToRange()
	assert.Equal(t, 4., s.Value)

	s.Max = -1
	s.clampValueToRange()
	assert.Equal(t, -1., s.Value)
}

func TestSpinnerBase_clampValueToRange_LessThanMin(t *testing.T) {
	sp := &spinner{}
	s := NewSpinnerBase(sp, 2, 4, 1)
	s.Value = 1.
	s.clampValueToRange()
	assert.Equal(t, 2., s.Value)
}

func TestSpinnerBase_SetValue(t *testing.T) {
	sp := &spinner{}
	s := NewSpinnerBase(sp, 2, 4, 1)
	s.SetValue(3.)
	assert.Equal(t, 3., s.Value)

	s.SetValue(5.)
	assert.Equal(t, 4., s.Value)

	s.SetValue(-1.)
	assert.Equal(t, 2., s.Value)
}
