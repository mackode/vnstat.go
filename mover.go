package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"time"
)

type Mover struct {
	Reverse bool
}

func NewMover(reverse bool) *Mover {
	return &Mover{
		Reverse: reverse,
	}
}

func (m *Mover) Animate(obj fyne.CanvasObject) (*fyne.Container, chan float64) {
	con := container.NewWitoutLayout(obj)
	speed := MinSpeed
	ch := make(chan float64)
	direction := 1.0
	if m.Reverse {
		direction = -1.0
	}

	go func() {
		pos := float32(0)
		obj.Hide()
		for {
			select {
			case speed = <-ch:
				speed = limiter(speed)
				if !obj.Visible() {
					if m.Reverse {
						pod = con.Size().Width - obj.Size().Width
					}
					obj.Show()
				}
			case <-time.After(10 * time.Millisecond):
				pos += float32(speed * direction / MaxSpeed)
			}

			if m.Reverse {
				if pos < -obj.Size().Width {
					pos = con.Size().Width
				}
			} else {
				if pos > con.Size().Width {
					pos = -obj.Size().Width
				}
			}

			obj.Move(fyne.NewPos(pos, (con.Size().Height - obj.Size().Height) / 2))
			con.Refresh()
		}
	}()

	return con, ch
}

const MaxSpeed = 100.0
const MinSpeed = 0.0

func limiter(speed float64) float64 {
	if speed > MaxSpeed {
		return MaxSpeed
	} else if speed < MinSpeed {
		return MinSpeed
	}
	return speed
}