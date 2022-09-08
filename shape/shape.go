package shape

import "math"

type Shape interface {
	GetSquare() (float32, error)
	GetType() string
}

type Square struct {
	side float32
}

type Triangle struct {
	side1 float32
	side2 float32
	angle float32
}

func (s *Square) GetSquare() (float32, error) {
	return s.side * s.side, nil
}

func (_ *Square) GetType() string {
	return "square"
}

func (t *Triangle) GetSquare() (float32, error) {
	return 0.5 * t.side1 * t.side2 * float32(math.Sin(float64(t.angle*math.Pi/180))), nil
}

func (_ *Triangle) GetType() string {
	return "triangle"
}

func NewSquare(side float32) *Square {
	return &Square{
		side: side,
	}
}

func NewTriangle(side1, side2, angle float32) *Triangle {
	return &Triangle{
		side1: side1,
		side2: side2,
		angle: angle,
	}
}
