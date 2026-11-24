package mallow

import (
	"math"

	"github.com/egoist/mygo/ui"
)

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

// GlyphTrash is a bin: its lid with a handle, and its body with two
// slats, for deleting.
func GlyphTrash(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		bin := new(ui.Path)
		// The lid, and the handle on it.
		bin.MoveTo(x+2.5, y+4.5).LineTo(x+13.5, y+4.5)
		bin.MoveTo(x+6, y+4.5).LineTo(x+6, y+2.5).LineTo(x+10, y+2.5).LineTo(x+10, y+4.5)
		// The body, narrower at the bottom.
		bin.MoveTo(x+4, y+4.5).LineTo(x+4.7, y+13.5).LineTo(x+11.3, y+13.5).LineTo(x+12, y+4.5)
		// The slats.
		bin.MoveTo(x+6.7, y+7).LineTo(x+6.9, y+11)
		bin.MoveTo(x+9.3, y+7).LineTo(x+9.1, y+11)
		p.StrokePath(bin, 1.3, color)
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

// drawn is a glyph drawn as paths in a box of 16 by 16: draw builds them
// with the box's top left corner at (x, y).
func drawn(c *ui.Context, draw func(p *ui.Painter, x, y float32)) {
	ui.Box(c).Size(16, 16).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) { draw(p, r.X, r.Y) })
}

// GlyphChevronDown points down: what opens a list to choose from.
func GlyphChevronDown(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		p.StrokePath(new(ui.Path).MoveTo(x+4.5, y+6.5).LineTo(x+8, y+10).LineTo(x+11.5, y+6.5), 1.4, color)
	})
}

// GlyphCopy is two sheets, one behind the other: for copying.
func GlyphCopy(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		sheets := new(ui.Path)
		// The sheet in front, and what shows of the one behind it.
		sheets.MoveTo(x+3, y+6).LineTo(x+10, y+6).LineTo(x+10, y+13.5).LineTo(x+3, y+13.5).Close()
		sheets.MoveTo(x+6, y+3.8).LineTo(x+6, y+2.5).LineTo(x+13, y+2.5).LineTo(x+13, y+10).LineTo(x+12, y+10)
		p.StrokePath(sheets, 1.3, color)
	})
}

// GlyphStop is a filled square: for stopping what is under way.
func GlyphStop(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(9, 9).Radius(2).Background(color)
}

// GlyphHistory is a clock: what was before, as earlier conversations.
func GlyphHistory(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		clock := new(ui.Path).Circle(x+8, y+8, 5.6)
		clock.MoveTo(x+8, y+4.8).LineTo(x+8, y+8).LineTo(x+10.3, y+9.4)
		p.StrokePath(clock, 1.3, color)
	})
}

// GlyphAttach is a paper clip: for attaching files to what is written.
func GlyphAttach(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		// From the end of the outer wire, round the bottom, up and round
		// the top, and down into the turn in the middle.
		clip := new(ui.Path).MoveTo(x+11, y+5).LineTo(x+11, y+11).CubeTo(x+11, y+15, x+5, y+15, x+5, y+11).
			LineTo(x+5, y+4.5).CubeTo(x+5, y+1.8, x+9, y+1.8, x+9, y+4.5).
			LineTo(x+9, y+10.5).CubeTo(x+9, y+11.9, x+7, y+11.9, x+7, y+10.5).LineTo(x+7, y+5)
		p.StrokePath(clip, 1.3, color)
	})
}

// GlyphFile is a sheet with its corner folded: a file.
func GlyphFile(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		sheet := new(ui.Path).MoveTo(x+4, y+2.5).LineTo(x+9.5, y+2.5).LineTo(x+12.5, y+5.5).
			LineTo(x+12.5, y+13.5).LineTo(x+4, y+13.5).Close()
		sheet.MoveTo(x+9.5, y+2.5).LineTo(x+9.5, y+5.5).LineTo(x+12.5, y+5.5)
		p.StrokePath(sheet, 1.3, color)
	})
}

// GlyphArrowUp is an arrow pointing up: for sending what was written.
func GlyphArrowUp(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		arrow := new(ui.Path).MoveTo(x+8, y+13).LineTo(x+8, y+3.5)
		arrow.MoveTo(x+3.8, y+7.6).LineTo(x+8, y+3.4).LineTo(x+12.2, y+7.6)
		p.StrokePath(arrow, 1.6, color)
	})
}

// GlyphCompose is a sheet with a pencil over its corner: for beginning
// something new to write, as a conversation.
func GlyphCompose(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		sheet := new(ui.Path).MoveTo(x+7.5, y+3).LineTo(x+3, y+3).LineTo(x+3, y+13).LineTo(x+13, y+13).LineTo(x+13, y+8.5)
		sheet.MoveTo(x+6.6, y+9.4).LineTo(x+7.2, y+7.2).LineTo(x+11.9, y+2.5).LineTo(x+13.5, y+4.1).LineTo(x+8.8, y+8.8).Close()
		p.StrokePath(sheet, 1.3, color)
	})
}

// GlyphEdit is a pencil: for changing what is written, as a name.
func GlyphEdit(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		// The pencil, from its point at the bottom left up to its end,
		// and the line where the end's band begins.
		pen := new(ui.Path).MoveTo(x+2.5, y+13.5).LineTo(x+3.2, y+10.4).LineTo(x+10.8, y+2.8).
			LineTo(x+13.2, y+5.2).LineTo(x+5.6, y+12.8).Close()
		pen.MoveTo(x+9.2, y+4.4).LineTo(x+11.6, y+6.8)
		p.StrokePath(pen, 1.3, color)
	})
}

// GlyphSettings is a gear: the settings.
func GlyphSettings(c *ui.Context, color ui.Color) {
	drawn(c, func(p *ui.Painter, x, y float32) {
		const (
			teeth        = 8
			inner, outer = 4.7, 6.7   // the radii of the wheel and of its teeth
			foot, top    = 0.21, 0.13 // half a tooth's angle, at each of them
		)
		gear := new(ui.Path)
		at := func(angle, radius float64, first bool) {
			px, py := x+8+float32(radius*math.Cos(angle)), y+8+float32(radius*math.Sin(angle))
			if first {
				gear.MoveTo(px, py)
			} else {
				gear.LineTo(px, py)
			}
		}
		for i := range teeth {
			a := 2 * math.Pi * float64(i) / teeth
			at(a-foot, inner, i == 0)
			at(a-top, outer, false)
			at(a+top, outer, false)
			at(a+foot, inner, false)
		}
		gear.Close()
		gear.Circle(x+8, y+8, 2.1)
		p.StrokePath(gear, 1.3, color)
	})
}
