// 创建新连接
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// 页面数据
type DATA_NewConnect struct {
	IPAddr string // ip地址
	Port   string // 端口号
}

// 页面UI
type newConnect struct {
	// 顶层布局
	content *fyne.Container
	// 弹层对象
	dia *dialog.CustomDialog
	// 组件
	labelIp   *widget.Label
	labelPort *widget.Label
	inputIp   *widget.Entry
	inputPort *widget.Entry
	btnAdd    *widget.Button
	btnCancel *widget.Button

	// 数据
	data DATA_NewConnect
}

// 创建页面方法
func newNewConnect() *newConnect {
	nc := &newConnect{}
	nc.labelIp = widget.NewLabel("IP地址")
	nc.labelPort = widget.NewLabel("端口")
	nc.inputIp = widget.NewEntry()
	nc.inputPort = widget.NewEntry()
	nc.btnAdd = widget.NewButton("添加", nil)
	nc.btnCancel = widget.NewButton("取消", func() { nc.dia.Hide() })

	// 内容对象
	nc.content = container.New(
		layout.NewFormLayout(),
		nc.labelIp, nc.inputIp,
		nc.labelPort, nc.inputPort,
		widget.NewLabel(""), container.NewHBox(nc.btnAdd, nc.btnCancel),
	)

	// 弹层对象
	nc.dia = dialog.NewCustomWithoutButtons("创建新连接", nc.content, mainWindow)

	return nc
}

// 清理数据 弹出连接弹窗
func (p *newConnect) pop() {
	// 清理数据
	p.data.IPAddr = ""
	p.data.Port = ""
	// 弹出弹窗
	p.dia.Show()
}
