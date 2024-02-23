package main

import (
	"os"

	"github.com/flopp/go-findfont"
	"github.com/goki/freetype/truetype"
	"github.com/vitsumoc/vmq-tool/app"
)

// 加载中文字体
func init() {
	fontPath, err := findfont.Find("simhei.ttf")
	if err != nil {
		panic(err)
	}

	// load the font with the freetype library
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	_, err = truetype.Parse(fontData)
	if err != nil {
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)
}

func main() {
	app.Run()
}
