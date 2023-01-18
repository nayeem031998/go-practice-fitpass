package main 
import "fmt"

func main() {

var a int = 46

var b *int
b = &a

fmt.Println("value stored in a : ", a)
fmt.Println("value stored in b : ", b)
fmt.Println("address of a : ", &a)

//how to change element for a particular address

var c *int
c = &a
*c = 23
fmt.Println(a)





}



