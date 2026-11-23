package mallow

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
)

// DateLayout is how Calendar writes a date: "2006-01-02".
const DateLayout = "2006-01-02"

var weekdayNames = [7]string{"一", "二", "三", "四", "五", "六", "日"}

// Calendar shows the month of *month, weeks starting on Monday, and picks a
// day into *value, as a DateLayout date ("" for none). Its arrows turn
// *month to the month before and after. today is marked. It reports
// whether *value changed.
func Calendar(c *ui.Context, label string, month *time.Time, value *string, today time.Time) bool {
	y, m, _ := month.Date()
	first := time.Date(y, m, 1, 0, 0, 0, 0, month.Location())
	changed := false
	ui.Column(c.Key(label)).Label(label).Gap(8).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(8).Children(func() {
			ui.Text(c, fmt.Sprintf("%d 年 %d 月", y, int(m))).Grow(1).FontSize(14).FontWeight(600).TextColor(Text)
			if IconButton(c, "上个月", GlyphChevronLeft, IconButtonStyle{}).Clicked() {
				*month = first.AddDate(0, -1, 0)
			}
			if IconButton(c, "下个月", GlyphChevronRight, IconButtonStyle{}).Clicked() {
				*month = first.AddDate(0, 1, 0)
			}
		})
		ui.Row(c).Children(func() {
			for _, n := range weekdayNames {
				ui.Text(c, n).Grow(1).Basis(0).FontSize(12).TextColor(Muted).TextAlign(ui.Center)
			}
		})
		// Six weeks hold any month; days of the months around fill the rest.
		day := first.AddDate(0, 0, -((int(first.Weekday()) + 6) % 7))
		for w := range 6 {
			ui.Row(c.Key(w)).Children(func() {
				for range 7 {
					if calendarDay(c, day, day.Month() == m, day.Format(DateLayout) == *value, sameDay(day, today)).Clicked() {
						*value, changed = day.Format(DateLayout), true
					}
					day = day.AddDate(0, 0, 1)
				}
			})
		}
	})
	return changed
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}

// calendarDay is a day of a Calendar: pressed in and in the accent while
// chosen, in the accent today, muted outside the month.
func calendarDay(c *ui.Context, day time.Time, inMonth, chosen, today bool) ui.Element {
	label := fmt.Sprintf("%d月%d日", int(day.Month()), day.Day())
	cell := ui.Row(c.Key(day.Format(DateLayout))).Grow(1).Basis(0).Justify(ui.Center)
	var b ui.Element
	cell.Children(func() {
		b = ui.ButtonBase(c).Label(label).Size(30, 30).Radius(RadiusPill).
			AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
		color, weight := Text, 400
		switch {
		case chosen:
			Recessed(b, Tiny)
			color, weight = Accent, 600
		case b.Hovered():
			Raised(b, Tiny)
		}
		switch {
		case chosen:
		case today:
			color, weight = Accent, 600
		case !inMonth:
			color = Muted.Alpha(0.55)
		}
		Soft(b).Children(func() {
			ui.Text(c, fmt.Sprint(day.Day())).FontSize(12).FontWeight(weight).TextColor(color)
		})
	})
	return b
}
