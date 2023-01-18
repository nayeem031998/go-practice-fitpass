package main 
import "fmt"

func swap(x , y *int)int {

var c int
c  = *x
*x = *y
*y = c

return c

}

func main() {

var a = 28
var b = 33

fmt.Println("Before swapping :" , a , b)

swap(&a , &b)

fmt.Println("After swapping :" , a , b)

}












