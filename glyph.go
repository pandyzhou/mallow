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

// GlyphTrash is a bin with its lid, for deleting.
func GlyphTrash(c *ui.Context, color ui.Color) {
	ui.Column(c).AlignItems(ui.Center).Gap(1.5).Children(func() {
		ui.Box(c).Size(11, 1.5).Radius(1).Background(color)
		ui.Box(c).Size(8, 8).Radius(2).Border(1.5, color)
	})
}

// GlyphCalendar is a calendar's page: a day, or dates.
func GlyphCalendar(c *ui.Context, color ui.Color) {
	ui.Column(c).Size(13, 12).Radius(3).Border(1.5, color).Padding(2.5, 0, 0, 0).Children(func() {
		ui.Box(c).Height(1.5).Background(color)
	})
}

// GlyphInbox is a tray: where things land before they are sorted.
func GlyphInbox(c *ui.Context, color ui.Color) {
	ui.Column(c).Size(13, 11).Radius(3).Border(1.5, color).Justify(ui.End).Padding(0, 2.5, 1.5, 2.5).Children(func() {
		ui.Box(c).Height(1.5).Radius(1).Background(color)
	})
}

// GlyphList is three lines: a list.
func GlyphList(c *ui.Context, color ui.Color) {
	ui.Column(c).Width(12).Gap(2.5).Children(func() {
		for range 3 {
			ui.Box(c).Height(1.5).Radius(1).Background(color)
		}
	})
}

// GlyphDone is a check in a ring: what is done.
func GlyphDone(c *ui.Context, color ui.Color) {
	ui.Row(c).Size(14, 14).Radius(7).Border(1.5, color).Justify(ui.Center).AlignItems(ui.Center).Children(func() {
		ui.Text(c, "✓").FontSize(8).FontWeight(700).TextColor(color)
	})
}
