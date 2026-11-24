// Gallery shows every mallow component, in the light and the dark palette.
// It runs as a desktop app and, compiled to WebAssembly, in a browser: the
// GitHub Pages site of the repository.
//
//	go run ./cmd/gallery                                  # desktop
//	go tool mygo dev -web cmd/gallery                     # browser, live reload
//	go tool mygo build -platform web cmd/gallery          # static site in cmd/gallery/build/web
package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/pandyzhou/mallow"
)

type gallery struct {
	win mallow.Window

	mode    string
	style   string
	nav     string
	period  string
	checks  [3]bool
	sync    bool
	notify  bool
	name    string
	note    string
	clicks  int
	primary int
	repeat  bool

	tasks      [3]bool
	chosen     int
	chip       string
	memo       string
	calMonth   time.Time
	calDate    string
	source     string
	menuOpen   bool
	menuChoice string

	popOpen   bool   // the popover shows
	popChoice string // what was chosen in it

	dialogOpen bool
	dialogName string // what is typed in the dialog
	dialogMade string // what the dialog last made
}

// galleryMarkdown is what the gallery shows as Markdown.
const galleryMarkdown = "## 今天的安排\n\n上午有 **两节课**，下午要交 `实验报告`。\n\n- 高等数学，A301\n- 大学英语，B102\n\n> 记得带学生证。"

func newGallery() *gallery {
	return &gallery{
		mode:   string(mallow.ModeSystem),
		style:  string(mallow.StyleNeumorphism),
		nav:    "组件",
		period: "week",
		checks: [3]bool{true, false, false},
		sync:   true,

		chosen:   1,
		chip:     "工作",
		calMonth: time.Now(),
	}
}

func (g *gallery) view(c *ui.Context) {
	mallow.UseStyle(c, mallow.Style(g.style), mallow.Mode(g.mode))
	ui.Column(c).Fill().Children(func() {
		mallow.TitleBar(c, g.win, func() {
			ui.Text(c, "mallow").FontSize(13).FontWeight(600).TextColor(mallow.Muted)
		})
		ui.Row(c).Grow(1).MinHeight(0).AlignItems(ui.Stretch).Gap(28).Padding(4, 24, 24, 24).Children(func() {
			g.sidebar(c)
			ui.Scroll(c).Grow(1).MinWidth(0).Padding(8, 12, 24, 12).Gap(32).Children(func() {
				g.sections(c)
			})
		})
	})
	g.dialog(c)
}

// dialog is the Dialog the gallery shows: it asks for a name, and makes
// something of it.
func (g *gallery) dialog(c *ui.Context) {
	mallow.Dialog(c, &g.dialogOpen, mallow.DialogOptions{Title: "新建相册", Width: 400}, func() {
		make := func() {
			if name := strings.TrimSpace(g.dialogName); name != "" {
				g.dialogMade, g.dialogOpen = name, false
			}
		}
		name := mallow.TextInput(c, &g.dialogName).Placeholder("相册的名字，回车确认").Label("相册的名字").AutoFocus()
		if name.Submitted() {
			make()
		}
		ui.Row(c).Gap(12).Justify(ui.End).Children(func() {
			if mallow.Button(c, "取消").Clicked() {
				g.dialogOpen = false
			}
			if mallow.PrimaryButton(c, "建好").Clicked() {
				make()
			}
		})
	})
}

func (g *gallery) sidebar(c *ui.Context) {
	mallow.Sidebar(c, mallow.SidebarOptions{
		Title:    "mallow",
		Subtitle: "新拟态组件库",
		Width:    200,
		Items:    []mallow.Item{{ID: "组件", Label: "组件"}, {ID: "表单", Label: "表单"}, {ID: "设置", Label: "设置"}},
		Selected: &g.nav,
		Bottom:   []mallow.Item{{ID: "关于", Label: "关于"}},
		Footer: func() {
			mallow.Label(c, "风格").PaddingX(8)
			mallow.Segmented(c, "风格", &g.style, []mallow.Option{
				{Value: string(mallow.StyleNeumorphism), Label: "新拟态"},
				{Value: string(mallow.StyleOrganic), Label: "自然有机"},
			})
			mallow.Label(c, "外观").PaddingX(8)
			mallow.Segmented(c, "外观", &g.mode, []mallow.Option{
				{Value: string(mallow.ModeLight), Label: "浅色"},
				{Value: string(mallow.ModeDark), Label: "深色"},
				{Value: string(mallow.ModeSystem), Label: "系统"},
			})
		},
	})
}

// section is a titled raised card of examples.
func section(c *ui.Context, title string, body func()) {
	mallow.Card(c, mallow.Medium).Padding(24).Gap(18).Children(func() {
		ui.Text(c, title).FontSize(16).FontWeight(600).TextColor(mallow.Text)
		mallow.Groove(c)
		body()
	})
}

func (g *gallery) sections(c *ui.Context) {
	ui.Column(c).Gap(4).Children(func() {
		mallow.Label(c, "COMPONENTS")
		ui.Text(c, "组件一览").FontSize(28).FontWeight(600).TextColor(mallow.Text)
	})
	ui.Row(c).Gap(28).AlignItems(ui.Start).Children(func() {
		ui.Column(c).Grow(1).Basis(0).MinWidth(0).Gap(28).Children(func() {
			section(c, "按钮 Button", func() {
				ui.Row(c).Gap(18).Wrap().AlignItems(ui.Center).Children(func() {
					if mallow.Button(c, "普通按钮").Clicked() {
						g.clicks++
					}
					if mallow.PrimaryButton(c, "主要按钮").Clicked() {
						g.primary++
					}
					mallow.IconButton(c, "上一个", mallow.GlyphChevronLeft, mallow.IconButtonStyle{Size: 36})
					mallow.IconButton(c, "添加", mallow.GlyphPlus, mallow.IconButtonStyle{Size: 36})
					mallow.IconButton(c, "删除", mallow.GlyphClose, mallow.IconButtonStyle{Size: 36, Danger: true})
				})
				ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
					mallow.Label(c, "无底图标按钮")
					if mallow.GhostIconButton(c, "每周重复", mallow.GlyphRepeat, mallow.IconButtonStyle{Active: g.repeat}).Clicked() {
						g.repeat = !g.repeat
					}
					mallow.GhostIconButton(c, "移除", mallow.GlyphClose, mallow.IconButtonStyle{Danger: true})
					mallow.GhostIconButton(c, "丢弃", mallow.GlyphTrash, mallow.IconButtonStyle{Danger: true})
					mallow.GhostIconButton(c, "改名", mallow.GlyphEdit, mallow.IconButtonStyle{})
					mallow.GhostIconButton(c, "偏好", mallow.GlyphSettings, mallow.IconButtonStyle{})
				})
				ui.Text(c, fmt.Sprintf("普通按钮点了 %d 次，主要按钮点了 %d 次", g.clicks, g.primary)).
					FontSize(12).TextColor(mallow.Muted)
			})
			section(c, "勾选 Check", func() {
				for i, label := range []string{"写周报", "读书 30 分钟", "整理房间"} {
					mallow.Check(c.Key(i), &g.checks[i], label)
				}
			})
			section(c, "列表行 ListRow · 小标签 Chip", func() {
				for i, title := range []string{"买菜", "交水电费", "预约体检"} {
					mallow.ListRow(c.Key(i), g.chosen == i).Children(func() {
						mallow.CheckMark(c, &g.tasks[i], "完成"+title)
						if mallow.Plain(c, "打开"+title).Grow(1).Padding(6, 4).Children(func() {
							ui.Text(c, title).FontSize(14).TextColor(mallow.Text)
						}).Clicked() {
							g.chosen = i
						}
					})
				}
				ui.Row(c).Gap(10).Wrap().Children(func() {
					for _, name := range []string{"收集箱", "工作", "生活"} {
						if mallow.Chip(c, name, g.chip == name).Clicked() {
							g.chip = name
						}
					}
				})
			})
			section(c, "来源列表 SourceList", func() {
				ui.Column(c).Width(220).Children(func() {
					mallow.SourceList(c, "来源", &g.source, []mallow.Item{
						{ID: "today", Label: "今日安排", Glyph: mallow.GlyphCalendar, Note: "4"},
						{ID: "inbox", Label: "未整理", Glyph: mallow.GlyphInbox, Note: "2"},
						{ID: "list", Label: "读书清单", Glyph: mallow.GlyphList},
						{ID: "done", Label: "做完的", Glyph: mallow.GlyphDone},
					})
				})
			})
			section(c, "右键菜单 ContextMenu", func() {
				target := mallow.Button(c, "在这里点右键")
				switch mallow.ContextMenu(c, target, &g.menuOpen, []mallow.MenuEntry{
					{ID: "new", Label: "再来一个", Glyph: mallow.GlyphPlus},
					{}, // a line: what follows is about this one
					{ID: "rename", Label: "改个名字", Glyph: mallow.GlyphEdit},
					{ID: "delete", Label: "删掉它", Glyph: mallow.GlyphTrash, Danger: true},
				}) {
				case "new":
					g.menuChoice = "新建"
				case "rename":
					g.menuChoice = "改名"
				case "delete":
					g.menuChoice = "删除"
				}
				ui.Text(c, "选了："+g.menuChoice).FontSize(12).TextColor(mallow.Muted)
			})
			section(c, "开关 Switch", func() {
				mallow.Switch(c.Key("sync"), &g.sync, "自动同步")
				mallow.Switch(c.Key("notify"), &g.notify, "到点提醒")
			})
		})
		ui.Column(c).Grow(1).Basis(0).MinWidth(0).Gap(28).Children(func() {
			section(c, "分段选择 Segmented", func() {
				mallow.Segmented(c, "范围", &g.period, []mallow.Option{
					{Value: "day", Label: "日"}, {Value: "week", Label: "周"}, {Value: "month", Label: "月"},
				})
				ui.Text(c, "当前："+g.period).FontSize(12).TextColor(mallow.Muted)
			})
			section(c, "输入框 TextInput", func() {
				mallow.TextInput(c, &g.name).Placeholder("你的名字").Label("名字")
				mallow.TextInput(c, &g.note).Placeholder("回车提交").Label("备注")
			})
			section(c, "多行输入 TextArea", func() {
				mallow.TextArea(c, &g.memo).Placeholder("备注，可以写多行").Label("多行备注")
			})
			section(c, "对话框 Dialog", func() {
				ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
					if mallow.Button(c, "打开对话框").Clicked() {
						g.dialogOpen, g.dialogName = true, ""
					}
					made := g.dialogMade
					if made == "" {
						made = "还没有"
					}
					ui.Text(c, "建好的相册："+made).FontSize(12).TextColor(mallow.Muted)
				})
			})
			section(c, "弹出层 Popover", func() {
				ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
					choice := g.popChoice
					if choice == "" {
						choice = "选一种水果"
					}
					anchor := mallow.Button(c, choice)
					if anchor.Clicked() {
						g.popOpen = !g.popOpen
					}
					mallow.Popover(c, anchor, &g.popOpen, func() {
						for _, fruit := range []string{"苹果", "橘子", "葡萄"} {
							item := mallow.Plain(c.Key(fruit), "选"+fruit).Padding(6, 12).MinWidth(120)
							item.Children(func() { ui.Text(c, fruit).FontSize(13).TextColor(mallow.Text) })
							if item.Clicked() {
								g.popChoice, g.popOpen = fruit, false
							}
						}
					})
					ui.Row(c).Gap(10).AlignItems(ui.Center).Children(func() {
						for _, glyph := range []func(*ui.Context, ui.Color){mallow.GlyphChevronDown, mallow.GlyphCopy, mallow.GlyphStop, mallow.GlyphHistory} {
							glyph(c, mallow.Muted)
						}
					})
				})
			})
			section(c, "Markdown", func() {
				mallow.Markdown(c, galleryMarkdown, false)
			})
			section(c, "日历 Calendar", func() {
				mallow.Calendar(c, "日期", &g.calMonth, &g.calDate, time.Now())
				date := g.calDate
				if date == "" {
					date = "未选"
				}
				ui.Text(c, "选中："+date).FontSize(12).TextColor(mallow.Muted)
			})
			section(c, "表面 Surface", func() {
				ui.Row(c).Gap(24).Children(func() {
					for _, s := range []struct {
						name string
						make func(*ui.Context) ui.Element
					}{
						{"凸起 Raised", func(c *ui.Context) ui.Element { return mallow.Raised(ui.Box(c), mallow.Large) }},
						{"凹陷 Recessed", func(c *ui.Context) ui.Element { return mallow.Recessed(ui.Box(c), mallow.Large) }},
					} {
						ui.Column(c.Key(s.name)).Grow(1).Gap(10).AlignItems(ui.Center).Children(func() {
							s.make(c).Size(88, 88).Radius(mallow.RadiusCard)
							ui.Text(c, s.name).FontSize(12).TextColor(mallow.Muted)
						})
					}
				})
				// GrooveV parts a row the way Groove parts a column.
				ui.Row(c).Gap(16).AlignItems(ui.Stretch).Children(func() {
					ui.Text(c, "竖向凹槽左边").Grow(1).Basis(0).PaddingY(12).FontSize(13).TextColor(mallow.Muted)
					mallow.GrooveV(c)
					ui.Text(c, "竖向凹槽右边").Grow(1).Basis(0).PaddingY(12).FontSize(13).TextColor(mallow.Muted)
				})
			})
		})
	})
}

func main() {
	g := newGallery()
	web := runtime.GOOS == "js"
	if web {
		// In a browser the page is the window: there are no window buttons
		// to draw, and the browser finds the fonts.
		mallow.Font = "system-ui"
	}
	mygo.App.WhenReady(func() {
		w := mygo.NewWindow(mygo.WindowOptions{
			Title:             "mallow gallery",
			Width:             1180,
			Height:            820,
			MinWidth:          900,
			MinHeight:         600,
			TitleBarStyle:     mygo.TitleBarHidden,
			HideWindowButtons: true,
			Content:           ui.View(g.view),
		})
		if !web {
			g.win = w
		}
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
