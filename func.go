package main 
import "fmt"

func swap(a , b *int)int{

var c int

 c = *a
*a = *b
*b =  c

return c
}

func plus(a , b int)int{
return a + b
}

func plusplus(a , b, c int)int{
return a + b + c
}


func  main() {

var p int = 32
var q int = 76

fmt.Println("\nBefore swapping : " , p , q)

swap(&p , &q)

fmt.Println("\nAfter swapping : " , p , q)

fmt.Println("\nAddition of two numbers :" , plus(67,39))

fmt.Println("\nAddition of three numbers :" , plusplus(87,23,95))

}









