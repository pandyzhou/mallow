package mallow

import "github.com/egoist/mygo/ui"

// Tooltip shows text by anchor while the pointer rests on it or the keyboard
// focus is on it: a small raised card floating over the surface.
func Tooltip(c *ui.Context, anchor ui.Element, text string) {
	ui.TooltipBase(c, anchor, func(tip ui.Element) {
		tip.Margin(6).Padding(6, 10).Radius(RadiusSmall)
		Raised(tip, Small).Children(func() {
			ui.Text(c, text).FontSize(12).TextColor(Text).SingleLine()
		})
	})
}
