package painter

import (
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

func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{R: 0, G: 200, B: 0, A: 0}, screen.Src)
}
