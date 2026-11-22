package mallow

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Radii.
const (
	RadiusCard    = 24
	RadiusControl = 16
	RadiusSmall   = 12
	RadiusPill    = 999
)

// Depth is how far an element stands out of, or sinks into, the surface: the
// shadow offset, whose blur is twice that. Keep elements at least twice
// their depth apart, so that their shadows don't meet.
type Depth float32

const (
	Tiny   Depth = 2 // marks inside controls: check discs, knobs, segments
	Small  Depth = 4 // buttons, hover, small controls
	Medium Depth = 6 // cards in dense layouts
	Large  Depth = 8 // panels, standalone cards and trays
)

// Raised makes e stand out of the surface.
func Raised(e ui.Element, d Depth) ui.Element {
	o := float32(d)
	return e.Background(Base).
		Shadow(o, o, 2*o, 0, Dark).
		Shadow(-o, -o, 2*o, 0, Light)
}

// Recessed makes e sink into the surface.
func Recessed(e ui.Element, d Depth) ui.Element {
	o := float32(d)
	return e.Background(Base).
		InsetShadow(o, o, 2*o, 0, Dark).
		InsetShadow(-o, -o, 2*o, 0, Light)
}

// Soft fades shadow, color and position changes over 200 ms.
func Soft(e ui.Element) ui.Element {
	return e.Transition(ui.ElementTransition{Duration: 200 * time.Millisecond})
}

// Card is a raised column: a panel or a card, as e.g.
// mallow.Card(c, mallow.Large).Padding(24).Children(…).
func Card(c *ui.Context, d Depth) ui.Element {
	return Raised(ui.Column(c).Radius(RadiusCard), d)
}

// Tray is a recessed column: a well that sets what it holds apart from what
// is around it, as raised cards rising out of it.
func Tray(c *ui.Context, d Depth) ui.Element {
	return Recessed(ui.Column(c).Radius(RadiusCard+8), d)
}

// Groove is a divider pressed into the surface: a thin channel, the way soft
// UI separates parts without a line of another color.
func Groove(c *ui.Context) ui.Element {
	return ui.Box(c).Height(4).Shrink(0).Radius(2).Background(Base).
		InsetShadow(1, 1, 2, 0, Dark).
		InsetShadow(-1, -1, 2, 0, Light)
}

// Label is the small caption above a group of controls.
func Label(c *ui.Context, s string) ui.Element {
	return ui.Text(c, s).FontSize(12).FontWeight(500).TextColor(Muted)
}
