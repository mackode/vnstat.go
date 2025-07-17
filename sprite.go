package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

type Sprite struct {
	xOff, yOff		int
	width, height	int
	xPad, yPad		int
	columns			int
	reversed 		bool
}

func NewSprite(reversed bool) *Sprite {
	return &Sprite{
		xOff:     313,
		yOff:     67,
		width:    205,
		height:   258,
		xPad:     27,
		yPad:     39,
		columns:  5,
		reversed: reversed,
	}
}

func (s *Sprite) Icons(file string) ([]image.Image, error) {
	icons := []image.Image{}
	img, err := loadPNG(file)
	if err != nil {
		return icons, err
	}

	for i := 0; i < 10; i++ {
		icon := s.extractIcon(img, i)
		if s.reversed {
			icon = flipH(icon)
		}
		icons = append(icons, icon)
	}

	return icons, nil
}

func loadPNG(pat string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Decode the PNG image
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (s *Sprite) extractIcon(sheet image.Image, idx int) image.Image {
	col := idx % s.columns
	row := idx / s.columns
	x := s.xOff + col * (s.width + s.xPad)
	y := s.yOff + row * (s.height + s.yPad)

	rect := image.Rect(x, y, x + s.width, y + s.height)
	icon := image.NewRGBA(rect)
	draw.Draw(icon, rect, sheet, image.Point{x, y}, draw.Src)

	return icon
}

func flipH(src image.Image) image.Image {
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	dst := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			flippedX := bounds.Max.X - x - 1
			dst.Set(flippedX, bounds.Min.Y + y, src.At(bounds.Min.X + x, bounds.Min.Y + y))
		}
	}

	return flipped
}