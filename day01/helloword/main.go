package main

//import "fmt"
import "unsafe"
// var x, y int

// var(
// 	b int
// 	c bool
// )

// var d, e int = 1 , 2
// var f, g bool = true , false

const(
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)


func main() {
	// var a = "hellow word"
	// fmt.Println(a)

	// var b bool
	// fmt.Println(b)

	// var c int
	// fmt.Println(c)
	// x := 100
	// fmt.Println(x)

	// const LENGTN int = 10
	// const WIDTH = 5
	// var area int
	// const a , b , c = 1 , false, "str"

	// area = LENGTN * WIDTH
	// fmt.Printf("面积为 : %d", area)

	// println()
	println(a, b, c)
}
