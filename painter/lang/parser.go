package lang

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/DipodidaeA/M-lab4/painter"
)

type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, []painter.Cord, error) {
	var opRes []painter.Operation
	var opWindowBackground []painter.Operation
	var opRect []painter.Operation
	var opUpdate []painter.Operation

	var cordRes []painter.Cord
	var cordWindowBackground []painter.Cord
	var cordRect []painter.Cord
	var cordUpdate []painter.Cord

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		commandLine := scanner.Text()
		op, x1, y1, x2, y2 := p.ParseText(commandLine)
		switch op {
		case "Fwhite":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.WhiteFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Fgreen":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.GreenFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Fred":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.RedFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Fblack":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.BlackFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Fblue":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.BlueFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Fyellow":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.YellowFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Rblack":
			opRect = append(opRect, painter.OperationFunc(painter.BlackRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Rred":
			opRect = append(opRect, painter.OperationFunc(painter.RedRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Rgreen":
			opRect = append(opRect, painter.OperationFunc(painter.GreenRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Rblue":
			opRect = append(opRect, painter.OperationFunc(painter.BlueRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "Ryellow":
			opRect = append(opRect, painter.OperationFunc(painter.YellowRect))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "reset":
			opWindowBackground = append(opWindowBackground, painter.OperationFunc(painter.BlackFill))
			cordWindowBackground = append(cordWindowBackground, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
			opRect = append(opRect, painter.OperationFunc(painter.BlackFill))
			cordRect = append(cordRect, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		case "update":
			opUpdate = append(opUpdate, painter.UpdateOp)
			cordUpdate = append(cordUpdate, painter.Cord{X1: x1, Y1: y1, X2: x2, Y2: y2})
		}
	}

	opRes = append(opRes, opWindowBackground...)
	opRes = append(opRes, opRect...)
	opRes = append(opRes, opUpdate...)

	cordRes = append(cordRes, cordWindowBackground...)
	cordRes = append(cordRes, cordRect...)
	cordRes = append(cordRes, cordUpdate...)

	return opRes, cordRes, nil
}

func (p *Parser) ParseText(lineText string) (string, float64, float64, float64, float64) {
	command := strings.Fields(lineText)
	size := len(command)
	com := command[0]
	x1, y1, x2, y2 := float64(0), float64(0), float64(0), float64(0)
	if size >= 3 {
		if command[1] != "" {
			x1, _ = strconv.ParseFloat(command[1], 64)
		}
		if command[2] != "" {
			y1, _ = strconv.ParseFloat(command[2], 64)
		}
	}
	if size == 5 {
		if command[3] != "" {
			x2, _ = strconv.ParseFloat(command[3], 64)
		}
		if command[4] != "" {
			y2, _ = strconv.ParseFloat(command[4], 64)
		}
	}
	return com, x1, y1, x2, y2
}
