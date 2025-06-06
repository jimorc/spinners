package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Spinnable is an interface for specifying if a widget is spinnable (i.e. is a spinner).
type Spinnable interface {
	CreateRenderer() fyne.WidgetRenderer
	Refresh()
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

// SetValue sets the Value field and ensures that it is between Min and Max.
func (s *SpinnerBase) SetValue(value float64) {
	if value == s.Value {
		return
	}
	s.Value = s.clampValue(value)
	s.spinner.Refresh()
}

// clampValueToRange ensures Value is clamped between Min and Max.
func (s *SpinnerBase) clampValueToRange() {
	s.Value = s.clampValue(s.Value)
}

// clampValue returns a value that is clamped between Min and Max.
func (s *SpinnerBase) clampValue(value float64) float64 {
	if value > s.Max {
		value = s.Max
	} else if value < s.Min {
		value = s.Min
	}
	return value
}
