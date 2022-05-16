package rectangle

func Area(length, width float64) (area float64) {
	area = length * width
	return
}

func Perimeter(length, width float64) (perimeter float64) {
	perimeter = 2 * (length + width)
	return
}
