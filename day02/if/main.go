package main

import "fmt"

func main() {
	// var a int = 100
	// var b int = 300

	// if a == 100{
	// 	if b == 200{
	// 		fmt.Printf("a的值等于100，b的值等于200\n")
	// 	} else {
	// 		fmt.Println("I am your father")
	// 	}
	// }
	// fmt.Println(a)
	
	// fmt.Println(b)
	// var grade string
	// var marks int = 90

	// switch marks {
	// case 90 : grade = "A"
	// case 80 : grade = "B"
	// case 50,60,70 : grade = "C"
	// default:grade = "D"

	// }

	// switch grade{
	// case "A" : 
	// 	fmt.Printf("优秀！\n")
	// case "B" , "C" :
	// 	fmt.Printf("良好\n")
	// case "D" :
	// 	fmt.Printf("及格\n")
	// default :
	// 	fmt.Printf("差\n")	
	// }
	// fmt.Printf("等级是 %s", grade)

	// var i , j int

	// for i = 2; i < 100; i++{
	// 	for j = 2; j <= (i / j); j++{
	// 		if(i % j == 0){
	// 			break
	// 		}
	// 	}
	// 	if(j > (i / j)){
	// 		fmt.Printf("%d\n", i)
	// 	}
	// }

	var a int = 100
	var b int = 200

	var ret = max(a,b)
	fmt.Printf("%d", ret)
}


func max(a, b int) int {
	var result int

	if(a > b){
		result = a
	} else {
		result = b
	}
	return result
}
