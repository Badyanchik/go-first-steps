package main

import (
	"fmt"
	"strings"
)

func average(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}

func add(arguments ...int) int {
	total := 0
	for _, i := range arguments {
		total += i
	}
	return total
}

func joinStringsArr(arguments ...string) string {
	return strings.Join(arguments[:], " ")
}

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x  - 1)
}

func fibonacci(x uint) uint {
	if x == 0 || x == 1 {
		return x
	}
	return fibonacci(x - 1) + fibonacci(x - 2)
}

func main()  {
	arr := []float64{23,4324,546432,346,33,1312,543}
	arrStrings := []string{"I", "love", "you", "Julia",}
	nextEven := makeEvenGenerator()
	fmt.Println(average(arr))
	fmt.Println(add(1,2,3,4,5))
	fmt.Println(joinStringsArr(arrStrings...))
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(factorial(10))
	fmt.Println(fibonacci(10))
}