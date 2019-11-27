package main

import "fmt"

func main()  {
	arr := []float64 {
		434,
		324,
		534,
		12,
		546,
	}
	x := arr[:2]
	y := make([]string, 3)
	copy(y, []string{"Badyanchik", "Julia", "Love"})
	fmt.Println(arr)
	fmt.Println(x)
	fmt.Println(y)
}