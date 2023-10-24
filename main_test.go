package main

import (
	"api/intro"
	"testing"
)

func TestFibo(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := intro.Fibo(item.a)
		if fib != item.n {
			t.Errorf("Fibo was incorrect, got %d expected %d", fib, item.n)
		}
	}
}

func main() {

	// var dos, tres, cuatro = intro.Return(1)

	// fmt.Println(dos, tres, cuatro)

	// var d, err = intro.Suma(1, 2, 4, 4)
	// if err != nil {
	// 	fmt.Println("something was wrong")
	// } else {
	// 	fmt.Println(d)
	// }
	// fmt.Println("main")

	// func(n int) {
	// 	fmt.Println(n * n)
	// }(5)
}
