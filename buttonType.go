package button

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

var basicAtlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
var buttonText = text.New(pixel.V(0, 0), basicAtlas)
var buttonsPos = []pixel.Vec{}
var buttonsText = []string{}
var i int

type buttonType struct {
	mousePos, buttonPos pixel.Vec
	text                string
	win                 pixel.Target
	min, max            int
	textSize            float64
}

//position 0,0 is down left corner
func (b *buttonType) createButton(buttonPos pixel.Vec, text string) {
	buttonsPos = append(buttonsPos, buttonPos)
	buttonsText = append(buttonsText, text)
}

func (b *buttonType) returnButtons() ([]string, []pixel.Vec) {
	return buttonsText, buttonsPos
}

func (b *buttonType) drawButtons(win pixel.Target, min, max int, textSize float64) {
	for i := min; i-1 < max; i++ {
		buttonText = text.New(pixel.V(buttonsPos[i].X, buttonsPos[i].Y), basicAtlas)
		fmt.Fprintln(buttonText, buttonsText[i])
		buttonText.Draw(win, pixel.IM.Scaled(buttonText.Orig, textSize))
		buttonText.Clear()
	}
	return
}

func (b *buttonType) interactButton(mousePos pixel.Vec, textSize float64) (bool, int) {
	i := 0
	for range buttonsText {
		buttonText = text.New(pixel.V(buttonsPos[i].X, buttonsPos[i].Y), basicAtlas)
		fmt.Fprintln(buttonText, buttonsText[i])
		Size := buttonText.Bounds().Size()
		Size.X *= textSize
		Size.Y *= textSize
		if buttonText.Bounds().ResizedMin(Size).Contains(mousePos) == true {
			return true, i + 1
		}
		buttonText.Clear()
		i++
	}
	return false, 0
}
