package app

import (
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/vitsumoc/vmq-tool/ui"
)

func Run() {
	// 主窗口
	a := fyneApp.NewWithID("vitsumoc.github.vmqtool")
	// 初始化UI 并运行窗口
	ui.Init(a)
}
