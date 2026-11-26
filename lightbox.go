package mallow

import (
	"cmp"
	"strings"
	"sync"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/markdown/parser"
)

// lightboxMargin is the least room between a picture shown large and the
// window's edges.
const lightboxMargin = 40

// Lightbox shows a picture large, over the window, while *open: as large
// as it is, up to twice that for a small one, and no larger than fits. A
// click anywhere, or Escape, shuts it.
func Lightbox(c *ui.Context, open *bool, src ui.ImageSource, label string) ui.Element {
	return ui.DialogBase(c, open, func(backdrop, panel ui.Element) {
		w, h := ui.ImageSize(src)
		if w <= 0 || h <= 0 {
			*open = false
			return
		}
		ww, wh := c.Size()
		scale := min((ww-2*lightboxMargin)/w, (wh-2*lightboxMargin)/h, 2)
		panel.Label(label)
		// The page fades behind it, as behind a dialog, and more.
		veil := Canvas.Alpha(0.92)
		pose := ui.Pose{}
		entered := ui.Local(panel, "entered", func() bool { return false })
		if !*entered {
			veil, pose, *entered = Canvas.Alpha(0), ui.Pose{Scale: 0.96, Fade: 1}, true
			c.Invalidate()
		}
		backdrop.Background(backdrop.AnimateColor("veil", veil, ui.Animation{Duration: 180 * time.Millisecond}))
		panel.AnimatePose(pose, ui.Snappy)
		shut := ui.ButtonBase(c).Label("关闭大图").Cursor(ui.CursorPointer).Radius(RadiusSmall).Clip()
		shut.Children(func() {
			ui.Image(c, src).Size(w*scale, h*scale)
		})
		if shut.Clicked() {
			*open = false
		}
	})
}

// picture is a picture of a Document, loading or loaded.
type picture struct {
	mu   sync.Mutex
	src  ui.ImageSource
	done bool
}

// pictures builds a paragraph of a Document that is pictures alone, each by
// itself or the whole of a link, in place of what Markdown would: each as
// large as it is, at most as wide as the view, its description showing until
// it loads. Loaded, a picture by itself is a button that shows it large; one
// that is a link follows it, as on GitHub, but shows large where the link
// only goes to a heading of the page (#…). It reports whether block i is
// such a paragraph.
func (d *Document) pictures(c *ui.Context, doc *parser.Document, i int) bool {
	if doc.Block(i).Kind != parser.Paragraph {
		return false
	}
	in := doc.Inlines(i)
	// The pictures, where they are among the inlines, each with the link it
	// is the whole of, -1 for none.
	type shown struct{ at, link int }
	var all []shown
	for k := 0; k < len(in); k = in[k].Next {
		switch in[k].Kind {
		case parser.Image:
			all = append(all, shown{k, -1})
		case parser.Link:
			at := linked(in, k)
			if at < 0 {
				return false
			}
			all = append(all, shown{at, k})
		case parser.SoftBreak, parser.LineBreak:
		case parser.Text:
			if strings.TrimSpace(in[k].Text) != "" {
				return false
			}
		default:
			return false
		}
	}
	if len(all) == 0 {
		return false
	}
	const em = 14
	ui.Row(c).Wrap().Gap(em * 0.5).AlignSelf(ui.Stretch).Children(func() {
		for _, p := range all {
			k := p.at
			alt := parser.PlainText(in[k+1 : in[k].Next])
			src := d.picture(c, in[k].Dest)
			title, link := in[k].Title, ""
			if p.link >= 0 {
				link = in[p.link].Dest
				title = cmp.Or(title, in[p.link].Title)
			}
			// A link to a heading of the page has nowhere else to go.
			follows := link != "" && !strings.HasPrefix(link, "#")
			if src == nil {
				label := alt
				if label == "" {
					label = "图片"
				}
				// Still a link while it shows as its description.
				var ph ui.Element
				if follows {
					ph = ui.ButtonBase(c.Key(k)).Label("打开链接：" + label).Cursor(ui.CursorPointer)
				} else {
					ph = ui.Row(c.Key(k)).Role(ui.RoleImage).Label(alt)
				}
				ph.Padding(em*0.5, em*0.75).Radius(em * 0.3).Background(Text.Alpha(0.07))
				ph.Children(func() { ui.Text(c, label).TextColor(Muted).FontSize(em * 0.875) })
				if follows && ph.Clicked() {
					d.follow(c, link)
				}
				continue
			}
			label := "看大图"
			if follows {
				label = "打开链接"
			}
			if alt != "" {
				label += "：" + alt
			}
			b := ui.ButtonBase(c.Key(k)).Label(label).Cursor(ui.CursorPointer).Radius(em * 0.3).Clip().MaxWidthPercent(100).Shrink(1)
			b.Children(func() {
				img := ui.Image(c, src).MaxWidthPercent(100)
				if w, h := ui.ImageSize(src); w > 0 && h > 0 {
					img.AspectRatio(w / h)
				}
			})
			if title != "" {
				Tooltip(c, b, title)
			}
			if b.Clicked() {
				if follows {
					d.follow(c, link)
				} else {
					d.zoomed, d.zoom, d.zoomLabel, d.zoomIn = true, src, alt, d.showing
				}
			}
		}
	})
	return true
}

// linked returns where the picture is among the inlines of a link, the one
// at k, that the picture is the whole of: -1 when the link holds anything
// else, or more pictures.
func linked(in []parser.Inline, k int) int {
	at := -1
	for j := k + 1; j < in[k].Next; j = in[j].Next {
		switch in[j].Kind {
		case parser.Image:
			if at >= 0 {
				return -1
			}
			at = j
		case parser.SoftBreak, parser.LineBreak:
		case parser.Text:
			if strings.TrimSpace(in[j].Text) != "" {
				return -1
			}
		default:
			return -1
		}
	}
	return at
}

// follow follows a link of the document as OnLink says, or, without it, as
// Markdown does: in the browser, where it is to the web.
func (d *Document) follow(c *ui.Context, url string) {
	if d.OnLink != nil {
		d.OnLink(url)
		return
	}
	if ui.SafeToOpen(url) {
		c.Services().OpenURL(url)
	}
}

// picture returns the picture at an address once it is loaded, and begins
// to load it the first time: nil until then, and for one that does not load.
func (d *Document) picture(c *ui.Context, url string) ui.ImageSource {
	if d.LoadImage == nil || url == "" {
		return nil
	}
	d.mu.Lock()
	if d.loaded == nil {
		d.loaded = map[string]*picture{}
	}
	p := d.loaded[url]
	fresh := p == nil
	if fresh {
		p = &picture{}
		d.loaded[url] = p
	}
	d.mu.Unlock()
	if fresh {
		load, services := d.LoadImage, c.Services()
		go func() {
			src, err := load(url)
			p.mu.Lock()
			if err == nil {
				p.src = src
			}
			p.done = true
			p.mu.Unlock()
			services.Invalidate()
		}()
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.src
}
