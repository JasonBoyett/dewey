package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	titleSize := 60
	textSizt := 30
	bgBlue := color.RGBA{9, 74, 109, 255}
	bgPurple := color.RGBA{55, 14, 127, 255}
	a := app.New()
	w := a.NewWindow("Clock")
	split := layout.NewGridLayoutWithRows(2)

	topCanvas := canvas.NewText(
		"Dewey the Digital Librarian",
		color.RGBA{70, 201, 181, 255},
	)
	bottomCanvas := canvas.NewText(
		"",
		color.RGBA{119, 21, 216, 255},
	)
	topCanvas.TextStyle.Bold = true
	bottomCanvas.TextStyle.Bold = true
	topCanvas.TextSize = float32(titleSize)
	bottomCanvas.TextSize = float32(textSizt)
	topCanvas.Alignment = fyne.TextAlignCenter
	bottomCanvas.Alignment = fyne.TextAlignCenter
	go updateTime(
		bottomCanvas,
		*time.FixedZone("CST", -7*60*60),
	)
	content := container.New(
		split,
		topCanvas,
		bottomCanvas,
	)
	gradient := canvas.NewHorizontalGradient(bgPurple, bgBlue)
	content = container.New(
		layout.NewBorderLayout(nil, nil, nil, nil),
		gradient,
		content,
	)
	go updateGradientStart(
		gradient,
		127,
		0,
	)
	go updateGradientEnd(
		gradient,
		109,
		30,
	)

	w.SetContent(content)
	w.SetMaster()
	w.Show()

	a.Run()
}

func updateTime(textCanvas *canvas.Text, zone time.Location) {
	for {
		time.Sleep(time.Second)
		t := time.Now().UTC().In(&zone)
		currentTime := t.Format("It's 03:04:05")
		textCanvas.Text = currentTime
		textCanvas.Refresh()
	}
}

func updateGradientStart(gradient *canvas.LinearGradient, max, min uint32) {
	b := uint32(0)
	increment := true
	for {
		r, g, _, a := gradient.StartColor.RGBA()
		if increment {
			b += 1
		} else {
			b -= 1
		}
		gradient.StartColor = color.RGBA{
			uint8(r),
			uint8(g),
			uint8(b),
			uint8(a),
		}
		gradient.Refresh()
		if b > uint32(max) {
			increment = false
		}
		if b <= min {
			increment = true
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func updateGradientEnd(gradient *canvas.LinearGradient, max, min uint32) {
	r := uint32(0)
	increment := true
	for {
		_, g, b, a := gradient.EndColor.RGBA()
		if increment {
			r += 1
		} else {
			r -= 1
		}
		gradient.EndColor = color.RGBA{
			uint8(r),
			uint8(g),
			uint8(b),
			uint8(a),
		}
		gradient.Refresh()
		if r > uint32(max) {
			increment = false
		}
		if r <= min {
			increment = true
		}
		time.Sleep(time.Millisecond * 50)
	}
}
