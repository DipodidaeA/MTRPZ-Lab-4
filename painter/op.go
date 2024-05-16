package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Do(t screen.Texture) (ready bool)
}

type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool {
	return true
}

type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

type Cord struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

var CordList []Cord

func PushCord(c Cord) {
	CordList = append(CordList, c)
}

func PullCord() Cord {
	res := CordList[0]
	CordList[0] = Cord{}
	CordList = CordList[1:]
	return res
}

func BlackFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

func WhiteFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.White, screen.Src)
}

func RedFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{R: 200, G: 0, B: 0, A: 0}, screen.Src)
}

func GreenFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{R: 0, G: 200, B: 0, A: 0}, screen.Src)
}

func BlueFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{R: 0, G: 0, B: 200, A: 0}, screen.Src)
}

func YellowFill(t screen.Texture) {
	PullCord()
	t.Fill(t.Bounds(), color.RGBA{R: 200, G: 200, B: 0, A: 0}, screen.Src)
}

func BlackRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 0, 0, 0, 0)
}

func RedRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 200, 0, 0, 0)
}

func GreenRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 0, 200, 0, 0)
}

func BlueRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 0, 0, 200, 0)
}

func YellowRect(t screen.Texture) {
	c := PullCord()
	sX1 := int(c.X1 * 800)
	sY1 := int(c.Y1 * 800)
	sX2 := int(c.X2 * 800)
	sY2 := int(c.Y2 * 800)
	RectFigureDraw(t, sX1, sY1, sX2, sY2, 200, 200, 0, 0)
}

func RectFigureDraw(t screen.Texture, x1, y1, x2, y2 int, r, g, b, a byte) {
	var pos image.Rectangle
	pos.Min.X = x1
	pos.Min.Y = y1
	pos.Max.X = x2
	pos.Max.Y = y2
	t.Fill(pos.Bounds(), color.RGBA{R: r, G: g, B: b, A: a}, screen.Src)
}
