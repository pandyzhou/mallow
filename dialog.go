package mallow

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// DialogOptions describe a Dialog.
type DialogOptions struct {
	// Title heads the card and names the dialog; "" for neither.
	Title string
	// Width is the card's width, 440 when zero.
	Width float32
}

// dialogMargin is the least room a dialog's card leaves to the window's
// edges.
const dialogMargin = 24

// Dialog shows a card over the window while *open is true, for what is
// asked or filled in apart from the page: the page behind it fades into
// the surface and cannot be used, and the keyboard stays in the card.
// body builds the card's content, a column; give the field to write in
// first AutoFocus. A click outside the card or Escape sets *open to false,
// which Dismissed reports on the card, the element returned: none while
// closed. The card rises into place as it shows, and is never larger than
// the window: what may grow taller than that goes in a Scroll.
func Dialog(c *ui.Context, open *bool, o DialogOptions, body func()) ui.Element {
	return ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		width := o.Width
		if width == 0 {
			width = 440
		}
		w, h := c.Size()
		over(panel).Label(o.Title).Width(width).MaxWidth(w - 2*dialogMargin).MaxHeight(h - 2*dialogMargin).
			Padding(24).Gap(16).Radius(RadiusCard)
		if look == Organic {
			// A card over the veil, with a line around it and the shadow
			// of what floats above the page: deeper than a card's on it.
			shade := Text.Alpha(0.18)
			if isDark {
				shade = ui.RGBA(0, 0, 0, 0.45)
			}
			panel.Background(Base).Border(1, Line).Shadow(0, 16, 48, 0, shade)
		} else {
			Raised(panel, Large)
		}
		// The page fades into the color of the surface, the window's: the
		// card's light and shadow fall on it as on the surface itself,
		// which a veil of black would not let them.
		veil := Canvas.Alpha(0.78)
		pose := ui.Pose{}
		// The first frame poses the card below its place, faded, and the
		// veil clear: the next ones bring them in from there.
		entered := ui.Local(panel, "entered", func() bool { return false })
		if !*entered {
			veil, pose, *entered = Canvas.Alpha(0), ui.Pose{Y: 10, Fade: 1}, true
			c.Invalidate()
		}
		backdrop.Background(backdrop.AnimateColor("veil", veil, ui.Animation{Duration: 180 * time.Millisecond}))
		panel.AnimatePose(pose, ui.Snappy)
		if o.Title != "" {
			ui.Text(c, o.Title).FontSize(18).FontWeight(600).TextColor(Text)
		}
		body()
	})
}
