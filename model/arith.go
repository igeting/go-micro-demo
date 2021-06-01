package model

import (
	"errors"
)

type Arith struct{}

type Args struct {
	A int
	B int
}

type Quotient struct {
	Pro int
	Quo int
	Rem int
}

func (a *Arith) Multiply(req Args, res *Quotient) error {
	res.Pro = req.A * req.B
	return nil
}

func (a *Arith) Divide(req Args, res *Quotient) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
