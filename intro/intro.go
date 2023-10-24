package intro

func Return(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func Suma(values ...int) (int, error) {
	var result int
	for _, v := range values {
		result += v
	}
	return result, nil
}
