package main

import "fmt"

// func swap(x, y string) (string, string){
// 	return y, x
// }


// func main() {
// 	a:= func (x, y string) (string, string){
// 		return y,x
// 	}
// 	fmt.Println(a("你好", "haha"))
// }

// type Books struct{
// 	titile string
// 	author string
// 	subject string
// 	book_id int
// }


// func main(){
// 	var books1 Books
// 	//var books2 Books

// 	books1.titile = "Go 语言"
// 	books1.author = "123"
// 	books1.subject = "3333"
// 	books1.book_id = 1

// 	fmt.Printf( "books 1 title : %s\n", books1.titile)
// 	fmt.Printf( "Book 1 author : %s\n", books1.author)
// 	fmt.Printf( "Book 1 subject : %s\n", books1.subject)
// 	fmt.Printf( "Book 1 book_id : %d\n", books1.book_id)
// }
func main(){
	var n[10] int
	var i int

	for i = 0; i < 10; i++{
		n[i] = i + 100
	}

	var number int = n[9]

	fmt.Printf("%d", number)
	// for j = 0; j < 10; j++{
	// 	fmt.Printf("i[%d] = %d ", j , n[j])
	// }
}