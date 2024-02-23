// 自定义组件 连接标签
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type connectLabel struct {
	widget.BaseWidget
	Name     *widget.Label
	BtnClose *widget.Button
	BtnDel   *widget.Button
}

// 构造
func NewConnectLabel(Name string) *connectLabel {
	comp := &connectLabel{
		Name:     widget.NewLabel(Name),
		BtnClose: widget.NewButton("断开", nil),
		BtnDel:   widget.NewButton("删除", nil),
	}
	comp.ExtendBaseWidget(comp)

	return comp
}

// 绘制
func (comp *connectLabel) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, container.NewHBox(
		comp.BtnClose, comp.BtnDel,
	), comp.Name)
	return widget.NewSimpleRenderer(c)
}

// 更新
func (comp *connectLabel) update() {}
