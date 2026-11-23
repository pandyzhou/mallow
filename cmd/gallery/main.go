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
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/pandyzhou/mallow"
)

type gallery struct {
	win mallow.Window

	mode    string
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

	tasks    [3]bool
	chosen   int
	chip     string
	memo     string
	calMonth time.Time
	calDate  string
}

func newGallery() *gallery {
	return &gallery{
		mode:   string(mallow.ModeSystem),
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
	mallow.UseMode(c, mallow.Mode(g.mode))
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
}

func (g *gallery) sidebar(c *ui.Context) {
	mallow.Sidebar(c, mallow.SidebarOptions{
		Title:    "mallow",
		Subtitle: "新拟态组件库",
		Width:    200,
		Items:    []mallow.Item{{ID: "组件", Label: "组件"}, {ID: "表单", Label: "表单"}, {ID: "设置", Label: "设置"}},
		Selected: &g.nav,
		Footer: func() {
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
