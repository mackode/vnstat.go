package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
)

const SpriteFile = "sprite.png"

func mkPanel(isDownload bool) (*fyne.Container, func(float64)) {
	ava := NewFlicker(isDownload)
	err := ava.LoadSprite(SpriteFile)
	if err != nil {
		panic(err)
	}
	avaCon, avaC := ava.Animate()
	avaCon.Resize(fyne.NewSize(100, 100))
	mv := NewMover(isDownload)
	mvCon, mvCh := mv.Animate(avaCon)
	meter := widget.NewLabel("")

	throttle := func(v float64) {
		meter.Text = toBitRate(v)
		meter.Refresh()
		avaCh <- speedFromRate(v)
		mvCh <- speedFromRate(v)
	}

	panel := container.NewVBox(meter, mvCon)
	return panel, throttle
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Bandwidth Marathon")
	down, downUpdate := mkPanel(true)
	up, upUpdate := mkPanel(false)

	border := canvas.NewRectangle(theme.DisabledColor())
	dual := container.NewVBox(down, up)
	all := container.NewMax(boder, dual)
	myWindow.SetContent(all)
	myWindow.Resize(fyne.NewSize(float32(800), float32(300)))

	go func() {
		for {
			rx, tx, err := vnstat()
			if err != nil {
				panic(err)
			}
			upUpdate(tx)
			downUpdate(rx)
			time.Sleep(3 * time.Second)
		}
	}()

	myWindow.Canvas().SetOnTypedKey(
		func(ev *fyne.KeyEvent) {
			os.Exit(0)
		}
	)

	myWindow.ShowAndRun()
}