package model

import "errors"

type Product struct {
	Name     string
	Color    string
	Category string
}

// Business rule happens here

type Color int

const (
	Red Color = iota + 1
	Green
	Blue
)

func (c Color) String() string {
	return [...]string{"Red", "Green", "Blue"}[c-1]
}

// The product only accept with 3 colors,
// because the stock limited for other colors
func (p Product) VerifyColor() error {
	switch p.Color {
	case Color.String(1): // Red
		return nil
	case Color.String(2): // Green
		return nil
	case Color.String(3): // Blue
		return nil
	}
	return errors.New("only accept red, green and blue colors")
}
