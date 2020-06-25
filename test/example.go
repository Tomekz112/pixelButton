package main

import (
	"fmt"

	"github.com/faiface/pixel"

	button "github.com/Tomekz112/pixelButton"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Buttons!",
		Bounds: pixel.R(0, 0, 1024, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	test := pixel.ZV
	test.X = 95
	test.Y = 95
	thirdtest := pixel.ZV
	thirdtest.X = 80
	thirdtest.Y = 30
	secondtest := pixel.ZV
	secondtest.X = 50
	secondtest.Y = 15
	button.CreateButton(secondtest, "Some_Button_number_1")
	button.CreateButton(test, "First Test Working! :)")
	button.CreateButton(thirdtest, "Awesome! :)")
	for !win.Closed() {
		win.Clear(colornames.Black)
		interact, number := button.InteractButton(win.MousePosition(), 2)
		if interact == true && win.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Println(number)
		}
		button.DrawButtons(win, 0, 2, 2)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
