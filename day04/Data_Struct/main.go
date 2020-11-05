package main

import "fmt"

//slice text
// func main(){
// 	// numbers := []int{0,1,2,3,4,5,6,7,8}
// 	// printSlice(numbers)

// 	// fmt.Println("numbers == ", numbers)

// 	// fmt.Println("numbers[1 : 4] == ", numbers[1:4])

// 	// fmt.Println("numbers[ : 3] == ",numbers[:3])

// 	// fmt.Println("numbers[3:] == ", numbers[3:])
	
// 	var numbers [] int
// 	PrintSlice(numbers)

// 	numbers = append(numbers, 0)
// 	PrintSlice(numbers)

// 	numbers = append(numbers, 1)
// 	PrintSlice(numbers)
	
	
// 	numbers = append(numbers, 2,3,4)
// 	PrintSlice(numbers)

// 	numbers1 := make([] int, len(numbers), (cap(numbers)) * 2)

// 	copy(numbers1, numbers)
// 	PrintSlice(numbers1)


// }

// func PrintSlice(x[] int){
// 	fmt.Printf("size = %d, cap = %d, slice = %v \n", len(x), cap(x), x)
// }


// func main(){
// 	var CountryCapitalMap map[string]string
// 	CountryCapitalMap = make(map[string] string)

// 	CountryCapitalMap["France"] = "巴黎"
// 	CountryCapitalMap["Japen"] = "东京"
// 	CountryCapitalMap["China"] = "北京"

// 	for country := range CountryCapitalMap{
// 		fmt.Println(country, " 首都是 ", CountryCapitalMap[country])
		
// 	}
	
// 	// for k, v := range CountryCapitalMap{
// 	// 	fmt.Println(k, " 首都是 ", v)
// 	// }

// 	delete(CountryCapitalMap, "France")

// 	for country := range CountryCapitalMap{
// 		fmt.Println(country, " 首都是 ", CountryCapitalMap[country])
		
// 	}

// 	// capacity, ok := CountryCapitalMap["American"]

// 	// if(ok){
// 	// 	fmt.Println("American的首都是", capacity)
// 	// }else{
// 	// 	fmt.Println("American的首都不存在")
// 	// }
// }

// func main(){
// 	vector := []int {1,2,3}
// 	var sum int

// 	for _, value := range vector{
// 		sum += value
// 	} 

// 	fmt.Println(sum)

// 	for key, value := range vector{
// 		if value == 2{
// 			fmt.Println(key)
// 		}
// 	}

// 	m_map := map[string]string{"a": "apple", "b":"banana"}
// 	for key, value := range m_map{
// 		fmt.Printf("%s -> %s\n", key, value)
// 	}


// }

// func fibonacci(num int) int{
// 	if(num < 2){
// 		return num
// 	}

// 	return fibonacci(num - 1) + fibonacci(num - 2)
	
// }

// func main(){
// 	var i int
// 	for i = 0; i <= 10; i++{
// 		fmt.Printf("i = %d\n", fibonacci(i))
// 	}
// }


func main(){
	var a int = 1
	var b float32

	b = float32(a)

	fmt.Printf("%f", b)

}