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
	variable	    float64
}

//CreateButton at give position, and given text
//position 0,0 is down left corner
func CreateButton(buttonPos pixel.Vec, text string) {
	buttonsPos = append(buttonsPos, buttonPos)
	buttonsText = append(buttonsText, text)
}

//ReturnButtons returns slice in which there are all buttons text and position
func ReturnButtons() ([]string, []pixel.Vec) {
	return buttonsText, buttonsPos
}

//Edits button in slice with given position(min)
func EditButton(min int, buttonPos pixel.Vec, text string) {
		buttonsPos[min] = buttonPos
		buttonsText[min] = text
}

//DrawButtons draw buttons from given position(min) to given position(max) in slice
func DrawButtons(win pixel.Target, min, max int, textSize float64) {
	for i := min; i-1 < max; i++ {
		buttonText = text.New(pixel.V(buttonsPos[i].X, buttonsPos[i].Y), basicAtlas)
		fmt.Fprintln(buttonText, buttonsText[i])
		buttonText.Draw(win, pixel.IM.Scaled(buttonText.Orig, textSize))
		buttonText.Clear()
	}
	return
}

//min is which button u want to draw and variable is which variable u want to add to button
func DrawVariableButtons(win pixel.Target, min int, textSize float64, variable float64) {
		buttonText = text.New(pixel.V(buttonsPos[min].X, buttonsPos[min].Y), basicAtlas)
		fmt.Fprintln(buttonText, buttonsText[min],variable)
		buttonText.Draw(win, pixel.IM.Scaled(buttonText.Orig, textSize))
		buttonText.Clear()
	        return
}

//InteractButton return true whenever mouse cursor is on button and number of button moues cursor
//is interacting with
func InteractButton(mousePos pixel.Vec, textSize float64) (bool, int) {
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

//DeleteButtons deletes button from min to max
//ex. we have slice with [0 1 2 3 4 5 6 7 8 9],
//if min = 9 and max = 10
//the slice will be:[1 2 3 4 5 6 7 8 10]
func DeleteButtons(min, max int) {
	buttonsPos = append(buttonsPos[:min], buttonsPos[max:]...)
	buttonsText = append(buttonsText[:min], buttonsText[max:]...)
}
