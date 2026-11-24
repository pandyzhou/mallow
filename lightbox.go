package mallow

import (
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

// pictures builds a paragraph of a Document that is pictures alone, in
// place of what Markdown would: each as large as it is, at most as wide as
// the view, its description showing until it loads, and, loaded, a button
// that shows it large. It reports whether block i is such a paragraph.
func (d *Document) pictures(c *ui.Context, doc *parser.Document, i int) bool {
	if doc.Block(i).Kind != parser.Paragraph {
		return false
	}
	in := doc.Inlines(i)
	n := 0
	for k := 0; k < len(in); k = in[k].Next {
		switch in[k].Kind {
		case parser.Image:
			n++
		case parser.SoftBreak, parser.LineBreak:
		case parser.Text:
			if strings.TrimSpace(in[k].Text) != "" {
				return false
			}
		default:
			return false
		}
	}
	if n == 0 {
		return false
	}
	const em = 14
	ui.Row(c).Wrap().Gap(em * 0.5).AlignSelf(ui.Stretch).Children(func() {
		for k := 0; k < len(in); k = in[k].Next {
			if in[k].Kind != parser.Image {
				continue
			}
			alt := parser.PlainText(in[k+1 : in[k].Next])
			src := d.picture(c, in[k].Dest)
			if src == nil {
				label := alt
				if label == "" {
					label = "图片"
				}
				ph := ui.Row(c.Key(k)).Padding(em*0.5, em*0.75).Radius(em * 0.3).Background(Text.Alpha(0.07))
				ph.Role(ui.RoleImage).Label(alt)
				ph.Children(func() { ui.Text(c, label).TextColor(Muted).FontSize(em * 0.875) })
				continue
			}
			label := "看大图"
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
			if in[k].Title != "" {
				Tooltip(c, b, in[k].Title)
			}
			if b.Clicked() {
				d.zoomed, d.zoom, d.zoomLabel, d.zoomIn = true, src, alt, d.showing
			}
		}
	})
	return true
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
