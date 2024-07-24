package mypackage

import (
	"fmt"
	"math"
	"reflect"
)

type figure2D interface {
	GetArea() float64
}

type Square struct {
	Base float64
}

type Rectangle struct {
	High  float64
	Width float64
}

type Trapezoid struct {
	BaseA float64
	BaseB float64
	High  float64
}

type Circle struct {
	Radio float64
}

func (s Square) GetArea() float64 {
	return s.Base * s.Base
}

func (r Rectangle) GetArea() float64 {
	return r.High * r.Width
}

func (t Trapezoid) GetArea() float64 {
	return ((t.BaseA + t.BaseB) / 2) * t.High
}

func (c Circle) GetArea() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

func CalculateArea(f figure2D) string {
	return fmt.Sprintf("The Area of %s: %.2f", reflect.TypeOf(f).Name(), f.GetArea())
}
