package mallow

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/markdown"
)

// markdownOptions are how every Markdown shows; their theme is made anew
// when the palette, the font or the line height changes.
var (
	markdownOptions markdown.Options
	markdownLook    struct {
		text, accent ui.Color
		font         string
		lineHeight   float32
	}
)

// Markdown shows src, text written in Markdown, as it reads: headings,
// lists, strong text, code, tables, in the palette's colors. While
// streaming, src is still being written, as a language model's answer,
// and what is added to it comes in smoothly; show it without streaming,
// in the same place, once it is whole. Its text can be selected.
func Markdown(c *ui.Context, src string, streaming bool) ui.Element {
	theme()
	if streaming {
		return markdown.Stream(c, src, &markdownOptions)
	}
	return markdown.View(c, src, &markdownOptions)
}

// Document shows Markdown that came from elsewhere, as a project's README:
// as Markdown does, with its pictures and with its links the app's to
// follow. Keep one for what is shown in a place: it holds the pictures
// loaded.
type Document struct {
	// LoadImage loads the picture at an address, on a goroutine of its
	// own; until it has, and when it fails, the picture shows as its
	// description. Without it every picture does.
	LoadImage func(url string) (ui.ImageSource, error)
	// OnLink follows a link, given its address as it is written. Without
	// it links to the web open in the browser.
	OnLink  func(url string)
	options markdown.Options
}

// View shows src, the document's text.
func (d *Document) View(c *ui.Context, src string) ui.Element {
	theme()
	d.options.Theme, d.options.LoadImage, d.options.OnLink = markdownOptions.Theme, d.LoadImage, d.OnLink
	return markdown.View(c, src, &d.options)
}

// theme makes the look of every Markdown anew when the palette, the font
// or the line height changed.
func theme() {
	lines := LineHeight
	if lines == 0 {
		lines = 1.5
	}
	if l := &markdownLook; markdownOptions.Theme == nil || l.text != Text || l.accent != Accent || l.font != Font || l.lineHeight != lines {
		l.text, l.accent, l.font, l.lineHeight = Text, Accent, Font, lines
		line := Muted.Alpha(0.35)
		markdownOptions.Theme = &markdown.Theme{
			FontSize:            14,
			Font:                Font,
			LineHeight:          lines,
			Text:                Text,
			Muted:               Muted,
			Link:                Accent,
			Strong:              Text,
			Code:                Text,
			CodeBackground:      Text.Alpha(0.07),
			QuoteBar:            line,
			QuoteText:           Muted,
			Rule:                line,
			CodeBlockBackground: Text.Alpha(0.05),
			CodeBlockBorder:     ui.Transparent,
			CodeBlockRadius:     RadiusSmall,
			TableBorder:         line,
			TableHeader:         Text.Alpha(0.04),
			Checkbox:            Accent,
		}
	}
}
