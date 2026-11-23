package mallow

import "github.com/egoist/mygo/ui"

// Glyphs drawn from boxes and text, sized for IconButton.

// GlyphMinimize is a short bar.
func GlyphMinimize(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(10, 1.5).Radius(1).Background(color)
}

// GlyphMaximize is a square outline.
func GlyphMaximize(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(9, 9).Radius(2).Border(1.5, color)
}

// GlyphRestore is two squares, one behind the other.
func GlyphRestore(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(11, 11).Children(func() {
		ui.Box(c).Absolute().Top(0).Right(0).Size(7, 7).Radius(1.5).Border(1.5, color)
		ui.Box(c).Absolute().Left(0).Bottom(0).Size(7, 7).Radius(1.5).Border(1.5, color).Background(Base)
	})
}

// GlyphClose is a cross.
func GlyphClose(c *ui.Context, color ui.Color) {
	ui.Text(c, "✕").FontSize(11).TextColor(color)
}

// GlyphChevronLeft and GlyphChevronRight point back and forward.
func GlyphChevronLeft(c *ui.Context, color ui.Color) {
	ui.Text(c, "‹").FontSize(18).TextColor(color).Margin(0, 0, 2, 0)
}

func GlyphChevronRight(c *ui.Context, color ui.Color) {
	ui.Text(c, "›").FontSize(18).TextColor(color).Margin(0, 0, 2, 0)
}

// GlyphPlus is a plus.
func GlyphPlus(c *ui.Context, color ui.Color) {
	ui.Text(c, "+").FontSize(16).TextColor(color)
}

// GlyphCheck is a check mark.
func GlyphCheck(c *ui.Context, color ui.Color) {
	ui.Text(c, "✓").FontSize(11).FontWeight(600).TextColor(color)
}

// GlyphRepeat is a circling arrow, for what comes back, as a weekly item.
func GlyphRepeat(c *ui.Context, color ui.Color) {
	ui.Text(c, "↻").FontSize(14).FontWeight(600).TextColor(color)
}
