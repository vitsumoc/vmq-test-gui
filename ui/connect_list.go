// 连接管理列表
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type UIConnectList struct {
	// 顶层布局
	content *fyne.Container // VBox 布局
	// 交互部分
	btnAdd *widget.Button // 添加连接

	// 数据部分
	connList *widget.List
	// TODO 此处直接引用 APP 的连接列表
}

func newUIConnList() *UIConnectList {
	ui := &UIConnectList{}

	listData := []string{"192.168.100.1:1883", "192.168.100.2:1883", "192.168.100.1:1883"}
	ui.connList = widget.NewList(
		func() int {
			return len(listData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(listData[i])
		},
	)
	ui.btnAdd = widget.NewButton("创建连接", func() {})
	ui.content = container.NewBorder(ui.btnAdd, nil, nil, nil, ui.connList)

	return ui
}

func NewConnectList() fyne.CanvasObject {
	// 首先是一个上下布局, 上面是按钮组, 下面是连接列表
	// 暂时处在 UI 设计阶段, 内容可以先静态
	return newUIConnList().content
}
