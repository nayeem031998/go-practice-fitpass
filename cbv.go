package main
import "fmt"

func swap(x , y int)int {

var c = 0
c = x
x = y
y = c

return c

}


func main() {

var a = 78
var b = 24

fmt.Println("Before swpping :" , a , b)

fmt.Println("\nAfter swapping :" , swap(78 , 24))

}


















