package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Spinnable is an interface for specifying if a widget is spinnable (i.e. is a spinner).
type Spinnable interface {
	CreateRenderer() fyne.WidgetRenderer
}

// SpinnerBase is the base widget for all spinner widgets.
type SpinnerBase struct {
	widget.DisableableWidget

	spinner Spinnable
	Value   float64
	Min     float64
	Max     float64
	Step    float64
}

// NewSpinnerBase iniitializes a new SpinnerBase widget. This function should only be
// called from within a child spinner New function. See, for example, the Spinner widget.
//
// Params:
//
//	min is the minimum value for the spinner.
//	max is the maximum value for the spinner.
//	step is the step size for the spinner.
//
// Returns:
//
//	A new SpinnerBase widget.
func NewSpinnerBase(sp Spinnable, min, max, step float64) *SpinnerBase {
	s := &SpinnerBase{spinner: sp, Min: min, Max: max, Step: step, Value: min}
	s.ExtendBaseWidget(s)
	return s
}

// CreateRenderer is a private method to fyne which links this widget to its
// renderer.
func (s *SpinnerBase) CreateRenderer() fyne.WidgetRenderer {
	return s.spinner.CreateRenderer()
}

func (s *SpinnerBase) clampValueToRange() {
	if s.Value > s.Max {
		s.Value = s.Max
	} else if s.Value < s.Min {
		s.Value = s.Min
	}
}
