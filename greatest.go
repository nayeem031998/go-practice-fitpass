package main
import "fmt"

func main() {

var a int
var b int
var c int

fmt.Println("Enter the first number: ")
fmt.Scanf("%d", &a)

fmt.Println("Enter the second number: ")
fmt.Scanf("%d", &b)

fmt.Println("Enter the third number: ")
fmt.Scanf("%d", &c)

if (a > b && a > c) {
fmt.Println("First number is greater")
}else if (b > a && b > c) {
fmt.Println("Second number is greater")
}else {
fmt.Println("Third number is greater")
}

}










