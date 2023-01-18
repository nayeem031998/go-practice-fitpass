package main 
import "fmt"

func main () {

var a int


fmt.Println("Enter the Year : ")
fmt.Scanf("%d", &a)

if ((a%4 == 0 && a%100 != 0) || (a%4 == 0 && a%100 == 0 && a%400 == 0)) {
     fmt.Println("This Year is a Leap Year")
}else {
     fmt.Println("This Year is not a Leap Year")
}

}














