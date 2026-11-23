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

// Recessed makes e sink into the surface, with inner shadows. They cost the
// GPU every pixel of e, the middle too, so keep it for controls; a large
// well takes RecessedWide.
func Recessed(e ui.Element, d Depth) ui.Element {
	o := float32(d)
	return e.Background(Base).
		InsetShadow(o, o, 2*o, 0, Dark).
		InsetShadow(-o, -o, 2*o, 0, Light)
}

// RecessedWide makes e, rounded by radius, sink into the surface as Recessed
// does, for large areas. Inner shadows are evaluated over every pixel of an
// element, so a well the size of half a window costs the GPU milliseconds a
// frame; here strips just outside each edge cast ordinary outer shadows into
// e, clipped to its corners, and outer shadows cost only the band where they
// show.
func RecessedWide(e ui.Element, d Depth, radius float32) ui.Element {
	o := float32(d)
	blur := 2 * o
	// A strip thicker than the blur's reach casts as a whole half-plane would.
	t := 3*blur + o
	return e.Background(Base).Radius(radius).Draw(func(p *ui.Painter, r ui.Rect) {
		p.Clip(r, radius, func() {
			x0, y0, x1, y1 := r.X-t, r.Y-t, r.X+r.W+t, r.Y+r.H+t
			// The shadow above and to the left, as the light comes from the
			// top left; the highlight below and to the right.
			p.Shadow(ui.Rect{X: x0, Y: y0, W: x1 - x0, H: t}, 0, o, o, blur, 0, Dark)
			p.Shadow(ui.Rect{X: x0, Y: y0, W: t, H: y1 - y0}, 0, o, o, blur, 0, Dark)
			p.Shadow(ui.Rect{X: x0, Y: r.Y + r.H, W: x1 - x0, H: t}, 0, -o, -o, blur, 0, Light)
			p.Shadow(ui.Rect{X: r.X + r.W, Y: y0, W: t, H: y1 - y0}, 0, -o, -o, blur, 0, Light)
		})
	})
}

// Slide is the motion of what moves from one place to another, as a
// NavList's recess: a quick spring, about as fast as macOS's selection,
// settling with a hint of bounce rather than swinging back and forth.
var Slide = ui.Animation{Spring: ui.SpringOptions{Duration: 280 * time.Millisecond, Bounce: 0.08}}

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
	return RecessedWide(ui.Column(c), d, RadiusCard+8)
}

// Groove is a divider pressed into the surface: a thin channel, the way soft
// UI separates parts without a line of another color.
func Groove(c *ui.Context) ui.Element {
	return ui.Box(c).Height(4).Shrink(0).Radius(2).Background(Base).
		InsetShadow(1, 1, 2, 0, Dark).
		InsetShadow(-1, -1, 2, 0, Light)
}

// GrooveV is Groove standing up, between the parts of a row: it is as tall
// as the row when the row stretches its children.
func GrooveV(c *ui.Context) ui.Element {
	return ui.Box(c).Width(4).Shrink(0).Radius(2).Background(Base).
		InsetShadow(1, 1, 2, 0, Dark).
		InsetShadow(-1, -1, 2, 0, Light)
}

// Label is the small caption above a group of controls.
func Label(c *ui.Context, s string) ui.Element {
	return ui.Text(c, s).FontSize(12).FontWeight(500).TextColor(Muted)
}
