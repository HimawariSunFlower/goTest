package drawer

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
)

//https://dxkite.cn/index.php/article/golang-gif-transparent gif与透明色背景

func MergeImages(first image.Image, subs ...image.Image) image.Image {
	wi := 900
	hi := 900
	//图片三合一绘图
	png := image.NewNRGBA(image.Rect(0, 0, wi, hi))
	pt := image.Pt(first.Bounds().Min.X, first.Bounds().Min.X)
	//画布   画布起画点  待画图   从待画图的哪里开始画    画图方法op,是否画在遮罩上
	draw.Draw(png, png.Bounds().Add(pt), first, first.Bounds().Min, draw.Over)
	for _, subImg := range subs {
		draw.Draw(png, png.Bounds().Add(image.Pt(subImg.Bounds().Min.X, subImg.Bounds().Min.X)), subImg, subImg.Bounds().Min, draw.Over)
	}
	return png
}

var picCloth = []image.Image{}
var picWeapon = []image.Image{}

func makeGif() *gif.GIF {

	mid := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	for i := 0; i < len(picCloth); i++ {
		var img image.Image

		img = MergeImages(picCloth[i], picWeapon[i])

		bounds := img.Bounds()

		p := image.NewPaletted(bounds, getSubPalette(img))

		draw.Draw(p, p.Bounds(), img, img.Bounds().Min, draw.Over)
		mid = append(mid, p)
		delays = append(delays, 100/60)
	}

	return &gif.GIF{
		Image: mid,
		Delay: delays,
	}
}

func getSubPalette(m image.Image) color.Palette {
	p := color.Palette{color.RGBA{0x00, 0x00, 0x00, 0x00}}
	p9 := color.Palette(palette.Plan9)
	b := m.Bounds()
	black := false
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := m.At(x, y)
			cc := p9.Convert(c)
			if cc == p9[0] {
				black = true
			}
			if inPalette(p, cc) == -1 {
				p = append(p, cc)
			}
		}
	}
	if len(p) < 256 && black == true {
		p[0] = color.RGBA{0x00, 0x00, 0x00, 0x00} // transparent
		p = append(p, p9[0])
	}
	return p
}

func inPalette(p color.Palette, c color.Color) int {
	ret := -1
	for i, v := range p {
		if v == c {
			return i
		}
	}
	return ret
}
