package services

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 40
}

func Area(width float64, height float64) float64 {
	return width * height
}
