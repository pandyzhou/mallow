package mallow

import "github.com/egoist/mygo/ui"

// Popover shows a small raised card below anchor while *open is true, for
// what is chosen there, as from a list: body builds its content, a column.
// A press outside the card and the anchor, or Escape, sets *open to false.
// Where there is no room below the anchor, it shows above it. It returns
// the card, none while closed.
func Popover(c *ui.Context, anchor ui.Element, open *bool, body func()) ui.Element {
	return ui.PopoverBase(c, anchor, open, func(panel ui.Element) {
		over(panel).Margin(6, 0).Padding(8).Gap(2).Radius(RadiusSmall + 4)
		Raised(panel, Medium)
		body()
	})
}
