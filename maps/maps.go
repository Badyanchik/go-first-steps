package main

import "fmt"

func main()  {
	x := map[string]map[string]string{
		"love": map[string]string{
			"name": "Julia",
			"age": "18",
		},
		"me": map[string]string{
			"name": "Bohdan",
			"age": "19",
		},
	}
	x["testDelete"] = map[string]string{
		"test": "test",
	}
	if testDelete, ok := x["testDelete"]; ok {
		fmt.Println(testDelete)
	}
	delete(x, "testDelete")
	if testDelete, ok := x["testDelete"]; ok {
		fmt.Println(testDelete)
	} else {
		fmt.Println("key testDelete is nil")
	}
	fmt.Println(x["love"]["name"])
	fmt.Println(x["me"]["name"])
	fmt.Println(len(x))
}