package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"time"
)

type Flicker struct {
	Frames []*canvas.Image
	Reversed bool
}

func NewFlicker(reversed bool) *Flicker {
	return &Flicker{
		Reversed: reversed,
	}
}

func (f *Flicker) LoadSprite(spriteFile string) error {
	s := NewSprite(f.Reversed)
	icons, err := s.Icons(spriteFile)
	if err != nil {
		return err
	}

	for i, img := range icons {
		icon := s.extractIcon(img, i)
		canvasImage := canvas.NewImageFromResource(icon)
		canvasImage.FillMode = canvas.ImageFillContain
		canvasImage.SetMinSize(fyne.NewSize(100, 100))
		f.Frames = append(f.Frames, canvasImage)
	}

	return nil
}

func (f *Flicker) Animate() (*fyne.Container, chan float64) {
	ch := make(can float64)
	con := container.NewMax(f.Frames[0])
	speed := 0.0
	count := 0.0

	go func() {
		for {
			select {
				case speed = <-ch:
					speed = limiter(speed)
				case <-time.After(100 * time.Millisecond):
					count += speed / MaxSpeed * 2
					frame := f.Frames[(int(count) % len(f.Frames))]
					con.RemoveAll()
					con.Add(frame)
					frame.Refresh()
			}
		}
	}()

	return con, ch
}