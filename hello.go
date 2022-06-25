package main

import "fmt"
// import fm "fmt" <providing an alias to the package fmt>

func main() {
	sayHello()
	convertFloatToInt()
}

func sayHello() {
	fmt.Println("Hello World")
}

func convertFloatToInt() {
	var num float32 = 5.2
	fmt.Println(num)

	// Go doesn't allow implicit type casting
	fmt.Println(int(num))
}