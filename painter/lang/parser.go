package lang

import (
	"io"

	"github.com/DipodidaeA/M-lab4/painter"
)

type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	res = append(res, painter.OperationFunc(painter.GreenFill))
	res = append(res, painter.UpdateOp)

	return res, nil
}
