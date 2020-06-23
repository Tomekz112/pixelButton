# pixelButton
library for making buttons in golang engine (pixel)

Starting
===============
First type in terminal/cmd:
<p><code>go get github.com/Tomekz112/pixelButton</code></p>
remember to also import pixel! _https://github.com/faiface/pixel_

<p>Create new file import <code>import (
"github.com/faiface/pixel"
button "github.com/Tomekz112/pixelButton"
)</code> and you are ready to go! :)</p>

Explanation
===============

CreateButton()
-------------
<p>Creates button with given position (pixel.Vec, pixel.ZV is at middle) and given text (string)</p>


ReturnButtons()
------------
<p>Returns all buttons text ([]string) and position ([]pixel.Vec)</p>

DrawButtons()
-------------
<p>Draw buttons in given window (pixel.Target), draw buttons from slice position min (int) to max (int)</p>
<p>textSize (float64) is how much the text is scaled</p>

InteractButton()
----------------
<p>returns true, if mouse cursor position is touching button, and returns, number of button (int) which cursor is touching, if mouse isnt touching</p>
<p>button will return false, and number -1</p>
<p>In mouse position just type<code>win.MousePosition()</code>, and in textSize, type what u choosed at <code>DrawButtons()</code></p>

DeleteButtons(min,max int) 
------------------------
<p>Removes buttons ex.</p>
<p>if we have slice <code>[1 2 3 4 5 6 7 8 9 10]</code> and give min = 9 max = 10, we will get slice: <code>[1 2 3 4 5 6 7 8 9]</code>
<p>if we have slice <code>[1 2 3 4 5 6 7 8 9 10</code> and give min = 0 max = 1, we will get slice: <code>[2 3 4 5 6 7 8 9 10]</code>
<p>if we have slice <code>[1 2 3 4 5 6 7 8 9 10</code> and give min = 0 max = 10, we will get slice: <code>[]</code>
<p>if we have slice <code>[1 2 3 4 5 6 7 8 9 10</code> and give min = 5 max = 7, we will get slice: <code>1 2 3 4 5 8 9 10</code>



