package main

import "fmt"

//Defining multiple variables
var (
	a = 5
	b = 10
	c = 15
)

var str string = "Hello Blaine!"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func multipleReturnTypes() (int, int) {
	return 5, 6
}

func main() {
	addition := 5
	addition += 1
	fmt.Println("value of addition: ", addition)
	const x string = "This is a constant called x"
	fmt.Printf("hello, world Bugsy!!\n")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(str)
	fmt.Println(x)

	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	var arr [5]int
	arr[4] = 100
	fmt.Println(arr)

	myMap := make(map[string]int)
	myMap["key"] = 10
	fmt.Println(myMap["key"])

	delete(myMap, "key")
	fmt.Println(myMap["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	if name, ok := elements["O"]; ok {
		fmt.Println(name, ok)
		if name, ok := elements["Un"]; ok {
			fmt.Println(name, ok)
		}
	}
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	//Multiple return types
	multipleReturnTypes()
	first, second := multipleReturnTypes()
	fmt.Println("First: ", first, " Second: ", second)

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
